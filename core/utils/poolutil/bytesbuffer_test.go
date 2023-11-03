package poolutil

import (
	"strconv"
	"testing"
)

func TestBytesBuffer(t *testing.T) {
	p := NewBytesBuffer(128)

	buf1 := p.Get()
	t.Logf("%p", buf1)
	for i := 0; i < 10; i++ {
		buf1.WriteString(strconv.Itoa(i))
	}
	t.Log(buf1.String())
	p.Put(buf1)

	buf2 := p.Get()
	t.Logf("%p", buf2)   // same address as buf1
	t.Log(buf2.String()) // empty string because of Reset

	buf3 := p.Get()
	t.Logf("%p", buf3) // new address different from buf1 and buf2
	p.Put(buf3)

	p.Put(buf2)

	buf4 := p.Get()
	t.Logf("%p", buf4)
	p.Put(buf4)
}
