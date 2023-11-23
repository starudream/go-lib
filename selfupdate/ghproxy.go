package selfupdate

import (
	"errors"
	"fmt"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/starudream/go-lib/core/v2/gh"
	"github.com/starudream/go-lib/core/v2/slog"
	"github.com/starudream/go-lib/core/v2/utils/maputil"
	"github.com/starudream/go-lib/core/v2/utils/osutil"
	"github.com/starudream/go-lib/resty/v2"
	"github.com/starudream/go-lib/resty/v2/ping"
)

const (
	githubDefault     = "https://github.com/"
	githubReleaseFile = "starudream/go-lib/releases/download/v2.0.0-rc.1/ping"
)

var githubProxies = []string{
	githubDefault,
	"https://ghproxy.net/https://github.com/",
	"https://gh-proxy.com/https://github.com/",
	"https://mirror.ghproxy.com/https://github.com/",
	"https://ghproxy.starudream.cn/https://github.com/",
}

type GithubProxy struct {
	proxies []string
	timeout time.Duration
	client  *resty.Client

	stats maputil.SyncMap[string, *ping.Statistics]
	costs maputil.SyncMap[string, time.Duration]
}

func (t *GithubProxy) Add(proxies ...string) *GithubProxy {
	for i := 0; i < len(proxies); i++ {
		proxy := proxies[i]
		u, err := url.Parse(proxy)
		if err != nil {
			slog.Warn("invalid github proxy: %v", err, slog.String("proxy", proxy))
		}
		u.ForceQuery = false
		u.RawQuery = ""
		u.Fragment = ""
		u.RawFragment = ""
		t.proxies = append(t.proxies, u.String())
	}
	return t
}

func (t *GithubProxy) SetTimeout(timeout time.Duration) *GithubProxy {
	t.timeout = timeout
	return t
}

func (t *GithubProxy) Check() *GithubProxy {
	if len(t.proxies) == 0 {
		t.proxies = make([]string, len(githubProxies))
		copy(t.proxies, githubProxies)
	}
	if t.timeout <= 0 {
		t.timeout = 3 * time.Second
	}
	t.client = resty.New().SetTimeout(t.timeout)
	slog.Debug("github proxy check, maybe take a few minutes", slog.Int("total", len(t.proxies)))
	wg := sync.WaitGroup{}
	wg.Add(2 * len(t.proxies))
	for _, proxy := range t.proxies {
		go func(proxy string) {
			defer wg.Done()
			t.stats.Store(proxy, t.ping(proxy))
		}(proxy)
		go func(proxy string) {
			defer wg.Done()
			t.costs.Store(proxy, t.down(proxy))
		}(proxy)
	}
	wg.Wait()
	return t
}

func (t *GithubProxy) ping(proxy string) *ping.Statistics {
	stat, err := ping.Ping(ping.WithAddr(osutil.Must1(url.Parse(proxy)).Host), ping.WithCount(1), ping.WithTimeout(3*time.Second))
	if err != nil {
		slog.Warn("github proxy ping error: %v", err, slog.String("proxy", proxy))
	}
	return stat
}

func (t *GithubProxy) down(proxy string) time.Duration {
	resp, err := t.client.R().SetDoNotParseResponse(true).Get(proxy + githubReleaseFile)
	defer gh.Close(resp.RawBody())
	if err != nil || resp.StatusCode() != 200 {
		if err == nil {
			err = errors.New(resp.Status())
		}
		slog.Warn("github proxy download error: %v", err, slog.String("proxy", proxy))
		return -1
	}
	return resp.Time()
}

func (t *GithubProxy) Items() []string {
	t.sort()
	items := make([]string, len(t.proxies))
	for i, proxy := range t.proxies {
		stat, _ := t.stats.Load(proxy)
		cost, _ := t.costs.Load(proxy)
		if proxy != githubDefault {
			proxy = strings.TrimSuffix(proxy, githubDefault)
		}
		if stat == nil || cost < 0 {
			items[i] = fmt.Sprintf("%-40s ***", proxy)
		} else {
			items[i] = fmt.Sprintf("%-40s %s/%s/%s", proxy, strconv.Itoa(int(stat.PacketLoss))+"%", stat.AvgRtt.Truncate(time.Microsecond), cost.Truncate(time.Millisecond))
		}
	}
	return items
}

func (t *GithubProxy) Fast() string {
	t.sort()
	if len(t.proxies) == 0 {
		return githubDefault
	}
	proxy := t.proxies[0]
	stat, _ := t.stats.Load(proxy)
	if stat == nil || stat.PacketLoss == 100 {
		return githubDefault
	}
	return proxy
}

func (t *GithubProxy) sort() {
	if len(t.proxies) == 0 {
		return
	}
	// cost > loss > rtt
	sort.Slice(t.proxies, func(i, j int) bool {
		pi, pj := t.proxies[i], t.proxies[j]

		ci, _ := t.costs.Load(pi)
		cj, _ := t.costs.Load(pj)
		if ci < 0 && cj < 0 {
		} else if ci < 0 {
			return false
		} else if cj < 0 {
			return true
		}

		si, _ := t.stats.Load(pi)
		sj, _ := t.stats.Load(pj)
		if si == nil && sj == nil {
			return pi < pj
		} else if si == nil {
			return false
		} else if sj == nil {
			return true
		}

		if si.PacketLoss == sj.PacketLoss {
			return si.AvgRtt < sj.AvgRtt
		}
		return si.PacketLoss < sj.PacketLoss
	})
}
