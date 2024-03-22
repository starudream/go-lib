package slog

import (
	"context"
	"log/slog"
)

type Attr = slog.Attr

var (
	Group = slog.Group

	Any = slog.Any

	Bool     = slog.Bool
	Duration = slog.Duration
	Float64  = slog.Float64
	Int      = slog.Int
	Int64    = slog.Int64
	String   = slog.String
	Time     = slog.Time
	Uint64   = slog.Uint64
)

type ctxkeyAttr struct{}

func WithAttrs(ctx context.Context, attrs ...Attr) context.Context {
	attrs = append(GetAttrs(ctx), attrs...)
	ctx = context.WithValue(ctx, ctxkeyAttr{}, attrs)
	return ctx
}

func GetAttrs(ctx context.Context) []Attr {
	attrs, _ := ctx.Value(ctxkeyAttr{}).([]Attr)
	return attrs
}
