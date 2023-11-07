package ntfy

import (
	"strings"
)

func clean(s string) string {
	return strings.ToUpper(strings.TrimSpace(s))
}
