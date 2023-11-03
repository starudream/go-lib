package logs

import (
	"io"
	"log/slog"
	"os"

	"github.com/lmittmann/tint"
	"github.com/mattn/go-colorable"
	"github.com/mattn/go-isatty"

	"github.com/starudream/go-lib/core/v2/slog/level"
	"github.com/starudream/go-lib/core/v2/utils/osutil"
)

func NewConsoleWriterHandler(writer io.Writer, source bool, level level.Leveler, color bool) slog.Handler {
	return tint.NewHandler(writer, &tint.Options{
		AddSource:   source,
		Level:       level,
		TimeFormat:  TimeFormat,
		NoColor:     !color,
		ReplaceAttr: NewAttrReplacer(color, true),
	})
}

func NewJSONHandler(writer io.Writer, source bool, level level.Leveler) slog.Handler {
	return slog.NewJSONHandler(writer, &slog.HandlerOptions{
		AddSource:   source,
		Level:       level,
		ReplaceAttr: NewAttrReplacer(false, false),
	})
}

func NewColorableStdoutHandler(level level.Leveler) slog.Handler {
	writer, color := colorable.NewColorableStdout(), true
	if f, ok := writer.(*os.File); ok {
		color = isatty.IsTerminal(f.Fd())
	}
	return NewConsoleWriterHandler(writer, osutil.DOT(), level, color || osutil.ArgTest())
}
