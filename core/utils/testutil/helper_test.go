package testutil

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/internal/testdata"
)

func TestDiff(t *testing.T) {
	x := diff(testdata.V1, testdata.V2, "", "")
	Log(t, "\n"+x)
	MustEqual(t, "A -> \"hello world\" != \"\"\nh -> \"\" != \"hidden\"", x)
}
