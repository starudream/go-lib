//go:build !linux

package filewriter

import (
	"os"
)

func chown(_ string, _ os.FileInfo) error {
	return nil
}
