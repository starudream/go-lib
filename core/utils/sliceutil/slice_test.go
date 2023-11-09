package sliceutil

import (
	"testing"
)

func TestGetValue(t *testing.T) {
	t.Log(GetValue([]string{"a", "b"}, 1))
	t.Log(GetValue([]string{"a", "b"}, 2))
	t.Log(GetValue([]string{"a", "b"}, 3, "c"))
}
