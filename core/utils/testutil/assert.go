package testutil

import (
	"fmt"
	"strings"

	"github.com/starudream/go-lib/core/v2/utils/reflectutil"
)

func Nil(t T, value any, msgAndArgs ...any) bool {
	t.Helper()
	if reflectutil.IsNil(value) {
		return true
	}
	return Fail(t, fmt.Sprintf("expected nil, but got [%T] %v", value, value), msgAndArgs...)
}

func NotNil(t T, value any, msgAndArgs ...any) bool {
	t.Helper()
	if !reflectutil.IsNil(value) {
		return true
	}
	return Fail(t, fmt.Sprintf("expected not nil, but got [%T] %v", value, value), msgAndArgs...)
}

func Equal(t T, expected, actual any, msgAndArgs ...any) bool {
	t.Helper()
	if err := validateEqualArgs(expected, actual); err != nil {
		return Fail(t, fmt.Sprintf("invalid operation: %#v == %#v (%v)", expected, actual, err), msgAndArgs...)
	}
	if !reflectutil.Equal(expected, actual) {
		return Fail(t, fmt.Sprintf("should be equal to %T\n%s", expected, diff(expected, actual, strings.Repeat(" ", 2), "")), msgAndArgs...)
	}
	return true
}

func NotEqual(t T, expected, actual any, msgAndArgs ...any) bool {
	t.Helper()
	if err := validateEqualArgs(expected, actual); err != nil {
		return Fail(t, fmt.Sprintf("invalid operation: %#v == %#v (%s)", expected, actual, err), msgAndArgs...)
	}
	if reflectutil.Equal(expected, actual) {
		return Fail(t, fmt.Sprintf("should not be equal to %T", expected), msgAndArgs...)
	}
	return true
}
