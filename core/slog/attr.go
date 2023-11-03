package slog

import (
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
