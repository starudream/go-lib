package otelgrpc

import (
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
)

var (
	NewClientHandler = otelgrpc.NewClientHandler
	NewServerHandler = otelgrpc.NewServerHandler
)
