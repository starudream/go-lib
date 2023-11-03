package config

import (
	"testing"
)

func TestFiles(t *testing.T) {
	files := Files()
	for i := 0; i < len(files); i++ {
		t.Log(files[i])
	}
}
