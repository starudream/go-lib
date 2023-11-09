package osutil

import (
	"encoding/json"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
)

type GoList struct {
	Dir        string `json:"Dir"`
	ImportPath string `json:"ImportPath"`
	Name       string `json:"Name"`
	Root       string `json:"Root"`
	Module     struct {
		Path      string `json:"Path"`
		Main      bool   `json:"Main"`
		Dir       string `json:"Dir"`
		GoMod     string `json:"GoMod"`
		GoVersion string `json:"GoVersion"`
	} `json:"Module"`
}

var (
	_goList     GoList
	_goListOnce sync.Once
)

func GetGoList() GoList {
	_goListOnce.Do(func() {
		output, err := exec.Command("go", "list", "-json").Output()
		if err == nil {
			_ = json.Unmarshal(output, &_goList)
		}
	})
	return _goList
}

func GoListRoot() string {
	if strings.HasPrefix(GetGoList().Module.Path, "github.com/starudream/go-lib") && strings.HasSuffix(GetGoList().Module.Path, "/v2") {
		return filepath.Join(GetGoList().Root, "..")
	}
	return GetGoList().Root
}
