package ntfy

import (
	"github.com/starudream/go-lib/core/v2/utils/optionutil"
)

type Options struct {
	extra   map[string]string
	headers map[string]string
}

func newOptions(options ...Option) *Options {
	return optionutil.Build(&Options{
		extra:   map[string]string{},
		headers: map[string]string{},
	}, options...)
}

type Option = optionutil.I[Options]

func WithExtra(extra map[string]string) Option {
	return optionutil.New(func(t *Options) {
		for k, v := range extra {
			t.extra[k] = v
		}
	})
}

func WithHeader(headers map[string]string) Option {
	return optionutil.New(func(t *Options) {
		for k, v := range headers {
			t.headers[k] = v
		}
	})
}
