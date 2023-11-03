package level

import (
	"fmt"
)

const (
	colorBlack = iota + 30
	colorRed
	colorGreen
	colorYellow
	colorBlue
	colorMagenta
	colorCyan
	colorWhite
)

var Colors = map[Level]int{
	Debug: 0,
	Info:  colorGreen,
	Warn:  colorYellow,
	Error: colorRed,
	Fatal: colorRed,
}

func Colorize(s string, c int, color bool) string {
	if !color {
		return s
	}
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", c, s)
}

//goland:noinspection GoMixedReceiverTypes
func (l Level) Color() int {
	return Colors[l]
}

//goland:noinspection GoMixedReceiverTypes
func (l Level) ColorString(color, short bool) string {
	if short {
		return Colorize(l.ShortString(), l.Color(), color)
	}
	return Colorize(l.String(), l.Color(), color)
}
