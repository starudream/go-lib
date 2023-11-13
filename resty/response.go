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
	re := &RespErr{Response: resp}
	if ee != nil {
		re.err = fmt.Errorf("execute error: %w", ee)
		return
	}
	if resp.IsError() {
		err := resp.Error().(Err)
		if err != nil {
			re.msg = err.String()
		} else {
			re.err = fmt.Errorf("response status: %s", resp.Status())
		}
		return
	}
	res := resp.Result().(Res)
	if res != nil && !res.IsSuccess() {
		re.msg = res.String()
		return
	}
	return res, nil
}

type RespErr struct {
	*Response
	err error
	msg string
}

func (e *RespErr) String() string {
	if e.err != nil {
		e.msg = e.err.Error()
	}
	return fmt.Sprintf("response status: %s, error: %s", e.Response.Status(), e.msg)
}

func (e *RespErr) Error() string {
	return e.String()
}

func AsRespErr(err error) (*RespErr, bool) {
	re := &RespErr{}
	ok := errors.As(err, &re)
	return re, ok
}
