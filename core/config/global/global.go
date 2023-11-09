package global

import (
	"path/filepath"

	"github.com/starudream/go-lib/core/v2/config"
	"github.com/starudream/go-lib/core/v2/slog/level"
	"github.com/starudream/go-lib/core/v2/utils/osutil"
)

type Config struct {
	LogConsoleDisabled bool        `json:"log.console.disabled,omitempty" yaml:"log.console.disabled,omitempty"`
	LogConsoleFormat   string      `json:"log.console.format,omitempty"   yaml:"log.console.format,omitempty"  `
	LogConsoleLevel    level.Level `json:"log.console.level"              yaml:"log.console.level"             `
	LogFileEnabled     bool        `json:"log.file.enabled"               yaml:"log.file.enabled"              `
	LogFileFormat      string      `json:"log.file.format,omitempty"      yaml:"log.file.format,omitempty"     `
	LogFileLevel       level.Level `json:"log.file.level"                 yaml:"log.file.level"                `
	LogFileFilename    string      `json:"log.file.filename"              yaml:"log.file.filename"             `
	LogFileMaxSize     int         `json:"log.file.max_size,omitempty"    yaml:"log.file.max_size,omitempty"   `
	LogFileMaxAge      int         `json:"log.file.max_age,omitempty"     yaml:"log.file.max_age,omitempty"    `
	LogFileMaxBackups  int         `json:"log.file.max_backups,omitempty" yaml:"log.file.max_backups,omitempty"`
	LogFileDailyRotate bool        `json:"log.file.daily_rotate"          yaml:"log.file.daily_rotate"         `
}

var _c = Config{
	LogConsoleDisabled: false,
	LogConsoleFormat:   "text",
	LogConsoleLevel:    level.Info,
	LogFileEnabled:     false,
	LogFileFormat:      "text",
	LogFileLevel:       level.Info,
	LogFileFilename: func() string {
		path, name := osutil.ExeDir(), osutil.ExeName()+".log"
		if osutil.ArgTest() {
			path = filepath.Join(osutil.GoListRoot(), "log")
		}
		return filepath.Join(path, name)
	}(),
	LogFileMaxSize:     100,
	LogFileMaxAge:      0,
	LogFileMaxBackups:  10,
	LogFileDailyRotate: true,
}

func init() {
	_ = config.Unmarshal("", &_c)
	config.LoadStruct(_c)
}

func C() Config {
	return _c
}
