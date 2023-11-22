package resty

import (
	"golang.org/x/net/http/httpproxy"

	"github.com/starudream/go-lib/core/v2/slog"
)

var ProxyFromEnv = httpproxy.FromEnvironment

func init() {
	proxy := ProxyFromEnv()
	if proxy.HTTPProxy == "" && proxy.HTTPSProxy == "" {
		return
	}
	slog.Debug("proxy env loaded", slog.String("http", proxy.HTTPProxy), slog.String("https", proxy.HTTPSProxy), slog.String("no", proxy.NoProxy))
}

func HasProxyEnv() bool {
	proxy := ProxyFromEnv()
	return proxy.HTTPProxy != "" || proxy.HTTPSProxy != ""
}
