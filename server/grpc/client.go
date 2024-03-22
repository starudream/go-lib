package grpc

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/starudream/go-lib/server/v2/otel/otelgrpc"
)

type (
	DialOption = grpc.DialOption
	ClientConn = grpc.ClientConn
)

func Dial(target string, opts ...DialOption) (*ClientConn, error) {
	return grpc.Dial(target, append([]DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(64 * 1024 * 1024)),
		grpc.WithStatsHandler(otelgrpc.NewClientHandler()),
	}, opts...)...)
}
