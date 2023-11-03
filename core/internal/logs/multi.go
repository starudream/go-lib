package logs

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/samber/lo"

	"github.com/starudream/go-lib/core/v2/utils/osutil"
)

type MultiHandler struct {
	handlers []slog.Handler
}

func NewMultiHandler(handlers ...slog.Handler) *MultiHandler {
	return &MultiHandler{lo.Filter(handlers, func(h slog.Handler, _ int) bool { return h != nil })}
}

var _ slog.Handler = (*MultiHandler)(nil)

func (h *MultiHandler) Enabled(ctx context.Context, level slog.Level) bool {
	for i := range h.handlers {
		if h.handlers[i].Enabled(ctx, level) {
			return true
		}
	}
	return false
}

func (h *MultiHandler) Handle(ctx context.Context, record slog.Record) error {
	for i := range h.handlers {
		if h.handlers[i].Enabled(ctx, record.Level) {
			err := h.try(func() error { return h.handlers[i].Handle(ctx, record) })
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (h *MultiHandler) try(fn func() error) (err error) {
	defer func() {
		if r := recover(); r != nil && err == nil {
			err = fmt.Errorf("[%s] %v", osutil.CallerString(1), r)
		}
	}()
	err = fn()
	return
}

func (h *MultiHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &MultiHandler{lo.Map(h.handlers, func(h slog.Handler, _ int) slog.Handler { return h.WithAttrs(attrs) })}
}

func (h *MultiHandler) WithGroup(name string) slog.Handler {
	return &MultiHandler{lo.Map(h.handlers, func(h slog.Handler, _ int) slog.Handler { return h.WithGroup(name) })}
}
