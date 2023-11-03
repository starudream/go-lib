package record

import (
	"context"
	"log/slog"
	"strings"
	"time"

	"github.com/starudream/go-lib/core/v2/slog/level"
	"github.com/starudream/go-lib/core/v2/utils/osutil"
)

func Handle(os ...OptionI) {
	opts := &Options{
		ctx:    context.Background(),
		logger: slog.Default(),
		time:   time.Now(),
		level:  level.Debug,
	}
	for i := 0; i < len(os); i++ {
		os[i].apply(opts)
	}

	if !opts.logger.Enabled(opts.ctx, opts.level.Level()) {
		return
	}

	var pc uintptr
	if len(opts.skipNames) > 0 {
		pc = osutil.CallerFn(func(frame osutil.CallerFrame) bool {
			for _, name := range opts.skipNames {
				if strings.Contains(frame.File, name) {
					return true
				}
			}
			return false
		})
	} else {
		pc = osutil.CallerPC(opts.skip)
	}

	record := slog.NewRecord(opts.time, opts.level.Level(), opts.msg, pc)
	if len(opts.attrs) > 0 {
		record.AddAttrs(opts.attrs...)
	}
	_ = opts.logger.Handler().Handle(opts.ctx, record)
}
