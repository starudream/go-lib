package poolutil

import (
	"bytes"
	"sync"
)

var BytesBuffer1024 = NewBytesBuffer(1024)

type BytesBuffer struct {
	pool sync.Pool
}

func NewBytesBuffer(capacity int) *BytesBuffer {
	return &BytesBuffer{
		pool: sync.Pool{
			New: func() any {
				return bytes.NewBuffer(make([]byte, 0, capacity))
			},
		},
	}
}

var _ Pool[*bytes.Buffer] = (*BytesBuffer)(nil)

func (p *BytesBuffer) Get() *bytes.Buffer {
	buf := p.pool.Get().(*bytes.Buffer)
	buf.Reset()
	return buf
}

func (p *BytesBuffer) Put(b *bytes.Buffer) {
	p.pool.Put(b)
}
