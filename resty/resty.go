package resty

import (
	"os"
	"runtime"
	"strconv"
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
	c.SetLogger(_logger)
	c.SetDebug(debug())
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
	_cOnce.Do(func() { _c = New() })
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

func debug() bool {
	v := osutil.DOT()
	if s := os.Getenv("RESTY_DEBUG"); s != "" {
		t, e := strconv.ParseBool(s)
		if e == nil {
			v = t
		}
	}
	return v
}
