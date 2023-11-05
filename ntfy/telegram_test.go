package ntfy

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"
)

func TestTelegram(t *testing.T) {
	testutil.Nil(t, _c.TelegramConfig.Notify(ctx, text))
}
