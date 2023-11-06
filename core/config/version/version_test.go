package version

import (
	"testing"
)

func TestGetVersionInfo(t *testing.T) {
	t.Log("\n" + GetVersionInfo().String())
}
