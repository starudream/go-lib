package ierr

import (
	"errors"
	"fmt"

	"google.golang.org/grpc/status"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

var (
	UnknownCode   = 500
	UnknownReason = ""
)

const (
	DefaultOKCode = 200
)

type Error struct {
	Code    int32  `json:"code"`
	Reason  string `json:"reason,omitempty"`
	Message string `json:"message,omitempty"`
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

var _ fmt.Stringer = (*Error)(nil)

func (e *Error) String() string {
	return e.Error()
}

func (e *Error) GRPCStatus() *status.Status {
	t, _ := status.New(toGRPCCode(int(e.Code)), e.Message).WithDetails(&errdetails.ErrorInfo{Reason: e.Reason})
	return t
}

func (e *Error) AppendMessage(format string, args ...any) *Error {
	if len(args) > 0 {
		format = fmt.Sprintf(format, args...)
	}
	if e.Message != "" {
		e.Message += "\n"
	}
	e.Message += format
	return e
}

func FromError(err error) *Error {
	if err == nil {
		return nil
	}
	if ne := new(Error); errors.As(err, &ne) {
		return ne
	}
	ge, ok := status.FromError(err)
	if !ok {
		return New(UnknownCode, UnknownReason, err.Error())
	}
	ne := New(fromGRPCCode(ge.Code()), UnknownReason, ge.Message())
	if dts := ge.Details(); len(dts) > 0 {
		if dt, ok := dts[0].(*errdetails.ErrorInfo); ok {
			ne.Reason = dt.Reason
		}
	}
	return ne
}

func Code(err error) int {
	if err == nil {
		return DefaultOKCode
	}
	return int(FromError(err).Code)
}

func Reason(err error) string {
	if err == nil {
		return UnknownReason
	}
	return FromError(err).Reason
}
