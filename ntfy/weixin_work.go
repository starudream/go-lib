package ntfy

import (
	"context"
	"fmt"

	"github.com/starudream/go-lib/resty/v2"
)

// https://developer.work.weixin.qq.com/document/path/91770

type WeixinWorkConfig struct {
	Key *string `json:"ntfy.weixin_work.key" yaml:"ntfy.weixin_work.key"`
}

var _ Interface = (*WeixinWorkConfig)(nil)

func (WeixinWorkConfig) Name() string {
	return "weixin_work"
}

func (c WeixinWorkConfig) Notify(_ context.Context, text string) error {
	if c.Key == nil || *c.Key == "" {
		return ErrNoConfig
	}
	req := &weixinWorkReq{MsgType: "text", Text: &weixinWorkText{Content: text}}
	addr := "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=" + *c.Key
	_, err := resty.ParseResp[*weixinWorkResp, *weixinWorkResp](
		R().SetBody(req).SetError(&weixinWorkResp{}).SetResult(&weixinWorkResp{}).AddRetryCondition(c.retry).Post(addr),
	)
	if err != nil {
		return fmt.Errorf("[ntfy/%s] %w", c.Name(), err)
	}
	return nil
}

func (c WeixinWorkConfig) retry(resp *resty.Response, err error) bool {
	return err != nil || resp.IsError()
}

type weixinWorkReq struct {
	MsgType string          `json:"msgtype"`
	Text    *weixinWorkText `json:"text"`
}

type weixinWorkText struct {
	Content string `json:"content"`
}

type weixinWorkResp struct {
	ErrCode *int   `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func (t *weixinWorkResp) IsSuccess() bool {
	return t != nil && t.ErrCode != nil && *t.ErrCode == 0
}

func (t *weixinWorkResp) String() string {
	return fmt.Sprintf("code: %d, msg: %s", *t.ErrCode, t.ErrMsg)
}
