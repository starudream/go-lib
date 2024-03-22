package jwt

import (
	"github.com/starudream/go-lib/core/v2/utils/optionutil"
)

type Option = optionutil.I[claims]

func WithId(id string) Option {
	return optionutil.New(func(c *claims) {
		if id != "" {
			c.Id = id
		}
	})
}

func WithMetadata(md map[string]string) Option {
	return optionutil.New(func(c *claims) {
		for k, v := range md {
			if k != "" {
				c.Metadata[k] = v
			}
		}
	})
}
