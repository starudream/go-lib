package config

import (
	"github.com/starudream/go-lib/core/v2/config/internal/providers/structs"
)

func LoadStruct(v any, tag ...string) {
	if len(tag) == 0 || tag[0] == "" {
		tag = []string{"yaml"}
	}
	_ = _k.Load(structs.ProviderWithDelim(v, tag[0], "."), nil)
}
