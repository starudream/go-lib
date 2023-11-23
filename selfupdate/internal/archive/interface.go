package archive

import (
	"errors"
)

type Interface interface {
	Get(name string) ([]byte, error)
}

var ErrFileNotFound = errors.New("file not found")
