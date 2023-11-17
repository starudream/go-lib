package main

import (
	"context"
	"errors"
	"net"
	"net/http"
	"time"

	"github.com/starudream/go-lib/cobra/v2"
	"github.com/starudream/go-lib/core/v2/slog"
	"github.com/starudream/go-lib/core/v2/utils/osutil"
	"github.com/starudream/go-lib/core/v2/utils/signalutil"
	"github.com/starudream/go-lib/service/v2"
)

var (
	rootCmd = cobra.NewRootCommand(func(c *cobra.Command) {
		c.Use = "root"
		cobra.AddConfigFlag(c)
		service.AddCommand(c, svc())
	})

	serverCmd = cobra.NewCommand(func(c *cobra.Command) {
		c.Use = "server"
		c.Short = "Run server"
		c.Run = func(cmd *cobra.Command, args []string) {
			start(context.Background())
		}
	})
)

func init() {
	rootCmd.AddCommand(serverCmd)
}

func main() {
	osutil.ExitErr(rootCmd.Execute())
}

func start(context.Context) {
	server := &http.Server{}

	ln := osutil.Must1(net.Listen("tcp", ":8080"))
	slog.Info("listening on %s", ln.Addr())

	go func() {
		err := server.Serve(ln)
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			osutil.PanicErr(err)
		}
	}()

	slog.Info("server started")

	<-signalutil.Defer(stop(server)).Done()
}

func stop(server *http.Server) func() {
	return func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		_ = server.Shutdown(ctx)

		slog.Info("server stopped")
	}
}

func svc() service.Service {
	return service.New("example", start, service.WithArguments("server"))
}
