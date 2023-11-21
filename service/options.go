package service

import (
	"github.com/starudream/go-lib/core/v2/utils/optionutil"
)

type Options struct {
	displayName string
	description string
	arguments   []string
	envVars     map[string]string
}

func newOptions(options ...Option) *Options {
	return optionutil.Build(&Options{}, options...)
}

type Option = optionutil.I[Options]

func WithDisplayName(displayName string) Option {
	return optionutil.New(func(t *Options) {
		t.displayName = displayName
	})
}

func WithDescription(description string) Option {
	return optionutil.New(func(t *Options) {
		t.description = description
	})
}

func WithArguments(arguments ...string) Option {
	return optionutil.New(func(t *Options) {
		t.arguments = arguments
	})
}

func WithEnvVars(envVars map[string]string) Option {
	return optionutil.New(func(t *Options) {
		t.envVars = envVars
	})
}
