package ierr

import (
	"errors"
	"fmt"

	"google.golang.org/grpc/status"
)

type Error struct {
	Code    int32
	Reason  string
	Message string
}

func New(code int, reason, format string, args ...any) *Error {
	if len(args) > 0 {
		format = fmt.Sprintf(format, args...)
	}
	return &Error{Code: int32(code), Reason: reason, Message: format}
}

var _ error = (*Error)(nil)

func (e *Error) Error() string {
	return fmt.Sprintf("error: code=%d, reason=%s, message=%s", e.Code, e.Reason, e.Message)
}

func (e *Error) GRPCStatus() *status.Status {
	return status.New(toGRPCCode(int(e.Code)), e.Message)
}

func (e *Error) FromError(err error) *Error {
	if err == nil {
		return nil
	}
	if ne := new(Error); errors.As(err, &ne) {
		return ne
	}
	return nil
}
