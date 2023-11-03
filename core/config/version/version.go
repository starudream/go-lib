package version

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"strings"
	"sync"
	"text/tabwriter"
	"time"
)

const unknown = "unknown"

var (
	// Output of "git describe". The prerequisite is that the
	// branch should be tagged using the correct versioning strategy.
	gitVersion = "devel"
	// SHA1 from git, output of $(git rev-parse HEAD)
	gitCommit = unknown
	// State of git tree, either "clean" or "dirty"
	gitTreeState = unknown
	// Build date in ISO8601 format, output of $(date -u +'%Y-%m-%dT%H:%M:%SZ')
	buildDate = unknown
	// goVersion is the used golang version.
	goVersion = unknown
	// compiler is the used golang compiler.
	compiler = unknown
	// platform is the used os/arch identifier.
	platform = unknown
)

type Info struct {
	GitVersion   string `json:"gitVersion"`
	GitCommit    string `json:"gitCommit"`
	GitTreeState string `json:"gitTreeState"`
	BuildDate    string `json:"buildDate"`
	GoVersion    string `json:"goVersion"`
	Compiler     string `json:"compiler"`
	Platform     string `json:"platform"`
}

func getBuildInfo() *debug.BuildInfo {
	bi, ok := debug.ReadBuildInfo()
	if !ok {
		return nil
	}
	return bi
}

func getGitVersion(bi *debug.BuildInfo) string {
	if bi == nil {
		return unknown
	}
	if bi.Main.Version == "(devel)" || bi.Main.Version == "" {
		return gitVersion
	}
	return bi.Main.Version
}

func getCommit(bi *debug.BuildInfo) string {
	return getKey(bi, "vcs.revision")
}

func getDirty(bi *debug.BuildInfo) string {
	modified := getKey(bi, "vcs.modified")
	if modified == "true" {
		return "dirty"
	}
	if modified == "false" {
		return "clean"
	}
	return unknown
}

func getBuildDate(bi *debug.BuildInfo) string {
	buildTime := getKey(bi, "vcs.time")
	t, err := time.Parse("2006-01-02T15:04:05Z", buildTime)
	if err != nil {
		return unknown
	}
	return t.Format("2006-01-02T15:04:05.000Z07:00")
}

func getKey(bi *debug.BuildInfo, key string) string {
	if bi == nil {
		return unknown
	}
	for _, iter := range bi.Settings {
		if iter.Key == key {
			return iter.Value
		}
	}
	return unknown
}

var (
	_info     Info
	_infoOnce sync.Once
)

// GetVersionInfo represents known information on how this binary was built.
func GetVersionInfo() Info {
	_infoOnce.Do(func() {
		buildInfo := getBuildInfo()
		gitVersion = getGitVersion(buildInfo)
		if gitCommit == unknown {
			gitCommit = getCommit(buildInfo)
		}
		if gitTreeState == unknown {
			gitTreeState = getDirty(buildInfo)
		}
		if buildDate == unknown {
			buildDate = getBuildDate(buildInfo)
		}
		if goVersion == unknown {
			goVersion = runtime.Version()
		}
		if compiler == unknown {
			compiler = runtime.Compiler
		}
		if platform == unknown {
			platform = fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH)
		}
		_info = Info{
			GitVersion:   gitVersion,
			GitCommit:    gitCommit,
			GitTreeState: gitTreeState,
			BuildDate:    buildDate,
			GoVersion:    goVersion,
			Compiler:     compiler,
			Platform:     platform,
		}
	})
	return _info
}

// String returns the string representation of the version info
func (i Info) String() string {
	b := strings.Builder{}
	w := tabwriter.NewWriter(&b, 0, 0, 2, ' ', 0)
	_, _ = fmt.Fprintf(w, "GitVersion:\t%s\n", i.GitVersion)
	_, _ = fmt.Fprintf(w, "GitCommit:\t%s\n", i.GitCommit)
	_, _ = fmt.Fprintf(w, "GitTreeState:\t%s\n", i.GitTreeState)
	_, _ = fmt.Fprintf(w, "BuildDate:\t%s\n", i.BuildDate)
	_, _ = fmt.Fprintf(w, "GoVersion:\t%s\n", i.GoVersion)
	_, _ = fmt.Fprintf(w, "Compiler:\t%s\n", i.Compiler)
	_, _ = fmt.Fprintf(w, "Platform:\t%s\n", i.Platform)
	_ = w.Flush()
	return b.String()
}
