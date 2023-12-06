package cobra

import (
	"os"

	"github.com/spf13/pflag"
)

type (
	Flag    = pflag.Flag
	FlagSet = pflag.FlagSet
)

func FlagArgs(fs *FlagSet, skipNames ...string) (args []string) {
	fs.ParseErrorsWhitelist.UnknownFlags = true

	visit := func(flag *Flag, value string) error {
		for _, n := range skipNames {
			if flag.Name == n {
				return nil
			}
		}
		args = append(args, "--"+flag.Name, value)
		return nil
	}

	_ = fs.ParseAll(os.Args[1:], visit)

	return args
}
