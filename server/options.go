package server

import (
	"time"

	"github.com/starudream/go-lib/core/v2/utils/optionutil"
)

type Options struct {
	hs Server
	gs Server

	timeout time.Duration
}

func newOptions(options ...Option) *Options {
	return optionutil.Build(&Options{
		timeout: 3 * time.Second,
	}, options...)
}

type Option = optionutil.I[Options]

func WithHTTP(hs Server) Option {
	return optionutil.New(func(o *Options) {
		o.hs = hs
	})
}

func WithGRPC(gs Server) Option {
	return optionutil.New(func(o *Options) {
		o.gs = gs
	})
}

func WithStopTimeout(timeout time.Duration) Option {
	return optionutil.New(func(o *Options) {
		if timeout >= 0 {
			o.timeout = timeout
		}
	})
}
