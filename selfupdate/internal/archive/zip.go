package archive

import (
	"archive/zip"
	"bytes"
	"io"
)

type Zip []byte

var _ Interface = (*Zip)(nil)

func (t Zip) Get(name string) ([]byte, error) {
	zr, err := zip.NewReader(bytes.NewReader(t), int64(len(t)))
	if err != nil {
		return nil, err
	}
	for _, file := range zr.File {
		if file.Name == name {
			return t.read(file)
		}
	}
	return nil, ErrFileNotFound
}

func (t Zip) read(file *zip.File) ([]byte, error) {
	r, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer r.Close()
	return io.ReadAll(r)
}
