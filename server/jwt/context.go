package jwt

import (
	"context"

	"github.com/starudream/go-lib/core/v2/utils/osutil"
	"github.com/starudream/go-lib/server/v2/ierr"
)

type ctxkey struct{}

func (c *claims) WithContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, ctxkey{}, c)
}

func FromContext(ctx context.Context) (Interface, error) {
	if c, ok := ctx.Value(ctxkey{}).(*claims); ok {
		return c, nil
	}
	return nil, ierr.Unauthorized(9999, "missing token")
}

func MustFromContext(ctx context.Context) Interface {
	return osutil.Must1(FromContext(ctx))
}
