package testutil

import (
	"fmt"
	"testing"

	"github.com/starudream/go-lib/core/v2/internal/testdata"
)

func TestMustNil(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		MustNil(t, error(nil), "hello")
	})
	t.Run("not nil", func(t *testing.T) {
		MustNotNil(t, fmt.Errorf("no error"), "hello")
	})

	t.Run("error", func(t *testing.T) {
		mt := new(MockT)
		MustNotNil(mt, error(nil))
		MustEqual(t, true, mt.Failed)
	})
}

func TestMustEqual(t *testing.T) {
	t.Run("v1 ne v2", func(t *testing.T) {
		mt := new(MockT)
		MustNotEqual(t, testdata.V1, testdata.V2)
		MustEqual(t, false, mt.Failed)
	})
	t.Run("v1 eq v3", func(t *testing.T) {
		mt := new(MockT)
		MustEqual(t, testdata.V1, testdata.V3)
		MustEqual(t, false, mt.Failed)
	})

	t.Run("error1", func(t *testing.T) {
		mt := new(MockT)
		MustEqual(mt, testdata.V1, testdata.V2)
		MustEqual(t, true, mt.Failed)
	})
	t.Run("error2", func(t *testing.T) {
		mt := new(MockT)
		MustNotEqual(mt, testdata.V1, testdata.V3)
		MustEqual(t, true, mt.Failed)
	})
}
