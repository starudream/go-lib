package tablew

import (
	"bytes"
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
