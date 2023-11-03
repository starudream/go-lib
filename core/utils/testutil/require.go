package testutil

func MustNil(t T, value any, msgAndArgs ...any) {
	t.Helper()
	if Nil(skipT(t), value, msgAndArgs...) {
		return
	}
	t.FailNow()
}

func MustNotNil(t T, value any, msgAndArgs ...any) {
	t.Helper()
	if NotNil(skipT(t), value, msgAndArgs...) {
		return
	}
	t.FailNow()
}

func MustEqual(t T, expected, actual any, msgAndArgs ...any) {
	t.Helper()
	if Equal(skipT(t), expected, actual, msgAndArgs...) {
		return
	}
	t.FailNow()
}

func MustNotEqual(t T, expected, actual any, msgAndArgs ...any) {
	t.Helper()
	if NotEqual(skipT(t), expected, actual, msgAndArgs...) {
		return
	}
	t.FailNow()
}
