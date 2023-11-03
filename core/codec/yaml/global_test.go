package yaml

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"

	"github.com/starudream/go-lib/core/v2/internal/testdata"
)

func TestG(t *testing.T) {
	s1, err := MarshalString(testdata.V1)
	testutil.LogNoErr(t, err, s1)

	var v4 = &testdata.V{}

	err = UnmarshalString(s1, v4)
	testutil.LogNoErr(t, err, v4)

	v5, err := UnmarshalTo[*testdata.V](s1)
	testutil.LogNoErr(t, err, v5)

	vs1, err := UnmarshalTo[[]*testdata.V](MustMarshal([]*testdata.V{testdata.V1}))
	testutil.LogNoErr(t, err, vs1)
}
