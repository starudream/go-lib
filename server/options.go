package server

import (
	"github.com/starudream/go-lib/core/v2/utils/optionutil"
)

type Options struct {
	hs Server
	gs Server
}

func newOptions(options ...Option) *Options {
	return optionutil.Build(&Options{}, options...)
}

type Option = optionutil.I[Options]

func WithHTTP(hs Server) Option {
	return optionutil.New(func(opts *Options) {
		opts.hs = hs
	})
}

func WithGRPC(gs Server) Option {
	return optionutil.New(func(opts *Options) {
		opts.gs = gs
	})
}
