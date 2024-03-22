package main

import (
	"context"

	"github.com/starudream/go-lib/cobra/v2"
	"github.com/starudream/go-lib/core/v2/utils/osutil"
	"github.com/starudream/go-lib/server/v2"
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
			run(context.Background())
		}
	})
)

func init() {
	rootCmd.AddCommand(serverCmd)
}

func main() {
	osutil.ExitErr(rootCmd.Execute())
}

func run(context.Context) {
	hs, gs := NewHTTPServer(), NewGRPCServer()
	osutil.PanicErr(server.Run(":8080", server.WithHTTP(hs), server.WithGRPC(gs)))
}

func svc() service.Service {
	return service.New("example", run, service.WithArguments("server"))
}
