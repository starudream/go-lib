package selfupdate

import (
	"bufio"
	"bytes"
	"crypto"
	_ "crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"runtime"
	"strconv"
	"strings"

	"github.com/cheggaaa/pb/v3"

	"github.com/starudream/go-lib/core/v2/gh"
	"github.com/starudream/go-lib/core/v2/slog"
	"github.com/starudream/go-lib/core/v2/utils/osutil"
	"github.com/starudream/go-lib/resty/v2"

	"github.com/starudream/go-lib/selfupdate/v2/internal/archive"
)

// GoReleaser
//
// {{ .Proxy }}{{ .Owner }}/{{ .Repo }}/releases/download/{{ .Version }}/{{ .Name or .Repo }}_{{ .Version }}_{{ .Os }}_{{ .Arch }}
type GoReleaser struct {
	Proxy string

	Owner string
	Repo  string
	Name  string
}

var _ Source = (*GoReleaser)(nil)

func (t *GoReleaser) Latest() (string, error) {
	t.proxy()
	resp, err := resty.R().SetDoNotParseResponse(true).Get(t.Proxy + t.Owner + "/" + t.Repo + "/releases/latest")
	defer gh.Close(resp.RawBody())
	if err != nil {
		return "", err
	}
	if resp.StatusCode() != 200 {
		return "", errors.New(resp.Status())
	}
	return t.findVersion(
		resp.RawResponse.Request.URL.Path,
		resp.Header().Get("X-Proxy-Redirect"),
	), nil
}

func (t *GoReleaser) proxy() {
	if t.Proxy == "" {
		t.Proxy = githubDefault
		if !resty.HasProxyEnv() && !resty.C().IsProxySet() {
			t.Proxy = (&GithubProxy{}).Check().Fast()
		}
	}
	if t.Proxy != githubDefault {
		slog.Debug("use proxy: %s", t.Proxy)
	}
}

func (t *GoReleaser) findVersion(vs ...string) string {
	for _, s := range vs {
		if s == "" {
			continue
		}
		sub := t.Owner + "/" + t.Repo + "/releases/tag/"
		idx := strings.LastIndex(s, sub)
		if idx >= 0 {
			return s[idx+len(sub):]
		}
	}
	return ""
}

func (t *GoReleaser) Binary(version string) ([]byte, error) {
	t.proxy()
	filename := t.filename(version)
	expected, err := t.checksum(version, filename)
	if err != nil {
		return nil, err
	}
	binary, err := t.download(version, filename)
	if err != nil {
		return nil, err
	}
	if actual := osutil.Must1(checksumFor(crypto.SHA256, binary)); !bytes.Equal(expected, actual) {
		return nil, fmt.Errorf("checksum not match. expected: %x, actual: %x", expected, actual)
	}
	return t.decompress(binary)
}

func (t *GoReleaser) checksum(version, filename string) ([]byte, error) {
	bs, err := t.download(version, "checksums.txt")
	if err != nil {
		return nil, err
	}
	sc := bufio.NewScanner(bytes.NewReader(bs))
	for sc.Scan() {
		ss := strings.Split(sc.Text(), "  ")
		if len(ss) == 2 && ss[1] == filename {
			return hex.DecodeString(ss[0])
		}
	}
	return nil, fmt.Errorf("not found checksum for %s", filename)
}

func (t *GoReleaser) download(version, filename string) ([]byte, error) {
	resp, err := resty.R().SetDoNotParseResponse(true).Get(t.Proxy + t.Owner + "/" + t.Repo + "/releases/download/" + version + "/" + filename)
	defer gh.Close(resp.RawBody())
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, errors.New(resp.Status())
	}
	size, _ := strconv.Atoi(resp.Header().Get("Content-Length"))
	bar := pb.New(size).Set(pb.Bytes, true).Set(pb.CleanOnFinish, true).Start()
	bs, err := io.ReadAll(bar.NewProxyReader(resp.RawBody()))
	if err != nil {
		return nil, err
	}
	bar.Finish()
	return bs, nil
}

func (t *GoReleaser) filename(version string) string {
	if t.Name == "" {
		t.Name = t.Repo
	}
	return fmt.Sprintf("%s_%s_%s_%s%s", t.Name, version, runtime.GOOS, runtime.GOARCH, t.ext())
}

func (t *GoReleaser) ext() string {
	//goland:noinspection GoBoolExpressions
	if runtime.GOOS == "windows" {
		return ".zip"
	}
	return ".tar.gz"
}

func (t *GoReleaser) decompress(data []byte) ([]byte, error) {
	switch t.ext() {
	case ".zip":
		return archive.Zip(data).Get(t.Name)
	case ".tar.gz":
		return archive.TarGz(data).Get(t.Name)
	default:
		panic("unreachable")
	}
}
