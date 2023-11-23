package version

import (
	"testing"
)

func TestGetVersionInfo(t *testing.T) {
	t.Log("\n" + GetVersionInfo().String())
}

func TestParse(t *testing.T) {
	text := `
GitVersion:    v2.0.0
GitCommit:     f4a79ceefe31e118b81ad65e3e68250462cc82d2
GitTreeState:  dirty
BuildDate:     2023-11-23T09:44:42.000Z
GoVersion:     go1.21.4
Compiler:      gc
Platform:      darwin/arm64
`
	t.Logf("\n%#v", Parse(text))
}
