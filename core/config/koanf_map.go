package config

import (
	"github.com/starudream/go-lib/core/v2/config/internal/providers/confmap"
)

func LoadMap(m map[string]any) {
	_ = _k.Load(confmap.Provider(m, "."), nil)
}
