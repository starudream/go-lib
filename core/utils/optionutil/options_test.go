package optionutil

import (
	"fmt"
	"testing"
)

type Options struct {
	k1 string
	k2 bool
}

type Option = I[Options]

func WithK1(k1 string) Option {
	return New(func(opts *Options) {
		opts.k1 = k1
	})
}

func WithK2(k2 bool) Option {
	return New(func(opts *Options) {
		opts.k2 = k2
	})
}

func T(opts ...Option) {
	v := Build(&Options{}, opts...)
	fmt.Printf("%+v\n", v)
}

func TestNew(t *testing.T) {
	T(WithK1("foo"), WithK2(true))
	T(WithK2(true))
}
