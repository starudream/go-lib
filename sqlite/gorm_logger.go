package sqlite

import (
	"context"
	"time"

	gormLogger "gorm.io/gorm/logger"

	"github.com/starudream/go-lib/core/v2/slog"
	"github.com/starudream/go-lib/core/v2/slog/record"
)

var _logger = &logger{}

type logger struct {
}

var _ gormLogger.Interface = (*logger)(nil)

func (l *logger) LogMode(gormLogger.LogLevel) gormLogger.Interface {
	return _logger
}

func (l *logger) Info(ctx context.Context, format string, args ...any) {
	l.Log(ctx, slog.LevelInfo, format, args...)
}

func (l *logger) Warn(ctx context.Context, format string, args ...any) {
	l.Log(ctx, slog.LevelWarn, format, args...)
}

func (l *logger) Error(ctx context.Context, format string, args ...any) {
	l.Log(ctx, slog.LevelError, format, args...)
}

func (l *logger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rows int64), err error) {
	elapsed := time.Since(begin)

	sql, rows := fc()

	attrs := []any{
		slog.String("sql", sql),
		slog.Duration("took", elapsed),
	}
	if rows >= 0 {
		attrs = append(attrs, slog.Int64("rows", rows))
	}

	switch {
	case err != nil && !IsErrNoRows(err):
		l.Log(ctx, slog.LevelError, "sql error: "+err.Error(), attrs...)
	case _c.SlowThreshold > 0 && _c.SlowThreshold < elapsed:
		l.Log(ctx, slog.LevelWarn, "sql slow", attrs...)
	default:
		l.Log(ctx, slog.LevelDebug, "sql", attrs...)
	}
}

func (l *logger) D(format string, argsAndAttrs ...any) {
	l.Log(context.Background(), slog.LevelDebug, format, argsAndAttrs...)
}

func (l *logger) Log(ctx context.Context, level slog.Level, format string, argsAndAttrs ...any) {
	record.Handle(
		record.WithLevel(level),
		record.WithSkip(1),
		record.WithSkipNames("@"),
		record.WithContext(ctx),
		record.WithMsgAndAttrs(format, argsAndAttrs...),
		record.WithAttrs(slog.String("module", "sqlite")),
	)
}
