package resty

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"

	"github.com/starudream/go-lib/core/v2/slog"
	"github.com/starudream/go-lib/core/v2/slog/record"
)

type logger struct {
}

var _ resty.Logger = (*logger)(nil)

func (l *logger) Debugf(format string, args ...any) {
	l.Log(slog.LevelDebug, format, args...)
}

func (l *logger) Warnf(format string, args ...any) {
	l.Log(slog.LevelWarn, format, args...)
}

func (l *logger) Errorf(format string, args ...any) {
	l.Log(slog.LevelError, format, args...)
}

func (l *logger) Log(level slog.Level, format string, args ...any) {
	record.Handle(
		record.WithLevel(level),
		record.WithSkipNames("@", "record/record.go", "resty/logger.go"),
		record.WithMsg(strings.TrimSuffix(fmt.Sprintf(format, args...), "\n")),
		record.WithAttrs(slog.String("module", "resty")),
	)
}
