package config

import (
	"os"

	"github.com/knadh/koanf/v2"

	"github.com/starudream/go-lib/core/v2/codec/json"
	"github.com/starudream/go-lib/core/v2/utils/osutil"

	"github.com/starudream/go-lib/core/v2/internal/logs"
)

var _k *koanf.Koanf

// The priority is the following:
// 1. env
// 2. flag's config file
// 3. files (if flag's config file is not set)
func init() {
	_k = koanf.New(".")

	// file
	files := Files()
	for i := 0; i < len(files); i++ {
		file := files[i]
		if file == "" {
			continue
		}
		logs.D("attempt to load config file", "file", file)
		err := LoadFile(file)
		if err != nil {
			// exit if the first file (from flag set --config or -c) is not exist
			if i < _filesMustSize || !os.IsNotExist(err) {
				osutil.ExitErr(err)
			}
			if !os.IsNotExist(err) {
				logs.D("config file load failed", "err", err)
			}
		} else {
			logs.D("config file loaded")
			break
		}
	}

	// env
	LoadEnv()
	logs.D("config env loaded")

	logs.D(json.MustMarshalString(_k.All()))
}

func All(path ...string) map[string]any {
	if len(path) == 0 {
		return _k.All()
	}
	return _k.Cut(path[0]).All()
}

func Raw(path ...string) map[string]any {
	if len(path) == 0 {
		return _k.Raw()
	}
	return _k.Cut(path[0]).Raw()
}

func Get(path string) Value {
	return NewValue(_k.Get(path))
}

func Set(path string, value any) {
	_ = _k.Set(path, value)
}

func Del(path string) {
	_k.Delete(path)
}

func Unmarshal(path string, o any) error {
	return _k.UnmarshalWithConf(path, o, koanf.UnmarshalConf{Tag: "yaml", FlatPaths: true})
}
