package gh

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
