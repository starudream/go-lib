package grpc

import (
	"context"

	"google.golang.org/grpc/metadata"
)

type MD = metadata.MD

func GetMD(ctx context.Context) MD {
	md, _ := metadata.FromIncomingContext(ctx)
	return md.Copy()
}
