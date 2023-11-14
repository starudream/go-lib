package resty

import (
	"errors"
	"fmt"

	"github.com/starudream/go-lib/core/v2/utils/reflectutil"
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
		return t, re
	}
	if resp.IsError() {
		err := resp.Error().(Err)
		if !reflectutil.IsNil(err) {
			re.esg = err.String()
		} else {
			re.err = fmt.Errorf("response status: %s", resp.Status())
		}
		return t, re
	}
	res := resp.Result().(Res)
	if !reflectutil.IsNil(res) && !res.IsSuccess() {
		re.esg = res.String()
		return t, re
	}
	return res, nil
}

type RespErr struct {
	*Response
	err error
	esg string
}

func (e *RespErr) String() string {
	if e.err != nil {
		e.esg = e.err.Error()
	}
	return fmt.Sprintf("response status: %s, error: %s", e.Response.Status(), e.esg)
}

func (e *RespErr) Error() string {
	return e.String()
}

func AsRespErr(err error) (*RespErr, bool) {
	re := &RespErr{}
	ok := errors.As(err, &re)
	return re, ok
}
