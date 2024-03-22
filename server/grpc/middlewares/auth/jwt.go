package auth

import (
	"context"
	"strings"

	"google.golang.org/grpc"

	"github.com/starudream/go-lib/server/v2/iconst"
	"github.com/starudream/go-lib/server/v2/ictx"
	"github.com/starudream/go-lib/server/v2/ierr"
	"github.com/starudream/go-lib/server/v2/jwt"

	"github.com/starudream/go-lib/server/v2/grpc/internal/annotation"
)

func Unary() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		claims, err := parse(ctx)
		if err != nil && !annotation.GetMethodOptions(info.FullMethod).GetSkipAuth() {
			return nil, err
		}

		ctx = claims.WithContext(ctx)

		return handler(ctx, req)
	}
}

func parse(ctx context.Context) (*jwt.Claims, error) {
	raw := ictx.Get(ctx, iconst.HeaderAuthorization)
	if raw == "" {
		return nil, ierr.Unauthorized(9999, "missing authorization header")
	}
	claims, err := jwt.Parse(strings.TrimPrefix(raw, "Bearer "))
	if err != nil {
		return nil, ierr.Unauthorized(9998, "parse authorization header error: %v", err)
	}
	return claims, nil
}
