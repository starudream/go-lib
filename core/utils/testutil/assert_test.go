package testutil

import (
	"fmt"
	"testing"

	"github.com/starudream/go-lib/core/v2/internal/testdata"
)

func TestNil(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		Nil(t, error(nil), "hello")
	})
	t.Run("not nil", func(t *testing.T) {
		NotNil(t, fmt.Errorf("no error"), "hello")
	})

	t.Run("error", func(t *testing.T) {
		mt := new(MockT)
		Equal(t, false, NotNil(mt, error(nil), "hello"))
	})
}

func TestEqual(t *testing.T) {
	t.Run("v1 ne v2", func(t *testing.T) {
		NotEqual(t, testdata.V1, testdata.V2)
	})
	t.Run("v1 eq v3", func(t *testing.T) {
		Equal(t, testdata.V1, testdata.V3)
	})

	t.Run("error1", func(t *testing.T) {
		mt := new(MockT)
		Equal(t, false, Equal(mt, testdata.V1, testdata.V2))
	})
	t.Run("error2", func(t *testing.T) {
		mt := new(MockT)
		Equal(t, false, NotEqual(mt, testdata.V1, testdata.V3))
	})
}
