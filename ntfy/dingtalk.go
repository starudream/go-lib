package ntfy

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strconv"
	"time"

	"github.com/starudream/go-lib/resty/v2"
)

// https://open.dingtalk.com/document/robots/robot-overview

type DingtalkConfig struct {
	Token  *string `json:"ntfy.dingtalk.token"  yaml:"ntfy.dingtalk.token"`
	Secret string  `json:"ntfy.dingtalk.secret" yaml:"ntfy.dingtalk.secret"`
}

var _ Interface = (*DingtalkConfig)(nil)

func (DingtalkConfig) Name() string {
	return "dingtalk"
}

func (c DingtalkConfig) Notify(_ context.Context, text string) error {
	if c.Token == nil || *c.Token == "" {
		return ErrNoConfig
	}
	req := &dingtalkReq{MsgType: "text", Text: &dingtalkText{Content: text}}
	addr := "https://oapi.dingtalk.com/robot/send?access_token=" + *c.Token
	if c.Secret != "" {
		ts, sign := c.Sign()
		addr += "&timestamp=" + ts + "&sign=" + sign
	}
	_, err := resty.ParseResp[*dingtalkResp, *dingtalkResp](
		resty.R().SetBody(req).SetError(&dingtalkResp{}).SetResult(&dingtalkResp{}).Post(addr),
	)
	if err != nil {
		return fmt.Errorf("[ntfy/dingtalk] %w", err)
	}
	return nil
}

func (c DingtalkConfig) Sign() (string, string) {
	milli := strconv.FormatInt(time.Now().UnixMilli(), 10)
	h := hmac.New(sha256.New, []byte(c.Secret))
	h.Write([]byte(milli + "\n" + c.Secret))
	return milli, base64.StdEncoding.EncodeToString(h.Sum(nil))
}

type dingtalkReq struct {
	MsgType string        `json:"msgtype"`
	Text    *dingtalkText `json:"text,omitempty"`
}

type dingtalkText struct {
	Content string `json:"content"`
}

type dingtalkResp struct {
	ErrCode *int   `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func (t *dingtalkResp) IsSuccess() bool {
	return t != nil && t.ErrCode != nil && *t.ErrCode == 0
}

func (t *dingtalkResp) String() string {
	return fmt.Sprintf("code: %d, msg: %s", *t.ErrCode, t.ErrMsg)
}
