package tablew

import (
	"bytes"
	"reflect"
	"strings"
)

type Option func(w *Table)

func Render(opts ...Option) string {
	buf := &bytes.Buffer{}
	w := NewWriter(buf)
	for i := 0; i < len(opts); i++ {
		opts[i](w)
	}
	w.Render()
	return buf.String()
}

func Structs(vs any, opts ...Option) string {
	if vs == nil {
		return "<nil>"
	}
	vt := reflect.TypeOf(vs)
	if vt.Kind() != reflect.Array && vt.Kind() != reflect.Slice {
		panic("must be array or slice")
	}
	vv := reflect.ValueOf(vs)
	if vv.Len() < 1 {
		return "<empty>"
	}
	fn := func(w *Table) {
		err := w.SetStructs(vs)
		if err != nil {
			panic(err)
		}
	}
	return Render(append(opts, fn)...)
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
