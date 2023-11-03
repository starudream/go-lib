package poolutil

import (
	"testing"
)

func TestBytes(t *testing.T) {
	p := NewBytes(1024, 10)

	b1 := p.Get()
	t.Logf("%p %d %d", b1, len(b1), cap(b1))
	p.Put(b1)

	b2 := p.Get()
	t.Logf("%p %d %d", b2, len(b2), cap(b2))

	b3 := p.Get()
	t.Logf("%p %d %d", b3, len(b3), cap(b3))
	p.Put(b3)

	p.Put(b2)

	b4 := p.Get()
	t.Logf("%p %d %d", b4, len(b4), cap(b4))
	p.Put(b4)
}
