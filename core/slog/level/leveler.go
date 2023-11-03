package level

import (
	"log/slog"
)

type Leveler = slog.Leveler

var _ Leveler = Level(0)

//goland:noinspection GoMixedReceiverTypes
func (l Level) Level() slog.Level {
	return slog.Level(l)
}
