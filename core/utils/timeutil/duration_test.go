package timeutil

import (
	"testing"
	"time"
)

func TestJitterDuration(t *testing.T) {
	for i := 0; i < 20; i++ {
		t.Log(JitterDuration(time.Second, time.Minute, i))
	}
}
