package testutil

import (
	"fmt"
)

func Log(t T, msgAndArgs ...any) {
	t.Helper()
	if str := msgAndArgsString(msgAndArgs...); str != "" {
		t.Log(str)
	}
}

func LogNoErr(t T, err error, msgAndArgs ...any) {
	t.Helper()
	if err != nil {
		FailNow(t, fmt.Sprintf("unexpected error: %v", err))
	} else {
		Log(t, msgAndArgs...)
	}
}
