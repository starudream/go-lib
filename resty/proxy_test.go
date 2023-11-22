package resty

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"
)

func TestHasProxyEnv(t *testing.T) {
	testutil.Log(t, HasProxyEnv())
}
