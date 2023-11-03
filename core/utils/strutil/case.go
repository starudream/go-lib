package strutil

import (
	"regexp"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var camelPattern = regexp.MustCompile("[^A-Z][A-Z]+")

func ToSnakeCase(s string) string {
	if !strings.ContainsRune(s, '_') {
		s = camelPattern.ReplaceAllStringFunc(s, func(x string) string {
			return x[:1] + "_" + x[1:]
		})
	}
	return strings.ToLower(s)
}

func ToCamelCase(s string) string {
	return toCamelInitCase(s, false)
}

func ToPascalCase(s string) string {
	return toCamelInitCase(s, true)
}

func toCamelInitCase(name string, initUpper bool) string {
	out := ""
	for i, p := range strings.Split(name, "_") {
		if !initUpper && i == 0 {
			out += p
			continue
		}
		out += cases.Title(language.Und).String(p)
	}
	return out
}
