package ping

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"
)

func TestPing(t *testing.T) {
	stat, err := Ping(WithAddr("www.google.com"))
	testutil.LogNoErr(t, err, stat)
}
