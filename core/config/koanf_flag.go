package config

import (
	"strings"

	"github.com/spf13/pflag"

	"github.com/starudream/go-lib/core/v2/config/internal/providers/posflag"
)

func LoadFlags(fs *pflag.FlagSet) {
	if fs == nil {
		return
	}
	_ = _k.Load(posflag.ProviderWithFlag(fs, ".", _k, flagCB(fs)), nil)
}

func flagCB(fs *pflag.FlagSet) func(f *pflag.Flag) (string, any) {
	return func(f *pflag.Flag) (string, any) {
		return strings.ReplaceAll(strings.ToLower(strings.TrimSpace(f.Name)), "-", "."), posflag.FlagVal(fs, f)
	}
}
