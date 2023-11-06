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
	_files      = make([]string, _filesMustSize, 20)
	_filesOnce  sync.Once
	_loadedFile string
)

func Files() []string {
	_filesOnce.Do(func() {
		// config file from flag --config or -c
		_files[0] = fileFromFlags()
		_files[1] = os.Getenv("CONFIG")
		_files[2] = os.Getenv("CONFIG_FILE")
		_files[3] = os.Getenv("APP_CONFIG_FILE")

		// test config file
		if root := osutil.GoListRoot(); root != "" && osutil.ArgTest() {
			_files = append(_files, filepathJoinDir(root, "app")...)
			_files = append(_files, filepathJoinDir(root, filepath.Base(root))...)
		}

		if name := os.Getenv("APP_NAME"); name != "" {
			_files = append(_files, filepathJoin(name)...)
		}

		// default config file
		_files = append(_files, filepathJoin(osutil.ExeName())...)
	})
	return _files
}

func LoadedFile() string {
	return _loadedFile
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

func filepathJoinDir(dir, base string) []string {
	return []string{
		filepath.Join(dir, base+".yaml"),
		filepath.Join(dir, base+".yml"),
	}
}

func filepathJoin(base string) []string {
	return []string{
		filepath.Join(osutil.WorkDir(), base+".yaml"),
		filepath.Join(osutil.WorkDir(), base+".yml"),
		filepath.Join(osutil.ExeDir(), base+".yaml"),
		filepath.Join(osutil.ExeDir(), base+".yml"),
	}
}
