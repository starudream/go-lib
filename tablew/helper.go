package tablew

import (
	"bytes"
	"strings"
)

func Render(cb func(w *Table)) string {
	buf := &bytes.Buffer{}
	w := NewWriter(buf)
	cb(w)
	w.Render()
	return buf.String()
}

func Structs(v any) string {
	return Render(func(w *Table) {
		err := w.SetStructs(v)
		if err != nil {
			panic(err)
		}
	})
}

type TableCell interface {
	TableCellString() string
}

type fieldTag struct {
	name   string
	ignore bool
}

func genFieldTag(s string) (t fieldTag) {
	if s == "" {
		return
	}

	ss := strings.Split(s, ",")

	if sl := len(ss); sl == 0 {
		return
	} else if sl >= 1 {
		t.name = ss[0]
	}

	for _, v := range ss[1:] {
		switch strings.ToLower(strings.TrimSpace(v)) {
		case "ignore":
			t.ignore = true
		}
	}

	return
}
