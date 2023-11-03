package osutil

import (
	"os"
	"strings"
	"sync"
)

var (
	_test     bool
	_testOnce sync.Once
)

func ArgTest() bool {
	fn := func() bool {
		for _, arg := range os.Args {
			if strings.HasPrefix(arg, "-test.") {
				return true
			}
		}
		return false
	}
	_testOnce.Do(func() {
		_test = fn()
	})
	return _test
}
