package slog

import (
	"context"
	"log/slog"

	"github.com/starudream/go-lib/core/v2/config/global"
	"github.com/starudream/go-lib/core/v2/slog/handler"
	"github.com/starudream/go-lib/core/v2/slog/record"

	"github.com/starudream/go-lib/core/v2/internal/logs"
)

type Logger = slog.Logger

func init() {
	SetDefault(New(
		logs.NewMultiHandler(
			handler.NewConsole(global.C()),
			handler.NewFile(global.C()),
		),
	))
}

var (
	New        = logs.New
	SetDefault = logs.SetDefault
)

func Debug(format string, argsAndAttrs ...any) {
	log(context.Background(), LevelDebug, format, argsAndAttrs...)
}

func Info(format string, argsAndAttrs ...any) {
	log(context.Background(), LevelInfo, format, argsAndAttrs...)
}

func Warn(format string, argsAndAttrs ...any) {
	log(context.Background(), LevelWarn, format, argsAndAttrs...)
}

func Error(format string, argsAndAttrs ...any) {
	log(context.Background(), LevelError, format, argsAndAttrs...)
}

func Fatal(format string, argsAndAttrs ...any) {
	log(context.Background(), LevelFatal, format, argsAndAttrs...)
}

func Log(ctx context.Context, level Level, format string, argsAndAttrs ...any) {
	log(ctx, level, format, argsAndAttrs...)
}

func log(ctx context.Context, level Level, format string, argsAndAttrs ...any) {
	record.Handle(
		record.WithContext(ctx),
		record.WithLevel(level),
		record.WithSkip(1),
		record.WithMsgAndAttrs(format, argsAndAttrs...),
	)
}
