package middlewares

import (
	"strings"

	"github.com/starudream/go-lib/server/v2/http"
	"github.com/starudream/go-lib/server/v2/iconst"
	"github.com/starudream/go-lib/server/v2/ierr"
	"github.com/starudream/go-lib/server/v2/jwt"
)

func JWT() http.Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerCtx(func(c *http.Context) error {
			_, err := jwt.FromContext(c)
			if err != nil {
				return err
			}

			next.ServeHTTP(c.Res, c.Req)

			return nil
		})
	}
}

func parse(c *http.Context) (jwt.Interface, error) {
	raw := c.Req.Header.Get(iconst.HeaderAuthorization)
	if raw == "" {
		return nil, ierr.Unauthorized(9999, "missing authorization header")
	}
	claims, err := jwt.Parse(strings.TrimPrefix(raw, "Bearer "))
	if err != nil {
		return nil, err
	}
	return claims, nil
}
