package resty

import (
	"github.com/starudream/go-lib/core/v2/utils/optionutil"
)

type ROptions struct {
	Headers map[string]string
}

type ROption = optionutil.I[ROptions]

func WithAccept(accept string) ROption {
	return optionutil.New(func(opts *ROptions) {
		if accept == "" {
			return
		}
		opts.Headers[HeaderAccept] = accept
	})
}

func WithUserAgent(useragent string) ROption {
	return optionutil.New(func(opts *ROptions) {
		if useragent == "" {
			return
		}
		opts.Headers[HeaderUserAgent] = useragent
	})
}
