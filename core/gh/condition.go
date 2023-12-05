package gh

import (
	"io"

	"github.com/starudream/go-lib/core/v2/utils/reflectutil"
)

func Ternary[T any](t bool, ifVal T, elseVal T) T {
	if t {
		return ifVal
	}
	return elseVal
}

func TernaryF[T any](t bool, ifFn func() T, elseFn func() T) T {
	if t {
		return ifFn()
	}
	return elseFn()
}

func Close(a ...any) {
	for i := 0; i < len(a); i++ {
		if b, ok := a[i].(io.Closer); ok && !reflectutil.IsNil(a[i]) {
			Silently(b.Close())
		}
	}
}

func Silently(_ ...any) {}
