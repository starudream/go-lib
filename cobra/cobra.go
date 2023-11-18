package cobra

import (
	"io"
	"os"

	"github.com/spf13/cobra"

	"github.com/starudream/go-lib/core/v2/config/version"
)

type Command = cobra.Command

func NewCommand(f ...func(c *Command)) *Command {
	cmd := &Command{}
	if len(f) > 0 && f[0] != nil {
		f[0](cmd)
	}
	return cmd
}

func NewRootCommand(f ...func(c *Command)) *Command {
	return NewCommand(func(c *Command) {
		// hidden help command
		c.SetHelpCommand(&Command{Hidden: true})
		// hidden completion command
		c.CompletionOptions.HiddenDefaultCmd = true
		// set version template
		c.SetVersionTemplate("{{ print .Version }}")
		c.Version = version.GetVersionInfo().String()
		// output
		c.SetOut(os.Stdout)
		c.SetErr(io.Discard)

		if len(f) > 0 && f[0] != nil {
			f[0](c)
		}
	})
}

func AddConfigFlag(c *Command, usage ...string) {
	if len(usage) == 0 || usage[0] == "" {
		usage = []string{"path to config file"}
	}
	c.PersistentFlags().StringP("config", "c", "", usage[0])
	_ = c.MarkPersistentFlagFilename("config")
}
