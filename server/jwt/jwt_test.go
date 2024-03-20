package jwt_test

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"
	"github.com/starudream/go-lib/server/v2/jwt"
)

func Test(t *testing.T) {
	s1, e1 := jwt.Sign("foo", "bar", "hello", jwt.WithId("123"), jwt.WithMetadata(map[string]string{"abc": "123"}))
	testutil.LogNoErr(t, e1, s1)

	c1, e2 := jwt.Parse(s1)
	testutil.LogNoErr(t, e2, c1)
}
