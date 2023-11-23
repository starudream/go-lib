package archive

import (
	"archive/tar"
	"bytes"
	"errors"
	"io"
)

type Tar []byte

var _ Interface = (*Tar)(nil)

func (t Tar) Get(name string) ([]byte, error) {
	return TarReader{bytes.NewReader(t)}.Get(name)
}

type TarReader struct {
	io.Reader
}

var _ Interface = (*TarReader)(nil)

func (t TarReader) Get(name string) ([]byte, error) {
	tr := tar.NewReader(t)
	for {
		th, err := tr.Next()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return nil, err
		}
		if th.Name == name {
			return io.ReadAll(tr)
		}
	}
	return nil, ErrFileNotFound
}
