package logs

import (
	"log"
	"log/slog"

	"github.com/starudream/go-lib/core/v2/slog/level"
	"github.com/starudream/go-lib/core/v2/utils/loutil"
	"github.com/starudream/go-lib/core/v2/utils/osutil"
)

const TimeFormat = "2006-01-02T15:04:05.000Z07:00"

var _logger = setDefault(New(NewColorableStdoutHandler(loutil.Ternary(osutil.DOT(), level.Debug, level.Info))))

var D = _logger.Debug

func SetDefault(logger *slog.Logger) *slog.Logger {
	setDefault(logger)
	_logger = logger
	return logger
}

func setDefault(logger *slog.Logger) *slog.Logger {
	log.SetOutput(slog.NewLogLogger(logger.Handler(), level.Debug.Level()).Writer())
	slog.SetDefault(logger)
	return logger
}

var New = slog.New
