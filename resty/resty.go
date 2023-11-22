package resty

import (
	"net"
	"net/http"
	"net/http/cookiejar"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"

	"github.com/go-resty/resty/v2"
	"golang.org/x/net/publicsuffix"

	"github.com/starudream/go-lib/core/v2/codec/json"
	"github.com/starudream/go-lib/core/v2/utils/osutil"
)

type (
	Client = resty.Client

	Request  = resty.Request
	Response = resty.Response
)

func New() *Client {
	c := newClient()
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

func R(rOptions ...ROption) *Request {
	opts := newROptions(rOptions...)
	return C().R().SetHeaders(opts.Headers)
}

func newClient() *Client {
	cookieJar, _ := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	return resty.NewWithClient(&http.Client{Transport: createTransport(), Jar: cookieJar})
}

func createTransport() *http.Transport {
	return &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		MaxIdleConnsPerHost:   runtime.GOMAXPROCS(0) + 1,
	}
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
