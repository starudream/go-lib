package resty

import (
	"errors"
	"fmt"
)

type iRespErr interface {
	String() string
}

type iRespRes interface {
	IsSuccess() bool
	String() string
}

func ParseResp[Err iRespErr, Res iRespRes](resp *Response, ee error) (t Res, _ error) {
	if ee != nil {
		return t, &RespErr{Response: resp, err: fmt.Errorf("execute error: %w", ee)}
	}
	if resp.IsError() {
		err := resp.Error().(Err)
		return t, &RespErr{Response: resp, msg: err.String()}
	}
	res := resp.Result().(Res)
	if !res.IsSuccess() {
		return t, &RespErr{Response: resp, msg: res.String()}
	}
	return res, nil
}

type RespErr struct {
	*Response
	err error
	msg string
}

func (e *RespErr) Error() string {
	if e.err != nil {
		e.msg = e.err.Error()
	}
	return fmt.Sprintf("response status: %s, error: %s", e.Response.Status(), e.msg)
}

func AsRespErr(err error) (*RespErr, bool) {
	re := &RespErr{}
	ok := errors.As(err, &re)
	return re, ok
}
