package gh

import (
	"testing"
)

type Closer struct{}

func (c Closer) Close() error { return nil }

func TestClose(t *testing.T) {
	c1 := Closer{}
	Close(c1)

	var c2 *Closer
	Close(c2)
}
