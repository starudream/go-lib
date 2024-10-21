package prepare

import (
	"context"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/google/uuid"

	"github.com/starudream/go-lib/core/v2/slog"
	"github.com/starudream/go-lib/server/v2/iconst"
	"github.com/starudream/go-lib/server/v2/ictx"
)

func Unary() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		c := ictx.FromContext(ctx)

		reqId := c.Get(iconst.HeaderXRequestID)
		if reqId == "" {
			reqId = "x" + uuid.Must(uuid.NewV7()).String()[1:]
		}
		_ = grpc.SetHeader(ctx, metadata.Pairs(iconst.HeaderXRequestID, reqId))

		attrs := append(slog.GetAttrs(ctx), slog.String("x-request-id", reqId))

		ip := c.Get(iconst.HeaderXRealIP)
		if ip == "" {
			ff := c.Get(iconst.HeaderXForwardedFor)
			if ff != "" {
				ip = strings.TrimSpace(strings.Split(ff, ",")[0])
			}
		}
		if ip != "" {
			attrs = append(attrs, slog.String("client-ip", ip))
		}

		ua := c.Get("V-"+iconst.HeaderUserAgent, iconst.HeaderUserAgent)
		if ua != "" {
			attrs = append(attrs, slog.String("user-agent", ua))
		}

		ctx = slog.WithAttrs(ctx, attrs...)

		return handler(ctx, req)
	}
}
