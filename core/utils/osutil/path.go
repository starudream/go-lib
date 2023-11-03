package osutil

import (
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var (
	_workDir     string
	_workDirOnce sync.Once
)

func WorkDir() string {
	_workDirOnce.Do(func() {
		wd, err := os.Getwd()
		ExitErr(err)
		_workDir = wd
	})
	return _workDir
}

var (
	_exeFull string
	_exeDir  string
	_exeName string
	_exeOnce sync.Once
)

func exeInit() {
	_exeOnce.Do(func() {
		exe, err := os.Executable()
		ExitErr(err)
		_exeFull = exe
		_exeDir = filepath.Dir(_exeFull)
		_exeName = strings.TrimSuffix(filepath.Base(_exeFull), filepath.Ext(_exeFull))
	})
}

func ExeFull() string {
	exeInit()
	return _exeFull
}

func ExeDir() string {
	exeInit()
	return _exeDir
}

func ExeName() string {
	exeInit()
	return _exeName
}
