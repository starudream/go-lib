package ictx_test

import (
	"context"
	"testing"

	"google.golang.org/grpc/metadata"

	"github.com/starudream/go-lib/core/v2/utils/testutil"
	"github.com/starudream/go-lib/server/v2/ictx"
)

func TestContext(t *testing.T) {
	md := metadata.Pairs("k1", "v1", "k2", "v2")
	md.Append("k3", "a", "b", "c")
	testutil.Log(t, md)

	c1 := metadata.NewIncomingContext(context.Background(), md)
	testutil.Log(t, c1)

	c2 := ictx.FromContext(c1).Set("k4", "v4").Append("k5", "a", "b", "c")
	testutil.Log(t, c2, c2.Get("k1"), c2.Get("k5"))
}
