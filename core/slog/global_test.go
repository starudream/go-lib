package slog

import (
	"fmt"
	"testing"
)

func init() {
	OnFatal = func() { fmt.Println("on fatal") }
}

func Test(t *testing.T) {
	Debug("hello %s", "world")
	Info("info", String("foo", "bar"))
	Warn("warn")
	Error("error")
	Fatal("fatal")
}

// func TestDailyRotate(t *testing.T) {
// 	for i := 0; i < 3; i++ {
// 		Info("hello %s", "world")
// 		time.Sleep(time.Second)
// 	}
// }
