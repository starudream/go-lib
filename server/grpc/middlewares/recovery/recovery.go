package recovery

import (
	"context"

	"google.golang.org/grpc"

	"github.com/starudream/go-lib/core/v2/slog"
	"github.com/starudream/go-lib/core/v2/utils/osutil"
	"github.com/starudream/go-lib/server/v2/ierr"
)

func Unary() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ any, err error) {
		defer func() {
			if r := recover(); r != nil {
				slog.Error("[%s] %v", osutil.CallerString(2), r, slog.String("stack", osutil.Stack(2)), slog.GetAttrs(ctx))
				err = ierr.InternalServer(9999, "%v", r)
			}
		}()
		return handler(ctx, req)
	}
}
