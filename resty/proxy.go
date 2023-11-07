package resty

import (
	"golang.org/x/net/http/httpproxy"

	"github.com/starudream/go-lib/core/v2/slog"
)

var ProxyFromEnv = httpproxy.FromEnvironment

func init() {
	proxy := ProxyFromEnv()
	slog.Debug("proxy config from env, http: %s, https: %s, no: %s", proxy.HTTPProxy, proxy.HTTPSProxy, proxy.NoProxy)
}
