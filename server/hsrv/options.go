package hsrv

import (
	"net/http"

	"github.com/starudream/go-lib/core/v2/utils/optionutil"
)

type Options struct {
	Addr    string
	Handler http.Handler
}

func newOptions(options ...Option) *Options {
	return optionutil.Build(&Options{Handler: Root()}, options...)
}

type Option = optionutil.I[Options]

func WithAddr(addr string) Option {
	return optionutil.New(func(opts *Options) {
		opts.Addr = addr
	})
}

func WithHandler(handler http.Handler) Option {
	return optionutil.New(func(opts *Options) {
		if handler == nil {
			return
		}
		opts.Handler = handler
	})
}
