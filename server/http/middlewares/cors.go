package middlewares

import (
	"github.com/rs/cors"

	"github.com/starudream/go-lib/server/v2/http"
)

func CORS() http.Middleware {
	opts := cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			"GET",
			"HEAD",
			"POST",
			"PUT",
			"PATCH",
			"DELETE",
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
	}
	return cors.New(opts).Handler
}
