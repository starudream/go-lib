package ntfy

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"
)

func TestDingtalk(t *testing.T) {
	testutil.Nil(t, _c.DingtalkConfig.Notify(ctx, text))
}
