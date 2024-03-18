package prepare

import (
	"context"

	"github.com/google/uuid"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/starudream/go-lib/core/v2/slog"
	"github.com/starudream/go-lib/server/v2/iconst"
	"github.com/starudream/go-lib/server/v2/ictx"
	"github.com/starudream/go-lib/server/v2/otel"
)

const (
	keyRequestId = "request-id"
	keyTraceId   = "trace-id"
	keySpanId    = "span-id"
)

func Unary() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		c := ictx.FromContext(ctx)

		reqId := c.Get(iconst.HeaderXRequestID)
		if reqId == "" {
			reqId = "X" + uuid.NewString()[1:]
		}
		_ = grpc.SetHeader(ctx, metadata.Pairs(iconst.HeaderXRequestID, reqId))

		attrs := []slog.Attr{slog.String(keyRequestId, reqId)}

		spanCtx := otel.SpanContextFromContext(ctx)
		if spanCtx.IsValid() {
			traceId, spanId := spanCtx.TraceID().String(), spanCtx.SpanID().String()
			attrs = append(attrs, slog.String(keyTraceId, traceId), slog.String(keySpanId, spanId))
		}

		ctx = slog.WithAttrs(ctx, attrs...)

		return handler(ctx, req)
	}
}
