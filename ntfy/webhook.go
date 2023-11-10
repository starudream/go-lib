package ntfy

import (
	"context"
	"fmt"

	"github.com/starudream/go-lib/core/v2/codec/json"
	"github.com/starudream/go-lib/resty/v2"
)

// WebhookConfig is the config for webhook.
//
// # ServerChan
//
//	ntfy:
//	  webhook:
//	    url: https://sctapi.ftqq.com/key.send
//	    method: post
//	    type: json
//	    key: desp
//	    extras:
//	    title: test
type WebhookConfig struct {
	URL *string `json:"ntfy.webhook.url" yaml:"ntfy.webhook.url"`
	// optional: GET or POST, default is GET
	Method string `json:"ntfy.webhook.method" yaml:"ntfy.webhook.method"`
	// only available when method is POST
	// optional: JSON or FORM, default is FORM
	//  - FORM is request by `application/x-www-form-urlencoded`
	//  - JSON is request by `application/json`
	Type string `json:"ntfy.webhook.type" yaml:"ntfy.webhook.type"`
	// the key of text in request body or query string
	// default is text
	Key string `json:"ntfy.webhook.key" yaml:"ntfy.webhook.key"`
	// extra values
	Extra map[string]string `json:"ntfy.webhook.extra" yaml:"ntfy.webhook.extra"`
	// headers
	Header map[string]string `json:"ntfy.webhook.headers" yaml:"ntfy.webhook.headers"`
}

var _ Interface = (*WebhookConfig)(nil)

func (c WebhookConfig) Name() string {
	return "webhook"
}

func (c WebhookConfig) Notify(_ context.Context, text string) error {
	if c.URL == nil || *c.URL == "" {
		return ErrNoConfig
	}
	if c.Key == "" {
		c.Key = "text"
	}
	if c.Extra == nil {
		c.Extra = map[string]string{}
	}
	c.Extra[c.Key] = text
	_, err := resty.ParseResp[*webhookResp, *webhookResp](
		func() (*resty.Response, error) {
			req := R().SetHeaders(c.Header).SetError(&webhookResp{}).SetResult(&webhookResp{})
			if clean(c.Method) == "POST" {
				if clean(c.Type) == "JSON" {
					return req.SetBody(c.Extra).Post(*c.URL)
				}
				return req.SetFormData(c.Extra).Post(*c.URL)
			}
			return req.SetQueryParams(c.Extra).Get(*c.URL)
		}(),
	)
	if err != nil {
		return fmt.Errorf("[ntfy/%s] %w", c.Name(), err)
	}
	return nil
}

type webhookResp map[string]any

func (t *webhookResp) IsSuccess() bool {
	return true
}

func (t *webhookResp) String() string {
	return json.MustMarshalString(t)
}
