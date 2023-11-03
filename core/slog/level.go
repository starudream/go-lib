package slog

import (
	"github.com/starudream/go-lib/core/v2/slog/level"
)

type Level = level.Level

const (
	LevelDebug = level.Debug
	LevelInfo  = level.Info
	LevelWarn  = level.Warn
	LevelError = level.Error
	LevelFatal = level.Fatal
)
