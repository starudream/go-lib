package ntfy

import (
	"context"
	"fmt"

	"github.com/starudream/go-lib/core/v2/config"
)

type Config struct {
	DingtalkConfig `yaml:",squash"`
	TelegramConfig `yaml:",squash"`
}

var _c = Config{}

func init() {
	_ = config.Unmarshal("", &_c)
}

func C() Config {
	return _c
}

var ErrNoConfig = fmt.Errorf("no notify config")

func Notify(ctx context.Context, text string) (err error) {
	var c Interface
	switch {
	case _c.DingtalkConfig.Token != nil:
		c = _c.DingtalkConfig
	case _c.TelegramConfig.Token != nil:
		c = _c.TelegramConfig
	}
	if c != nil {
		return c.Notify(ctx, text)
	}
	return ErrNoConfig
}
