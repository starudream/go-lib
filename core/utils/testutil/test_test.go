package testutil

import (
	"fmt"
	"strings"
)

type MockT struct {
	Failed bool
}

var _ T = (*MockT)(nil)

func (t *MockT) Helper() {
}

func (t *MockT) Log(a ...any) {
	fmt.Println(strings.Trim(fmt.Sprint(a...), "\n"))
}

func (t *MockT) Error(a ...any) {
	fmt.Println(strings.Trim(fmt.Sprint(a...), "\n"))
}

func (t *MockT) FailNow() {
	t.Failed = true
}
