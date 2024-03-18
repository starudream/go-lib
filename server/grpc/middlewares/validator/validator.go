package validator

import (
	"context"

	"google.golang.org/grpc"

	"github.com/starudream/go-lib/server/v2/ierr"
)

func Unary() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		if v, ok := req.(interface{ Validate() error }); ok {
			err = v.Validate()
			if err != nil {
				return nil, ierr.BadRequest("validate error", err.Error())
			}
		}
		return handler(ctx, req)
	}
}
