package recovery

import (
	"context"

	"google.golang.org/grpc"

	"github.com/starudream/go-lib/core/v2/slog"
	"github.com/starudream/go-lib/core/v2/utils/osutil"
	"github.com/starudream/go-lib/server/v2/otel"
)

func Unary() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		spanCtx := otel.SpanContextFromContext(ctx)
		if spanCtx.IsValid() {
			spanId, traceId := spanCtx.SpanID().String(), spanCtx.TraceID().String()
			ctx = slog.WithAttrs(ctx, slog.String("span_id", spanId), slog.String("trace_id", traceId))
		}
		defer func() {
			if r := recover(); r != nil {
				slog.Error("[%s] %v", osutil.CallerString(2), r, slog.String("stack", osutil.Stack(2)))
			}
			// todo: wrap error
		}()
		return handler(ctx, req)
	}
}
