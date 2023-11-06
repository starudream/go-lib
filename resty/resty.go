package resty

import (
	"runtime"
	"sync"

	"github.com/go-resty/resty/v2"

	"github.com/starudream/go-lib/core/v2/codec/json"
	"github.com/starudream/go-lib/core/v2/utils/osutil"
)

type (
	Client = resty.Client

	Request  = resty.Request
	Response = resty.Response
)

func New() *Client {
	c := resty.New()
	c.SetDisableWarn(true)
	c.SetLogger(&logger{})
	c.SetDebug(osutil.DOT())
	c.SetDebugBodyLimit(1 << 16) // 65536
	c.SetJSONMarshaler(json.Marshal)
	c.SetJSONUnmarshaler(json.Unmarshal)
	c.SetHeader(HeaderUserAgent, runtime.Version())
	return c
}

var (
	_c     *Client
	_cOnce sync.Once
)

func C() *Client {
	_cOnce.Do(func() {
		_c = New()
	})
	return _c
}

func R(ros ...rOptionI) *Request {
	opts := &rOptions{
		Headers: map[string]string{},
	}
	for i := 0; i < len(ros); i++ {
		ros[i].apply(opts)
	}
	return C().R().SetHeaders(opts.Headers)
}
