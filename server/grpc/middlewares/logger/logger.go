package logger

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"

	"github.com/starudream/go-lib/core/v2/slog"
)

var marshalOptions = protojson.MarshalOptions{
	UseEnumNumbers:  true,
	EmitUnpopulated: true,
}

func marshal(v any) string {
	if m, ok := v.(proto.Message); ok {
		bs, _ := marshalOptions.Marshal(m)
		return string(bs)
	}
	return fmt.Sprintf("(%T)", v)
}

func Unary() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		attrs := slog.GetAttrs(ctx)

		slog.Info("req: %s", marshal(req), attrs)

		defer func(start time.Time) {
			attrs = append(attrs, slog.Duration("took", time.Since(start)))

			if err != nil {
				slog.Error("resp: %v", err, attrs)
			} else {
				slog.Info("resp: %s", marshal(resp), attrs)
			}
		}(time.Now())

		return handler(ctx, req)
	}
}
