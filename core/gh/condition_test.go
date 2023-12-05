package gh

import (
	"testing"
)

type Closer struct{}

func (c Closer) Close() error { return nil }

func TestClose(t *testing.T) {
	c1 := Closer{}
	Close(c1)

	c2 := &Closer{}
	c2 = nil
	Close(c2)
}
