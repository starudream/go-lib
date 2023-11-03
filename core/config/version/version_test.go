package version

import (
	"testing"
)

func TestGetVersionInfo(t *testing.T) {
	t.Log(GetVersionInfo().String())
}
