package main

import (
	"github.com/starudream/go-lib/cobra/v2"
	"github.com/starudream/go-lib/core/v2/slog"
)

var rootCmd = cobra.NewRootCommand(func(c *cobra.Command) {
	c.Use = "app"
	c.Run = func(cmd *cobra.Command, args []string) {
		slog.Info("root", slog.Any("args", args))
	}
	cobra.AddConfigFlag(c)
})

func init() {
	rootCmd.AddCommand(serveCmd)
}
