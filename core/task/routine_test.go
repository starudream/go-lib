package task

import (
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	Run(func() { panic("run") })
}

func TestGo(t *testing.T) {
	Go(func() { panic("go") })
	time.Sleep(time.Second)
}
