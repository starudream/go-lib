package validator

import (
	"context"

	"google.golang.org/grpc"

	"github.com/starudream/go-lib/server/v2/ierr"
)

type Validator interface {
	Validate() error
}

type ValidateError interface {
	Field() string
	Reason() string
}

func Unary() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		if v, ok1 := req.(Validator); ok1 {
			err := v.Validate()
			if err != nil {
				return nil, ierr.BadRequest(9999, err.Error())
			}
		}
		return handler(ctx, req)
	}
}
