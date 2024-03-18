package recovery

import (
	"context"
	"testing"

	"github.com/starudream/go-lib/core/v2/slog"
	"github.com/starudream/go-lib/server/v2/ictx"
)

func TestUnary(t *testing.T) {
	ctx := slog.WithAttrs(ictx.FromContext(context.Background()), slog.String("foo", "bar"))

	handler := func(ctx context.Context, req any) (any, error) {
		panic("hello")
		return nil, nil
	}

	_, _ = Unary()(ctx, nil, nil, handler)
}
