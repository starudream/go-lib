package main

import (
	"github.com/starudream/go-lib/cobra/v2"
	"github.com/starudream/go-lib/core/v2/slog"
)

var serveCmd = cobra.NewCommand(func(c *cobra.Command) {
	c.Use = "serve"
	c.Run = func(cmd *cobra.Command, args []string) {
		slog.Info("serve", slog.Any("args", args))
		select {}
	}
})
