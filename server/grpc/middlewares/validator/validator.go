package validator

import (
	"context"

	"google.golang.org/grpc"
)

func Unary() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		if v, ok := req.(interface{ Validate() error }); ok {
			err = v.Validate()
			if err != nil {
				// todo: wrap error
				return nil, err
			}
		}
		return handler(ctx, req)
	}
}
