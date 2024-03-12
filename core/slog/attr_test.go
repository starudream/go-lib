package slog_test

import (
	"context"
	"testing"

	"github.com/starudream/go-lib/core/v2/slog"
)

func TestAttrs(t *testing.T) {
	ctx1 := context.Background()

	ctx2 := slog.WithAttrs(ctx1,
		// slog.Bool("bool", true),
		// slog.Duration("duration", time.Hour+time.Minute+time.Second),
		// slog.Float64("float64", math.Pi),
		// slog.Int("int", math.MaxInt),
		// slog.Int64("int64", math.MaxInt64),
		slog.String("string", "foo"),
		// slog.Time("time", time.Now()),
		// slog.Uint64("uint64", math.MaxUint64),
	)

	slog.Info("ctx1", slog.GetAttrs(ctx1), slog.String("string", "bar"))
	slog.Info("ctx2", slog.GetAttrs(ctx2), slog.String("string", "bar"))
}
