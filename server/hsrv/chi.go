package hsrv

import (
	"github.com/go-chi/chi/v5"
)

type (
	Route       = chi.Route
	Context     = chi.Context
	RouteParams = chi.RouteParams
	Mux         = chi.Mux
	Router      = chi.Router
	Routes      = chi.Routes
	Middlewares = chi.Middlewares
)

var (
	NewRouter       = chi.NewRouter
	NewRouteContext = chi.NewRouteContext
	RouteContext    = chi.RouteContext
	URLParam        = chi.URLParam
	URLParamFromCtx = chi.URLParamFromCtx
)
