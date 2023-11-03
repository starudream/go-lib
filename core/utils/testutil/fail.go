package testutil

import (
	"strings"

	"github.com/starudream/go-lib/core/v2/utils/osutil"
)

func Fail(t T, failure string, msgAndArgs ...any) bool {
	t.Helper()
	fail(t, failure, msgAndArgs...)
	return false
}

func FailNow(t T, failure string, msgAndArgs ...any) {
	t.Helper()
	fail(t, failure, msgAndArgs...)
	t.FailNow()
}

func fail(t T, failure string, msgAndArgs ...any) {
	t.Helper()

	rs := make([][2]string, 0)

	if nt, ok := t.(interface{ Name() string }); ok {
		rs = append(rs, [2]string{"NAME", nt.Name()})
	}

	rs = append(rs, [2]string{"FAILURE", failure})

	if str := msgAndArgsString(msgAndArgs...); str != "" {
		rs = append(rs, [2]string{"MESSAGE", str})
	}

	skip := 3
	if tt, ok := t.(_skipT); ok {
		skip += tt.skip
	}
	rs = append(rs, [2]string{"STACK", osutil.Stack(skip)})

	t.Error("\n" + render(strings.Repeat(" ", 4), "", " : ", rs))
}
