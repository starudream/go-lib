package testutil

type T interface {
	Helper()
	Log(...any)
	Error(...any)
	FailNow()
}

type _skipT struct {
	T
	skip int
}

func skipT(t T, skip ...int) _skipT {
	if len(skip) == 0 {
		skip = []int{1}
	}
	return _skipT{t, skip[0]}
}
