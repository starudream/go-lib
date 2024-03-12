package slog_test

import (
	"fmt"
	"testing"

	"github.com/starudream/go-lib/core/v2/slog"
)

func init() {
	slog.OnFatal = func() { fmt.Println("on fatal") }
}

func Test(t *testing.T) {
	slog.Debug("hello %s", "world")
	slog.Info("info", slog.String("foo", "bar"))
	slog.Warn("warn")
	slog.Error("error")
	slog.Fatal("fatal")
}

// func TestDailyRotate(t *testing.T) {
// 	for i := 0; i < 3; i++ {
// 		slog.Info("hello %s", "world")
// 		time.Sleep(time.Second)
// 	}
// }
