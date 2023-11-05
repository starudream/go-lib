package resty

import (
	"fmt"
)

type respErr interface {
	String() string
}

type respRes interface {
	IsSuccess() bool
	String() string
}

func ParseResp[Err respErr, Res respRes](resp *Response, ee error) (t Res, _ error) {
	if ee != nil {
		return t, fmt.Errorf("execute error: %w", ee)
	}
	if resp.IsError() {
		err := resp.Error().(Err)
		return t, fmt.Errorf("response status: %s error: %s", resp.Status(), err.String())
	}
	res := resp.Result().(Res)
	if !res.IsSuccess() {
		return t, fmt.Errorf("response status: %s error: %s", resp.Status(), res.String())
	}
	return res, nil
}
