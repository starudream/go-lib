package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"

	"github.com/Masterminds/semver/v3"
	"github.com/samber/lo"
	"golang.org/x/mod/modfile"

	"github.com/starudream/go-lib/core/v2/slog"
	"github.com/starudream/go-lib/core/v2/utils/osutil"
)

func main() {
	ms := map[string]*Module{}
	sc := bufio.NewScanner(bytes.NewReader(lo.Must1(exec.Command("git", "tag", "-l").Output())))
	for sc.Scan() {
		ss := strings.SplitN(sc.Text(), "/", 2)
		if len(ss) != 2 {
			slog.Warn("skip", slog.String("tag", sc.Text()))
			continue
		}
		name := ss[0]
		version, err := semver.NewVersion(ss[1])
		if err != nil {
			slog.Error("skip", slog.String("tag", sc.Text()), slog.Any("error", err))
			continue
		}
		if _, ok := ms[name]; !ok {
			ms[name] = &Module{Name: name, Require: map[string]string{}}
		}
		ms[name].Versions = append(ms[name].Versions, version)
	}
	px, sx := "github.com/starudream/go-lib/", "/v2"
	for n := range ms {
		sort.Sort(semver.Collection(ms[n].Versions))
		file := filepath.Join(osutil.WorkDir(), n, "go.mod")
		mod := lo.Must1(modfile.ParseLax(file, lo.Must1(os.ReadFile(file)), nil))
		for _, r := range mod.Require {
			if !r.Indirect && strings.HasPrefix(r.Mod.Path, px) {
				name := strings.TrimSuffix(strings.TrimPrefix(r.Mod.Path, px), sx)
				ms[n].Require[name] = r.Mod.Version
			}
		}
	}
	for n := range ms {
		ms[n].nrCnt = len(ms[n].Require)
	}
	cnt, left := 1, len(ms)
tag:
	dm := map[string]bool{}
	for n := range ms {
		if ms[n].Next != "" {
			continue
		}
		if ms[n].nrCnt == 0 {
			version := ms[n].Versions[len(ms[n].Versions)-1]
			ms[n].Next = version.IncPatch().String()
			dm[n] = true
			fmt.Printf("git tag -a -m \"\" %s\n", n+"/v"+ms[n].Next)
			left--
		}
	}
	for n := range dm {
		for m := range ms {
			if _, ok := ms[m].Require[n]; ok {
				ms[m].nrCnt--
			}
		}
	}
	if left > 0 {
		fmt.Printf("----------\n")
		cnt++
		goto tag
	}
}

type Module struct {
	Name     string            `json:"name"`
	Versions []*semver.Version `json:"versions"`
	Next     string            `json:"next,omitempty"`
	Require  map[string]string `json:"require,omitempty"`

	nrCnt int
}
