package osutil

import (
	"fmt"
	"os"
	"strings"
)

var Exit = func(t int, s string) {
	if s != "" {
		wPrintln(t, s)
	}
	if t >= 0 && !ArgTest() {
		os.Exit(t)
	}
}

func ExitErr(err error, a ...any) {
	RunErr(Exit, err, a...)
}

var Panic = func(t int, s string) {
	panic(s)
}

func PanicErr(err error, a ...any) {
	if err == nil {
		return
	}
	RunErr(Panic, err, a...)
}

func RunErr(fn func(t int, s string), err error, a ...any) {
	t := -1
	if len(a) > 0 {
		if i, ok := a[0].(int); ok {
			t = i
			a = a[1:]
		}
	}

	f := ""
	if len(a) > 0 {
		if s, ok := a[0].(string); ok && strings.Contains(s, "%") {
			f = s
			a = a[1:]
		}
	}
	if f != "" || len(a) > 0 {
		if f == "" {
			f = "%v"
		} else {
			f = strings.TrimSuffix(f, "\n")
		}
	}

	if err == nil {
		if f != "" {
			if DOT() {
				f = fmt.Sprintf("[%s] ", CallerString(2)) + f
			}
			f = fmt.Sprintf(f, a...)
		}
		fn(t, f)
		return
	}

	if t < 0 {
		t = 1
	}

	if f != "" {
		f = ", " + f
	}

	fn(t, fmt.Sprintf(fmt.Sprintf("[%s] exit: %d, error: %v", CallerString(2), t, err)+f, a...))
}

func wPrintln(t int, s string) {
	w := os.Stderr
	if t == 0 {
		w = os.Stdout
	}
	_, _ = fmt.Fprintln(w, strings.TrimSuffix(s, "\n"))
}
