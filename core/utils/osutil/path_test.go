package osutil

import (
	"testing"
)

func TestWorkDir(t *testing.T) {
	t.Log(WorkDir())
}

func TestExe(t *testing.T) {
	t.Log(ExeFull())
	t.Log(ExeDir())
	t.Log(ExeName())
}
