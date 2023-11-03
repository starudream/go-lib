package json

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"

	"github.com/starudream/go-lib/core/v2/internal/testdata"
)

func TestG(t *testing.T) {
	s1, err := MarshalString(testdata.V1)
	testutil.LogNoErr(t, err, s1)
	testutil.MustEqual(t, `{"a":"hello world","b":{"bool":true,"float":3.14,"int":16}}`, s1)

	s2, err := MarshalIndentString(testdata.V1)
	testutil.LogNoErr(t, err, s2)

	var v2 *testdata.V
	var v3 testdata.V
	var v4 = &testdata.V{}

	err = UnmarshalString(s1, v2)
	testutil.NotNil(t, err)

	err = UnmarshalString(s1, v3)
	testutil.NotNil(t, err)

	err = UnmarshalString(s1, v4)
	testutil.LogNoErr(t, err, v4)

	v5, err := UnmarshalTo[*testdata.V](s1)
	testutil.LogNoErr(t, err, v5)

	vs1, err := UnmarshalTo[[]*testdata.V](MustMarshal([]*testdata.V{testdata.V1}))
	testutil.LogNoErr(t, err, vs1)
}
