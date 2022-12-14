package httpx

import (
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"

	"github.com/starudream/go-lib/codec/json"
	"github.com/starudream/go-lib/codec/xml"
	"github.com/starudream/go-lib/config"
	"github.com/starudream/go-lib/log"

	"github.com/starudream/go-lib/internal/httpproxy"
	"github.com/starudream/go-lib/internal/ilog"
)

type (
	Request  = resty.Request
	Response = resty.Response
)

var _c *resty.Client

func init() {
	_c = resty.New()
	_c.JSONMarshal = json.Marshal
	_c.JSONUnmarshal = json.Unmarshal
	_c.XMLMarshal = xml.Marshal
	_c.XMLUnmarshal = xml.Unmarshal

	_c.SetTimeout(5 * time.Minute)
	_c.SetLogger(&logger{Logger: log.With().Str("span", "http").Logger()})
	_c.SetDisableWarn(true)
	_c.SetDebug(config.GetBool("debug"))

	pc := httpproxy.FromEnvironment()
	if pc != nil && (pc.HTTPProxy != "" || pc.HTTPSProxy != "") {
		ilog.X.Debug().Msgf("proxy: %s", json.MustMarshalString(pc))
	}
}

func SetTimeout(timeout time.Duration) {
	_c.SetTimeout(timeout)
}

var hdrUserAgentKey = http.CanonicalHeaderKey("User-Agent")

func SetUserAgent(ua string) {
	_c.SetHeader(hdrUserAgentKey, ua)
}

func R() *resty.Request {
	return _c.R()
}
