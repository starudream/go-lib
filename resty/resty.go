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

var (
	_c     *Client
	_cOnce sync.Once
)

func C() *Client {
	_cOnce.Do(func() {
		_c = resty.New()
		_c.SetDisableWarn(true)
		_c.SetLogger(&logger{})
		_c.SetDebug(osutil.DOT())
		_c.SetDebugBodyLimit(1 << 16) // 65536
		_c.SetJSONMarshaler(json.Marshal)
		_c.SetJSONUnmarshaler(json.Unmarshal)
		_c.SetHeader(HeaderUserAgent, runtime.Version())
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
