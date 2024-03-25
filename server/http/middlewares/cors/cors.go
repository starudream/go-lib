package cors

import (
	"net/http"

	"github.com/rs/cors"

	"github.com/starudream/go-lib/server/v2/http/middlewares"
)

type Options = cors.Options

var New = cors.New

func AllowAll() middlewares.Middleware {
	c := New(Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodHead,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
	})
	return c.Handler
}
