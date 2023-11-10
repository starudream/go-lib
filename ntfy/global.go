package ntfy

import (
	"context"
	"fmt"
	"net/url"
	"sync"
	"time"

	"github.com/starudream/go-lib/core/v2/config"
	"github.com/starudream/go-lib/core/v2/slog"
	"github.com/starudream/go-lib/resty/v2"
)

type Config struct {
	Timeout time.Duration `json:"ntfy.timeout" yaml:"ntfy.timeout"`
	Proxy   string        `json:"ntfy.proxy"   yaml:"ntfy.proxy"`
	Retry   int           `json:"ntfy.retry"   yaml:"ntfy.retry"`

	DingtalkConfig   `yaml:",squash"`
	TelegramConfig   `yaml:",squash"`
	WeixinWorkConfig `yaml:",squash"`

	WebhookConfig `yaml:",squash"`
}

var _c = Config{}

func init() {
	_ = config.Unmarshal("", &_c)

	_ = config.Unmarshal("ntfy.webhook.extra", &_c.WebhookConfig.Extra)
	_ = config.Unmarshal("ntfy.webhook.header", &_c.WebhookConfig.Header)

	if _c.Timeout <= 0 {
		_c.Timeout = 10 * time.Second
	}
	if _c.Proxy != "" {
		u, err := url.Parse(_c.Proxy)
		if err != nil {
			slog.Fatal("[ntfy] proxy config error: %v", err)
		}
		slog.Info("[ntfy] proxy: %s", u.String())
	}
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
	case _c.WebhookConfig.URL != nil:
		c = _c.WebhookConfig
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
		_cli.SetTimeout(_c.Timeout)
		if _c.Proxy != "" {
			_cli.SetProxy(_c.Proxy)
		}
		if _c.Retry > 0 {
			_cli.SetRetryCount(_c.Retry)
		}
	})
	return _cli.R()
}
