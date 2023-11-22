package ping

import (
	"time"

	"github.com/starudream/go-lib/core/v2/utils/optionutil"
)

type Options struct {
	addr       string
	interval   time.Duration
	timeout    time.Duration
	count      int
	privileged bool
}

func newOptions(options ...Option) *Options {
	return optionutil.Build(&Options{
		interval: time.Second,
		timeout:  5 * time.Second,
		count:    3,
	}, options...)
}

type Option = optionutil.I[Options]

func WithAddr(addr string) Option {
	return optionutil.New(func(t *Options) {
		t.addr = addr
	})
}

func WithInterval(interval time.Duration) Option {
	return optionutil.New(func(t *Options) {
		t.interval = interval
	})
}

func WithTimeout(timeout time.Duration) Option {
	return optionutil.New(func(t *Options) {
		t.timeout = timeout
	})
}

func WithCount(count int) Option {
	return optionutil.New(func(t *Options) {
		t.count = count
	})
}

func WithPrivileged(privileged bool) Option {
	return optionutil.New(func(t *Options) {
		t.privileged = privileged
	})
}
