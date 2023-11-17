package service

import (
	"github.com/starudream/go-lib/core/v2/slog"
	"github.com/starudream/go-lib/core/v2/slog/record"
)

func log(level slog.Level, msg string) {
	record.Handle(
		record.WithLevel(level),
		record.WithSkip(1),
		record.WithMsg(msg),
		record.WithAttrs(slog.String("module", "service")),
	)
}
