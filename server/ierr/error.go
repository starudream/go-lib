package ierr

import (
	"errors"
	"fmt"

	"google.golang.org/grpc/status"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

var (
	UnknownStatus = 500
	UnknownCode   = 9999
)

const (
	DefaultStatus = 200
)

type Error struct {
	status int32

	Code     int32             `json:"code"`
	Message  string            `json:"message,omitempty"`
	Metadata map[string]string `json:"metadata,omitempty"`
}

func New(status, code int, format string, args ...any) *Error {
	if len(args) > 0 {
		format = fmt.Sprintf(format, args...)
	}
	return &Error{status: int32(status), Code: int32(code), Message: format}
}

var _ error = (*Error)(nil)

func (e *Error) Error() string {
	return fmt.Sprintf("error: status=%d, code=%d, message=%s, metadata=%v", e.status, e.Code, e.Message, e.Metadata)
}

var _ fmt.Stringer = (*Error)(nil)

func (e *Error) String() string {
	return e.Error()
}

func (e *Error) GRPCStatus() *status.Status {
	t, _ := status.New(toGRPCCode(int(e.status)), e.Message).WithDetails(&errdetails.ErrorInfo{Metadata: e.Metadata})
	return t
}

func (e *Error) WithMetadata(metadata map[string]string) *Error {
	e.Metadata = metadata
	return e
}

func (e *Error) AppendMetadata(metadata map[string]string) *Error {
	if e.Metadata == nil {
		e.Metadata = map[string]string{}
	}
	for k, v := range metadata {
		e.Metadata[k] = v
	}
	return e
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
		return New(UnknownStatus, UnknownCode, err.Error())
	}
	ne := New(fromGRPCCode(ge.Code()), UnknownCode, ge.Message())
	if dts := ge.Details(); len(dts) > 0 {
		if dt, ok := dts[0].(*errdetails.ErrorInfo); ok {
			ne.Metadata = dt.Metadata
		}
	}
	return ne
}

func Status(err error) int {
	if err == nil {
		return DefaultStatus
	}
	return int(FromError(err).status)
}
