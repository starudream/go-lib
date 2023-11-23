package main

import (
	"github.com/starudream/go-lib/cobra/v2"
	"github.com/starudream/go-lib/core/v2/utils/fmtutil"
	"github.com/starudream/go-lib/core/v2/utils/osutil"
	"github.com/starudream/go-lib/selfupdate/v2"
)

var rootCmd = cobra.NewRootCommand(func(c *cobra.Command) {
	c.Use = "selfupdate"
	c.Run = run
})

func main() {
	// run(nil, nil)
	osutil.ExitErr(rootCmd.Execute())
}

func run(*cobra.Command, []string) {
	// source := &selfupdate.GoReleaser{
	// 	Owner: "starudream",
	// 	Repo:  "secret-tunnel",
	// 	Name:  "secret-tunnel-client",
	// }

	source := &selfupdate.GoReleaser{
		Owner: "starudream",
		Repo:  "douyu-task",
	}

	osutil.ExitErr(selfupdate.Update(source, func() bool { return fmtutil.Scan("update? (y/n): ") == "y" }))
}
