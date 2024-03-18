package sqlite

import (
	"errors"

	"gorm.io/gorm"

	"modernc.org/sqlite"
	sqlite3 "modernc.org/sqlite/lib"
)

// Translate it will translate the error to native gorm errors.
func (dialector Dialector) Translate(err error) error {
	var se *sqlite.Error
	if errors.As(err, &se) {
		switch se.Code() {
		case sqlite3.SQLITE_CONSTRAINT_UNIQUE, sqlite3.SQLITE_CONSTRAINT_PRIMARYKEY:
			return gorm.ErrDuplicatedKey
		case sqlite3.SQLITE_CONSTRAINT_FOREIGNKEY:
			return gorm.ErrForeignKeyViolated
		}
	}
	return err
}
