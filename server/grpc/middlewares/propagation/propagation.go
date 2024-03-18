package propagation

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func Unary() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		md, _ := metadata.FromIncomingContext(ctx)
		md = md.Copy()
		ctx = metadata.NewOutgoingContext(ctx, md)
		return handler(ctx, req)
	}
}
