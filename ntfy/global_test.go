package ntfy

import (
	"context"
	"testing"

	"github.com/starudream/go-lib/core/v2/codec/yaml"
	"github.com/starudream/go-lib/core/v2/config"
	"github.com/starudream/go-lib/core/v2/utils/testutil"
)

var (
	ctx  = context.Background()
	text = "hello world"
)

func TestConfig(t *testing.T) {
	testutil.Log(t, "\n"+yaml.MustMarshalString(C()))

	testutil.Log(t, "\n"+yaml.MustMarshalString(config.Raw("ntfy")))
}

func TestNotify(t *testing.T) {
	testutil.Nil(t, Notify(ctx, text))
}
