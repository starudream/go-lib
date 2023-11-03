package config

import (
	"os"
	"path/filepath"
	"sync"

	"github.com/spf13/pflag"

	"github.com/starudream/go-lib/core/v2/utils/osutil"

	"github.com/starudream/go-lib/core/v2/config/internal/parsers/yaml"
	"github.com/starudream/go-lib/core/v2/config/internal/providers/file"
)

func LoadFile(name string) error {
	if name == "" {
		return nil
	}
	return _k.Load(file.Provider(name), yaml.Parser())
}

var (
	_filesMustSize = 10
	// files[0:_filesMustSize] must be success if not empty
	_files     = make([]string, _filesMustSize, 20)
	_filesOnce sync.Once
)

func Files() []string {
	_filesOnce.Do(func() {
		// config file from flag --config or -c
		_files[0] = fileFromFlags()
		_files[1] = os.Getenv("CONFIG")
		_files[2] = os.Getenv("CONFIG_FILE")
		_files[3] = os.Getenv("APP_CONFIG_FILE")

		// test config file
		if osutil.ArgTest() {
			_files = append(_files,
				filepath.Join(osutil.GoListRoot(), "app.yaml"),
				filepath.Join(osutil.GoListRoot(), "../app.yaml"),
				filepath.Join(osutil.GoListRoot(), "../bin/app.yaml"),
			)
		}

		// default config file
		_files = append(_files,
			filepath.Join(osutil.WorkDir(), osutil.ExeName()+".yaml"),
			filepath.Join(osutil.WorkDir(), osutil.ExeName()+".yml"),
			filepath.Join(osutil.ExeDir(), osutil.ExeName()+".yaml"),
			filepath.Join(osutil.ExeDir(), osutil.ExeName()+".yml"),
		)
	})
	return _files
}

func fileFromFlags() (name string) {
	fs := pflag.NewFlagSet("", pflag.ContinueOnError)
	fs.StringVarP(&name, "config", "c", "", "")
	_ = fs.Parse(os.Args[1:])
	if name != "" {
		name, _ = filepath.Abs(name)
	}
	return
}
