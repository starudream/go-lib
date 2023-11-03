package cobra

import (
	"bytes"
	"strings"
	"testing"

	"github.com/starudream/go-lib/core/v2/slog"
	"github.com/starudream/go-lib/core/v2/utils/testutil"
)

var (
	rootCmd = NewRootCommand(func(c *Command) {
		c.Use = "test"
		c.Run = func(cmd *Command, args []string) {
			slog.Info("root", slog.Any("args", args))
		}
		AddConfigFlag(c, "", "")
	})

	echoCmd = NewCommand(func(c *Command) {
		c.Use = "echo"
		c.Aliases = []string{"e"}
		c.Short = "Echo command"
		c.Run = func(cmd *Command, args []string) {
			slog.Info("echo", slog.Any("args", args))
		}
	})
)

func init() {
	rootCmd.AddCommand(echoCmd)
}

func exec(args ...string) (string, error) {
	buf := &bytes.Buffer{}
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	rootCmd.SetArgs(args)
	err := rootCmd.Execute()
	return strings.TrimSuffix(buf.String(), "\n"), err
}

func TestHelp(t *testing.T) {
	out, err := exec("-h")
	testutil.LogNoErr(t, err, "\n"+out)
}

func TestVersion(t *testing.T) {
	out, err := exec("-v")
	testutil.LogNoErr(t, err, "\n"+out)
}
