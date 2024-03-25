package middlewares

import (
	"net/http"
)

type (
	Handler     = http.Handler
	HandlerFunc = http.HandlerFunc

	Middleware = func(next Handler) Handler
)
