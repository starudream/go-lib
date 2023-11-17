package cron

import (
	"time"

	"github.com/robfig/cron/v3"

	"github.com/starudream/go-lib/core/v2/slog"
	"github.com/starudream/go-lib/core/v2/slog/record"
)

var _logger = &logger{}

type logger struct {
}

var _ cron.Logger = (*logger)(nil)

func (l *logger) Info(msg string, kvs ...any) {
	l.Log(slog.LevelInfo, msg, kvs...)
}

func (l *logger) Error(error, string, ...any) {
	panic("unreachable")
}

func (l *logger) E(msg string, kvs ...any) {
	l.Log(slog.LevelError, msg, kvs...)
}

func (l *logger) Log(level slog.Level, msg string, kvs ...any) {
	attrs := []slog.Attr{slog.String("module", "cron")}
	for i := 0; i < len(kvs); i += 2 {
		k, v := kvs[i].(string), kvs[i+1]
		if k == "entry" {
			id := v.(EntryID)
			attrs = append(attrs, slog.Int("id", int(id)))
			name, ok := _cnm.Load(id)
			if ok {
				attrs = append(attrs, slog.String("name", name.(string)))
			}
			v = nil
		} else if k == "now" {
			v = nil
		} else if t, ok := v.(time.Time); ok {
			v = t.Local().Format("2006-01-02T15:04:05.000Z07:00")
		}
		if v == nil {
			continue
		}
		attrs = append(attrs, slog.Any(k, v))
	}
	record.Handle(
		record.WithLevel(level),
		record.WithSkip(1),
		record.WithMsg(msg),
		record.WithAttrs(attrs...),
	)
}
