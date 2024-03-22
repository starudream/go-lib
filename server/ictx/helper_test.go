package ictx

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"
	"github.com/starudream/go-lib/server/v2/jwt"
)

func TestSYSTEM(t *testing.T) {
	c1 := SYSTEM()
	testutil.Log(t, c1)

	c2 := SYSTEM("foo", "bar", "zoo")
	testutil.Log(t, c2)

	c3 := jwt.MustFromContext(c2)
	testutil.Log(t, c3.ISS(), c3.SUB(), c3.AUD())
}
