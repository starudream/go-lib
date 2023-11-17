package filewriter

import (
	"github.com/starudream/go-lib/core/v2/utils/signalutil"

	"github.com/starudream/go-lib/core/v2/internal/logs"
)

func Sync(logger *Logger) {
	for {
		<-signalutil.Defer(func() { _ = logger.Close(); logs.D("log file closed") }).Done()
	}
}
