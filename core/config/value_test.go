package config

import (
	"testing"
)

func TestValue(t *testing.T) {
	t.Log(NewValue("X").String("Y"))
	t.Log(NewValue("").String("Y"))
	t.Log(NewValue(nil).String("X"))
}
