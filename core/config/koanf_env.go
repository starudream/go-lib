package config

import (
	"strings"

	"github.com/starudream/go-lib/core/v2/config/internal/providers/env"
)

func LoadEnv() {
	_ = _k.Load(env.ProviderWithValue("", ".", envCB()), nil)
}

func envCB() func(key string, value string) (string, any) {
	prefix := "app."
	return func(key string, value string) (string, any) {
		key = strings.ReplaceAll(strings.ToLower(strings.TrimSpace(key)), "__", ".")
		if !strings.HasPrefix(key, prefix) {
			return "", nil
		}
		key = strings.TrimPrefix(key, prefix)
		return key, value
	}
}
