package sliceutil

import (
	"testing"
)

func TestSyncSlice(t *testing.T) {
	s := SyncSlice[string]{}
	s.Append("a", "b")
	t.Log(s.Data())
	t.Log(s.Index(1))
	// t.Log(s.Index(2))
	s.Delete(0)
	t.Log(s.Data())
}
