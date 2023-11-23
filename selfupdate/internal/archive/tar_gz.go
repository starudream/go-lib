package archive

import (
	"bytes"
	"compress/gzip"
)

type TarGz []byte

var _ Interface = (*TarGz)(nil)

func (t TarGz) Get(name string) ([]byte, error) {
	gr, err := gzip.NewReader(bytes.NewReader(t))
	if err != nil {
		return nil, err
	}
	defer func() { _ = gr.Close() }()
	return TarReader{gr}.Get(name)
}
