package ntfy

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"
)

func TestWebhook(t *testing.T) {
	testutil.Nil(t, _c.WebhookConfig.Notify(ctx, text))
}
