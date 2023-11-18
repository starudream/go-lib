package sqlite

import (
	"gorm.io/gorm"

	"github.com/starudream/go-lib/core/v2/slog"

	sqlite "github.com/starudream/go-lib/sqlite/v2/internal/driver"
)

type GormDB = gorm.DB

func open(dsn string) (*GormDB, error) {
	dial, opts :=
		sqlite.Open(dsn),
		&gorm.Config{
			Logger: _logger,

			PrepareStmt:       true,
			AllowGlobalUpdate: true,
			TranslateError:    true,

			DisableForeignKeyConstraintWhenMigrating: true,
		}
	db, err := gorm.Open(dial, opts)
	if err != nil {
		return nil, err
	}

	raw, err := db.DB()
	if err != nil {
		return nil, err
	}
	raw.SetMaxOpenConns(1)

	_logger.D("sqlite initialized", slog.String("dsn", dsn))

	return db, nil
}
