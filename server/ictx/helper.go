package ictx

import (
	"context"

	"github.com/starudream/go-lib/server/v2/iconst"
	"github.com/starudream/go-lib/server/v2/jwt"
)

func SYSTEM(args ...any) *Context {
	var (
		ctx = context.Background()
		str = [3]string{"*", "SYSTEM", ""}
		idx = 0

		c jwt.Interface
	)
	for i := 0; i < len(args); i++ {
		switch v := args[i].(type) {
		case jwt.Interface:
			c = v
		case context.Context:
			ctx = v
		case string:
			if idx < 3 {
				str[idx] = v
			}
			idx++
		}
	}
	if c == nil {
		c = jwt.New(str[0], str[1], str[2])
	}
	raw := c.Privilege()
	return FromContext(c.WithContext(ctx)).Set(iconst.HeaderAuthorization, raw)
}
