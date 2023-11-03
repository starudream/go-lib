package resty

type rOptions struct {
	Headers map[string]string
}

type rOptionI interface {
	apply(*rOptions)
}

type rOptionF struct {
	f func(*rOptions)
}

var _ rOptionI = (*rOptionF)(nil)

func (of rOptionF) apply(opts *rOptions) {
	of.f(opts)
}

func newROptionFunc(f func(*rOptions)) *rOptionF {
	return &rOptionF{f: f}
}

//goland:noinspection GoExportedFuncWithUnexportedType
func WithAccept(accept string) rOptionI {
	return newROptionFunc(func(opts *rOptions) {
		if accept == "" {
			return
		}
		opts.Headers[HeaderAccept] = accept
	})
}

//goland:noinspection GoExportedFuncWithUnexportedType
func WithUserAgent(useragent string) rOptionI {
	return newROptionFunc(func(opts *rOptions) {
		if useragent == "" {
			return
		}
		opts.Headers[HeaderUserAgent] = useragent
	})
}
