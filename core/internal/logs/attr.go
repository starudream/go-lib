package logs

import (
	"log/slog"

	"github.com/starudream/go-lib/core/v2/slog/level"
	"github.com/starudream/go-lib/core/v2/utils/osutil"
)

func NewAttrReplacer(color, short bool) func(groups []string, attr slog.Attr) slog.Attr {
	return func(groups []string, attr slog.Attr) slog.Attr {
		if len(groups) > 0 {
			return attr
		}
		switch key := attr.Key; key {
		case slog.SourceKey:
			if source, ok := attr.Value.Any().(*slog.Source); ok {
				source.File = osutil.CallerFormatPath(source.File)
				return slog.Any(key, source)
			}
		case slog.LevelKey:
			if l, ok := attr.Value.Any().(slog.Level); ok {
				return slog.String(key, level.Level(l).ColorString(color, short))
			}
		}
		return attr
	}
}
