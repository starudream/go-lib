package tablew

import (
	"testing"
)

func TestStructs(t *testing.T) {
	type T struct {
		A string
		B int
	}
	s := Structs([]T{
		{"a", 1},
		{"b", 2},
	})
	t.Log("\n" + s)
}
