package sqlite

import (
	"path/filepath"
	"time"

	"github.com/starudream/go-lib/core/v2/config"
	"github.com/starudream/go-lib/core/v2/utils/osutil"
)

type Config struct {
	DSN           string        `json:"sqlite.dsn"            yaml:"sqlite.dsn"`
	SlowThreshold time.Duration `json:"sqlite.slow_threshold" yaml:"sqlite.slow_threshold"`
}

var _c = Config{
	DSN: func() string {
		path, name := osutil.ExeDir(), osutil.ExeName()+".sqlite"
		if osutil.ArgTest() {
			path = filepath.Join(osutil.GoListRoot(), "bin")
		}
		return filepath.Join(path, name)
	}(),
	SlowThreshold: time.Second,
}

var _db *GormDB

func init() {
	_ = config.Unmarshal("", &_c)
	config.LoadStruct(_c)

	db, err := open(_c.DSN)
	osutil.ExitErr(err)
	_db = db
}

func DB() *GormDB {
	return _db
}
