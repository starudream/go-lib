package handler

import (
	"log/slog"
	"os"

	"github.com/starudream/go-lib/core/v2/config/global"

	"github.com/starudream/go-lib/core/v2/internal/logs"
)

func NewConsole(cfg global.Config) slog.Handler {
	if cfg.LogConsoleDisabled {
		return nil
	}

	logs.D("log console enabled", "format", cfg.LogConsoleFormat, "level", cfg.LogConsoleLevel)

	if cfg.LogConsoleFormat == "json" {
		return logs.NewJSONHandler(os.Stdout, true, cfg.LogConsoleLevel)
	}

	return logs.NewColorableStdoutHandler(cfg.LogConsoleLevel)
}
