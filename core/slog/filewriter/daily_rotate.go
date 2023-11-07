package filewriter

import (
	"fmt"
	"time"

	"github.com/starudream/go-lib/core/v2/utils/osutil"
	"github.com/starudream/go-lib/core/v2/utils/signalutil"

	"github.com/starudream/go-lib/core/v2/internal/logs"
)

func RotateDaily(logger *Logger) {
	if logger == nil {
		return
	}

	fn := func() {
		var err error
		defer func() {
			if r := recover(); r != nil && err == nil {
				err = fmt.Errorf("[%s] %v", osutil.CallerString(1), r)
			}
			if err != nil {
				logs.D("log file rotate daily failed", "err", err)
			} else {
				logs.D("log file rotate daily success")
			}
		}()
		logs.D("attempt to rotate log file daily")
		err = logger.Rotate()
	}

	for {
		now := time.Now()
		tom := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location())

		select {
		// case <-time.After(time.Second):
		// 	fn()
		case <-time.After(tom.Sub(now)):
			fn()
		case <-signalutil.Done():
			_ = logger.Close()
			return
		}
	}
}
