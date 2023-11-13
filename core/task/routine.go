package task

import (
	"github.com/starudream/go-lib/core/v2/slog"
	"github.com/starudream/go-lib/core/v2/utils/osutil"
)

func Run(fn func()) {
	defer func() {
		if r := recover(); r != nil {
			slog.Error("[%s] %v", osutil.CallerString(2), r, slog.String("stack", osutil.Stack(2)))
		}
	}()

	fn()
}

func Go(fn func()) {
	go Run(fn)
}
