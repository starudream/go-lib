package tablew

import (
	"fmt"
	"testing"
)

type T struct {
	A string `json:"a" table:",ignore"`
	B TB     `table:"foo"`
}

type TB int

func (v TB) TableCellString() string {
	return fmt.Sprintf("%02d", v)
}

func TestStructs(t *testing.T) {
	ts := []T{
		{"a", 1},
		{"b", 2},
	}
	t.Log("\n" + Structs(ts))
}
