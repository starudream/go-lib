package handler

import (
	"log/slog"

	"github.com/starudream/go-lib/core/v2/config/global"

	"github.com/starudream/go-lib/core/v2/internal/logs"
	"github.com/starudream/go-lib/core/v2/slog/filewriter"
)

func NewFile(cfg global.Config) slog.Handler {
	if !cfg.LogFileEnabled {
		return nil
	}

	file := &filewriter.Logger{
		Filename:   cfg.LogFileFilename,
		MaxSize:    cfg.LogFileMaxSize,
		MaxAge:     cfg.LogFileMaxAge,
		MaxBackups: cfg.LogFileMaxBackups,
		LocalTime:  true,
		Compress:   true,
	}

	if cfg.LogFileDailyRotate {
		go filewriter.RotateDaily(file)
	}

	logs.D("log file enabled", "format", cfg.LogFileFormat, "level", cfg.LogFileLevel, "file", file.Filename)

	if cfg.LogFileFormat == "json" {
		return logs.NewJSONHandler(file, true, cfg.LogFileLevel)
	}

	return logs.NewConsoleWriterHandler(file, true, cfg.LogFileLevel, false)
}
