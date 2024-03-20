package recovery

import (
	"context"
	"testing"

	"github.com/starudream/go-lib/core/v2/slog"
	"github.com/starudream/go-lib/server/v2/ictx"
	"github.com/starudream/go-lib/server/v2/ierr"
)

func TestUnary(t *testing.T) {
	ctx := slog.WithAttrs(ictx.FromContext(context.Background()), slog.String("foo", "bar"))

	handler := func(ctx context.Context, req any) (any, error) {
		s := []string{"a", "b", "c"}
		_ = s[5:]
		return nil, ierr.InternalServer(0, "unreachable")
	}

	_, _ = Unary()(ctx, nil, nil, handler)
}
