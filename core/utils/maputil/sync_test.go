package maputil

import (
	"testing"
)

func TestSyncMap(t *testing.T) {
	m := SyncMap[string, int]{}
	t.Log(m.Load("a"))
	t.Log(m.LoadOrStore("a", 1))
	m.Store("a", 2)
	t.Log(m.CompareAndSwap("a", 3, 1))
	t.Log(m.LoadAndDelete("a"))
	t.Log(m.LoadAndDelete("a"))
	t.Log(m.CompareAndDelete("a", 1))
}
