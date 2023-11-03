package record

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/starudream/go-lib/core/v2/slog/level"
)

type Options struct {
	ctx       context.Context
	logger    *slog.Logger
	time      time.Time
	level     level.Level
	skip      int
	skipNames []string
	msg       string
	attrs     []slog.Attr
}

type OptionI interface {
	apply(*Options)
}

type optionF struct {
	f func(*Options)
}

var _ OptionI = (*optionF)(nil)

func (of *optionF) apply(opts *Options) {
	of.f(opts)
}

func newOptionFunc(f func(*Options)) *optionF {
	return &optionF{f: f}
}

func WithContext(ctx context.Context) OptionI {
	return newOptionFunc(func(o *Options) {
		o.ctx = ctx
	})
}

func WithLogger(logger *slog.Logger) OptionI {
	return newOptionFunc(func(o *Options) {
		o.logger = logger
	})
}

func WithTime(time time.Time) OptionI {
	return newOptionFunc(func(o *Options) {
		o.time = time
	})
}

func WithLevel(level level.Level) OptionI {
	return newOptionFunc(func(o *Options) {
		o.level = level
	})
}

func WithSkip(skip int) OptionI {
	return newOptionFunc(func(o *Options) {
		o.skip = 2 + skip
	})
}

func WithSkipNames(skipNames ...string) OptionI {
	return newOptionFunc(func(o *Options) {
		o.skipNames = skipNames
	})
}

func WithMsg(format string, args ...any) OptionI {
	return newOptionFunc(func(o *Options) {
		if len(args) == 0 {
			o.msg = format
		} else {
			o.msg = fmt.Sprintf(format, args...)
		}
	})
}

func WithAttrs(attrs ...slog.Attr) OptionI {
	return newOptionFunc(func(o *Options) {
		o.attrs = append(o.attrs, attrs...)
	})
}

func WithMsgAndAttrs(format string, argsAndAttrs ...any) OptionI {
	args, attrs := func() (_ []any, attrs []slog.Attr) {
		n := len(argsAndAttrs)
		if n == 0 {
			return nil, nil
		}
		for i := n - 1; i >= 0; i-- {
			if attr, ok := argsAndAttrs[i].(slog.Attr); ok {
				attrs = append(attrs, attr)
				continue
			}
			if i == n-1 {
				return argsAndAttrs, nil
			}
			return argsAndAttrs[:i+1], attrs
		}
		return nil, attrs
	}()
	return newOptionFunc(func(o *Options) {
		if len(args) == 0 {
			o.msg = format
		} else {
			o.msg = fmt.Sprintf(format, args...)
		}
		o.attrs = append(o.attrs, attrs...)
	})
}
