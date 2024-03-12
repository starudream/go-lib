package hggw

import (
	"context"
	"fmt"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"github.com/starudream/go-lib/core/v2/utils/optionutil"
	"github.com/starudream/go-lib/server/v2"
	"github.com/starudream/go-lib/server/v2/http"
)

type Server struct {
	*http.Server
	mux *runtime.ServeMux

	gwHandlers []gwHandler

	muxOpts  []runtime.ServeMuxOption
	dialOpts []grpc.DialOption
}

func NewServer(options ...Option) *Server {
	s := optionutil.Build(&Server{
		Server:  http.NewServer(),
		muxOpts: []runtime.ServeMuxOption{},
		dialOpts: []grpc.DialOption{
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(64 * 1024 * 1024)),
		},
	}, options...)
	s.mux = runtime.NewServeMux(s.muxOpts...)
	s.Mount("/", s.mux)
	return s
}

var _ server.Server = (*Server)(nil)

func (s *Server) Start(ln net.Listener) error {
	endpoint := fmt.Sprintf("%s:%d", localIP(), ln.Addr().(*net.TCPAddr).Port)
	conn, err := grpc.Dial(endpoint, s.dialOpts...)
	if err != nil {
		return err
	}
	for _, handler := range s.gwHandlers {
		err = handler(context.Background(), s.mux, conn)
		if err != nil {
			return err
		}
	}
	return s.Server.Start(ln)
}

func (s *Server) Stop(timeout time.Duration) {
	s.Server.Stop(timeout)
}

type gwHandler func(context.Context, *runtime.ServeMux, *grpc.ClientConn) error

func (s *Server) RegisterHandler(fn gwHandler) {
	s.gwHandlers = append(s.gwHandlers, fn)
}
