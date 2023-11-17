package fmtutil

import (
	"bufio"
	"fmt"
	"os"
)

func Scan(hint string) (s string) {
	in := bufio.NewScanner(os.Stdin)
	for fmt.Print(hint); in.Scan(); fmt.Print(hint) {
		if s = in.Text(); s != "" {
			break
		}
	}
	return
}
