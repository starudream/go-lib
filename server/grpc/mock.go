//go:build !release

package grpc

import (
	"context"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"

	"github.com/starudream/go-lib/core/v2/utils/osutil"
)

func NewMock(s *Server) (*ClientConn, func(m interface{ Run() int })) {
	ln := bufconn.Listen(256 * 1024)

	go func() {
		if err := s.Start(ln); err != nil {
			osutil.PanicErr(err)
		}
	}()

	dialer := func(context.Context, string) (net.Conn, error) { return ln.Dial() }
	conn, err := Dial("", grpc.WithContextDialer(dialer))
	if err != nil {
		osutil.PanicErr(err)
	}

	wait := func(m interface{ Run() int }) {
		code := m.Run()
		s.Stop(time.Second)
		osutil.ExitErr(ln.Close(), code)
	}

	return conn, wait
}
