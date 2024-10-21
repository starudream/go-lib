package logger

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"

	"github.com/starudream/go-lib/core/v2/codec/json"
	"github.com/starudream/go-lib/core/v2/slog"
	"github.com/starudream/go-lib/server/v2/ierr"
	"github.com/starudream/go-lib/server/v2/jwt"

	"github.com/starudream/go-lib/server/v2/grpc/internal/annotation"
	"github.com/starudream/go-lib/server/v2/grpc/internal/fieldmask"
)

var marshalOptions = protojson.MarshalOptions{
	UseEnumNumbers:  true,
	EmitUnpopulated: true,
	UseProtoNames:   false,
}

func marshal(v any, paths []string) string {
	switch x := v.(type) {
	case proto.Message:
		if len(paths) > 0 {
			x = proto.Clone(x)
			fieldmask.Prune(x, paths)
		}
		bs, _ := marshalOptions.Marshal(x)
		return string(bs)
	case *ierr.Error:
		bs, _ := json.Marshal(x)
		return string(bs)
	}
	return fmt.Sprintf("(%T)", v)
}

func Unary() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		attrs := append(slog.GetAttrs(ctx),
			slog.String("kind", "server"),
			slog.String("method", info.FullMethod),
		)

		claims, _ := jwt.FromContext(ctx)
		if claims != nil {
			attrs = append(attrs,
				slog.String("jwt-issuer", claims.ISS()),
				slog.String("jwt-subject", claims.SUB()),
				slog.String("jwt-audience", claims.AUD()),
			)
		}

		slog.Info("req: %s", marshal(req, annotation.GetMethodOptions(info.FullMethod).GetReqMaskPaths()), attrs)

		defer func(start time.Time) {
			attrs = append(attrs, slog.Duration("took", time.Since(start)))

			if err != nil {
				err = ierr.FromError(err)
				slog.Error("resp: %s", marshal(err, nil), attrs)
			} else {
				slog.Info("resp: %s", marshal(resp, annotation.GetMethodOptions(info.FullMethod).GetRespMaskPaths()), attrs)
			}
		}(time.Now())

		ctx = slog.WithAttrs(ctx, attrs...)

		return handler(ctx, req)
	}
}

func UnaryClient() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, resp any, conn *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) (err error) {
		attrs := append(slog.GetAttrs(ctx),
			slog.String("kind", "client"),
			slog.String("method", method),
		)

		slog.Info("req: %s", marshal(req, annotation.GetMethodOptions(method).GetReqMaskPaths()), attrs)

		defer func(start time.Time) {
			attrs = append(attrs, slog.Duration("took", time.Since(start)))

			if err != nil {
				err = ierr.FromError(err)
				slog.Error("resp: %s", marshal(err, nil), attrs)
			} else {
				slog.Info("resp: %s", marshal(resp, annotation.GetMethodOptions(method).GetRespMaskPaths()), attrs)
			}
		}(time.Now())

		return invoker(ctx, method, req, resp, conn, opts...)
	}
}
