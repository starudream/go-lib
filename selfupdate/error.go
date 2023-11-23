package selfupdate

import (
	"errors"
)

// RollbackError takes an error value returned by Apply and returns the error, if any,
// that occurred when attempting to roll back from a failed update. Applications should
// always call this function on any non-nil errors returned by Apply.
//
// If no rollback was needed or if the rollback was successful, RollbackError returns nil,
// otherwise it returns the error encountered when trying to roll back.
func RollbackError(err error) (error, bool) {
	var e *rollbackErr
	if errors.As(err, &e) {
		return e.rollbackErr, true
	}
	return nil, false
}

type rollbackErr struct {
	error             // original error
	rollbackErr error // error encountered while rolling back
}
