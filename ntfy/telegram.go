package ntfy

import (
	"context"
	"fmt"

	"github.com/starudream/go-lib/resty/v2"
)

// https://core.telegram.org/bots/api#sendmessage

type TelegramConfig struct {
	Token  *string `json:"ntfy.telegram.token"   yaml:"ntfy.telegram.token"`
	ChatId string  `json:"ntfy.telegram.chat_id" yaml:"ntfy.telegram.chat_id"`
}

var _ Interface = (*TelegramConfig)(nil)

func (TelegramConfig) Name() string {
	return "telegram"
}

func (c TelegramConfig) Notify(_ context.Context, text string) error {
	if c.Token == nil || *c.Token == "" {
		return ErrNoConfig
	}
	req := &telegramReq{ChatId: c.ChatId, Text: text}
	addr := "https://api.telegram.org/bot" + *c.Token + "/sendMessage"
	_, err := resty.ParseResp[*telegramResp, *telegramResp](
		R().SetBody(req).SetError(&telegramResp{}).SetResult(&telegramResp{}).AddRetryCondition(c.retry).Post(addr),
	)
	if err != nil {
		return fmt.Errorf("[ntfy/telegram] %w", err)
	}
	return nil
}

func (c TelegramConfig) retry(resp *resty.Response, err error) bool {
	return err != nil || resp.IsError()
}

type telegramReq struct {
	ChatId string `json:"chat_id"`
	Text   string `json:"text"`
}

type telegramResp struct {
	OK          bool   `json:"ok"`
	ErrorCode   int    `json:"error_code,omitempty"`
	Description string `json:"description,omitempty"`
}

func (t *telegramResp) IsSuccess() bool {
	return t.OK
}

func (t *telegramResp) String() string {
	return fmt.Sprintf("ok: %v, code: %d, desc: %s", t.OK, t.ErrorCode, t.Description)
}
