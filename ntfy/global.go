package ntfy

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/starudream/go-lib/core/v2/config"
	"github.com/starudream/go-lib/resty/v2"
)

type Config struct {
	Timeout time.Duration `json:"ntfy.timeout"  yaml:"ntfy.timeout"`

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

var (
	_cli     *resty.Client
	_cliOnce sync.Once
)

func R() *resty.Request {
	_cliOnce.Do(func() {
		_cli = resty.New()
	})
	return _cli.R()
}
