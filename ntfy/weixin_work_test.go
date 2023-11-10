package ntfy

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"
)

func TestWeixinWork(t *testing.T) {
	testutil.Nil(t, _c.WeixinWorkConfig.Notify(ctx, text))
}
