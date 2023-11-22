package ping

import (
	"github.com/go-ping/ping"

	"github.com/starudream/go-lib/core/v2/slog"
	"github.com/starudream/go-lib/core/v2/utils/signalutil"
)

type (
	Packet     = ping.Packet
	Statistics = ping.Statistics
)

func Ping(options ...Option) (*Statistics, error) {
	opts := newOptions(options...)

	pinger, err := ping.NewPinger(opts.addr)
	if err != nil {
		return nil, err
	}

	pinger.Interval = opts.interval
	pinger.Timeout = opts.timeout
	pinger.Count = opts.count

	pinger.SetPrivileged(opts.privileged)

	pinger.OnRecv = func(pkt *Packet) {
		slog.Debug("%d bytes from %s: icmp_seq=%d time=%v ttl=%v",
			pkt.Nbytes, pkt.IPAddr, pkt.Seq, pkt.Rtt, pkt.Ttl, slog.String("addr", pkt.Addr))
	}
	pinger.OnDuplicateRecv = func(pkt *Packet) {
		slog.Debug("%d bytes from %s: icmp_seq=%d time=%v ttl=%v (DUP!)",
			pkt.Nbytes, pkt.IPAddr, pkt.Seq, pkt.Rtt, pkt.Ttl, slog.String("addr", pkt.Addr))
	}
	pinger.OnFinish = func(stats *Statistics) {
		slog.Debug("%d packets transmitted, %d packets received, %d duplicates, %v%% packet loss, min/avg/max/stddev = %v/%v/%v/%v",
			stats.PacketsSent, stats.PacketsRecv, stats.PacketsRecvDuplicates, stats.PacketLoss,
			stats.MinRtt, stats.AvgRtt, stats.MaxRtt, stats.StdDevRtt, slog.String("addr", stats.Addr))
	}

	slog.Debug("ping start", slog.String("addr", pinger.Addr()), slog.String("ip", pinger.IPAddr().String()))

	go func() {
		<-signalutil.Defer(pinger.Stop).Done()
	}()

	err = pinger.Run()
	if err != nil {
		return nil, err
	}

	return pinger.Statistics(), nil
}
