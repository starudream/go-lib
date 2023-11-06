package testutil

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/kr/pretty"

	"github.com/starudream/go-lib/core/v2/utils/reflectutil"
)

func msgAndArgsString(a ...any) string {
	if len(a) == 0 {
		return ""
	}

	f := ""
	if s, ok := a[0].(string); ok && strings.Contains(s, "%") {
		f = s
		a = a[1:]
	}

	if len(a) == 0 {
		if f == "" {
			return "<EMPTY STRING>"
		}
		return fmt.Sprint(f)
	}
	if f == "" {
		f = strings.Repeat("%v\n", len(a)-1) + "%v"
		for i := 0; i < len(a); i++ {
			a[i] = pretty.Sprint(a[i])
		}
	}
	return fmt.Sprintf(f, a...)
}

func validateEqualArgs(expected, actual any) error {
	if expected == nil && actual == nil {
		return nil
	}
	if reflectutil.IsFunc(expected) || reflectutil.IsFunc(actual) {
		return fmt.Errorf("cannot compare functions: %T %T", expected, actual)
	}
	return nil
}

func diff(expected, actual any, prefix, indent string) string {
	ds := pretty.Diff(expected, actual)
	rs := make([][2]string, len(ds))
	for i, d := range ds {
		ss := strings.SplitN(d, ": ", 2)
		if len(ss) < 2 {
			ss = append(ss, "", "")
		}
		rs[i] = [2]string{ss[0], ss[1]}
	}
	return render(prefix, indent, " -> ", rs)
}

func render(prefix, indent, sep string, rows [][2]string) string {
	var (
		ma  = 0
		ks  = make([]string, len(rows))
		vs  = make([]string, len(rows))
		buf = &bytes.Buffer{}
	)
	for i, row := range rows {
		if l := len(row[0]); l > ma {
			ma = l
		}
		ks[i], vs[i] = row[0], row[1]
	}
	for i := 0; i < len(ks); i++ {
		if i > 0 {
			buf.WriteString("\n")
			buf.WriteString(indent)
		}
		buf.WriteString(prefix)
		buf.WriteString(ks[i])
		if vs[i] == "" {
			continue
		}
		buf.WriteString(strings.Repeat(" ", ma-len(ks[i])))
		buf.WriteString(sep)
		ss := strings.Split(vs[i], "\n")
		for j := 0; j < len(ss); j++ {
			if j > 0 {
				buf.WriteString("\n")
				buf.WriteString(strings.Repeat(" ", len(prefix)+len(indent)+ma+len(sep)))
			}
			buf.WriteString(ss[j])
		}
	}
	return buf.String()
}
