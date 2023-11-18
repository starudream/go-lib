package record

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/starudream/go-lib/core/v2/slog/level"
	"github.com/starudream/go-lib/core/v2/utils/optionutil"
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

type Option = optionutil.I[Options]

func WithContext(ctx context.Context) Option {
	return optionutil.New(func(o *Options) {
		o.ctx = ctx
	})
}

func WithLogger(logger *slog.Logger) Option {
	return optionutil.New(func(o *Options) {
		o.logger = logger
	})
}

func WithTime(time time.Time) Option {
	return optionutil.New(func(o *Options) {
		o.time = time
	})
}

func WithLevel(level level.Level) Option {
	return optionutil.New(func(o *Options) {
		o.level = level
	})
}

func WithSkip(skip int) Option {
	return optionutil.New(func(o *Options) {
		o.skip = 2 + skip
	})
}

func WithSkipNames(skipNames ...string) Option {
	return optionutil.New(func(o *Options) {
		o.skipNames = skipNames
	})
}

func WithMsg(format string, args ...any) Option {
	return optionutil.New(func(o *Options) {
		if len(args) == 0 {
			o.msg = format
		} else {
			o.msg = fmt.Sprintf(format, args...)
		}
	})
}

func WithAttrs(attrs ...slog.Attr) Option {
	return optionutil.New(func(o *Options) {
		o.attrs = append(o.attrs, attrs...)
	})
}

func WithMsgAndAttrs(format string, argsAndAttrs ...any) Option {
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
		return nil, reverse(attrs)
	}()
	return optionutil.New(func(o *Options) {
		if len(args) == 0 {
			o.msg = format
		} else {
			o.msg = fmt.Sprintf(format, args...)
		}
		o.attrs = append(o.attrs, attrs...)
	})
}

func reverse[T any](vs []T) []T {
	length := len(vs)
	half := length / 2
	for i := 0; i < half; i++ {
		j := length - 1 - i
		vs[i], vs[j] = vs[j], vs[i]
	}
	return vs
}
