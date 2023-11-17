package osutil

import (
	"fmt"
	"os"
	"testing"
)

var (
	err = fmt.Errorf("test error")
	ner = error(nil)
)

func TestExitErr(t *testing.T) {
	Exit = func(t int, s string) { fmt.Printf("exit: %d, %s\n", t, s) }

	t.Run("no-error", func(t *testing.T) { ExitErr(ner) })
	t.Run("no-error-code", func(t *testing.T) { ExitErr(ner, 0) })
	t.Run("no-error-msg", func(t *testing.T) { ExitErr(ner, "foo %s", "bar") })

	t.Run("error", func(t *testing.T) { ExitErr(err) })
	t.Run("error-code", func(t *testing.T) { ExitErr(err, 2) })
	t.Run("error-msg", func(t *testing.T) { ExitErr(err, "helle %s, %v", "world", map[string]any{"foo": "bar"}) })
}

func TestPanicErr(t *testing.T) {
	Panic = func(t int, s string) { fmt.Printf("panic: %d, %s\n", t, s) }

	t.Run("no-error", func(t *testing.T) { PanicErr(ner) })

	t.Run("error", func(t *testing.T) { PanicErr(err) })
	t.Run("error-msg", func(t *testing.T) { PanicErr(err, "helle %s, %v", "world", map[string]any{"foo": "bar"}) })

	t.Run("must1", func(t *testing.T) { Must1(os.Open("foo")) })
}

func TestRunErr(t *testing.T) {
	fn := func(t int, s string) { fmt.Printf("run: %d, %s\n", t, s) }

	t.Run("no-error", func(t *testing.T) { RunErr(fn, ner) })

	t.Run("error", func(t *testing.T) { RunErr(fn, err) })
}
