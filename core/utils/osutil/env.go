package osutil

import (
	"os"
	"strconv"
	"strings"
	"sync"
)

var (
	_debug     bool
	_debugOnce sync.Once
)

func EnvDebug() bool {
	_debugOnce.Do(func() {
		_debug, _ = strconv.ParseBool(strings.ToLower(os.Getenv("DEBUG")))
	})
	return _debug
}

// DOT returns true if DEBUG is enabled or test mode is enabled.
func DOT() bool {
	return EnvDebug() || ArgTest()
}
