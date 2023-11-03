package level

import (
	"encoding"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type Level int

const (
	Debug Level = -4
	Info  Level = 0
	Warn  Level = 4
	Error Level = 8
	Fatal Level = 16
)

var Strings = map[Level]string{
	Debug: "DEBUG",
	Info:  "INFO",
	Warn:  "WARN",
	Error: "ERROR",
	Fatal: "FATAL",
}

//goland:noinspection GoMixedReceiverTypes
func (l Level) String() string {
	if s, ok := Strings[l]; ok {
		return s
	}
	return "UNKNOWN"
}

var ShortStrings = map[Level]string{
	Debug: "DBG",
	Info:  "INF",
	Warn:  "WRN",
	Error: "ERR",
	Fatal: "FTL",
}

//goland:noinspection GoMixedReceiverTypes
func (l Level) ShortString() string {
	if s, ok := ShortStrings[l]; ok {
		return s
	}
	return "???"
}

var _ json.Marshaler = Level(0)

//goland:noinspection GoMixedReceiverTypes
func (l Level) MarshalJSON() ([]byte, error) {
	return strconv.AppendQuote(nil, l.String()), nil
}

var _ json.Unmarshaler = (*Level)(nil)

//goland:noinspection GoMixedReceiverTypes
func (l *Level) UnmarshalJSON(bs []byte) error {
	s, err := strconv.Unquote(string(bs))
	if err != nil {
		return err
	}
	return l.parse(s)
}

var _ encoding.TextMarshaler = Level(0)

//goland:noinspection GoMixedReceiverTypes
func (l Level) MarshalText() (bs []byte, err error) {
	return []byte(l.String()), nil
}

var _ encoding.TextUnmarshaler = (*Level)(nil)

//goland:noinspection GoMixedReceiverTypes
func (l *Level) UnmarshalText(bs []byte) error {
	return l.parse(string(bs))
}

//goland:noinspection GoMixedReceiverTypes
func (l *Level) parse(s string) (err error) {
	switch strings.ToUpper(strings.TrimSpace(s)) {
	case "DEBUG", "DBG", "D":
		*l = Debug
	case "INFO", "INF", "I", "":
		*l = Info
	case "WARN", "WAN", "W":
		*l = Warn
	case "ERROR", "ERR", "E":
		*l = Error
	case "FATAL", "FTL", "F":
		*l = Fatal
	default:
		err = fmt.Errorf("unknown level string %q", s)
	}
	return
}
