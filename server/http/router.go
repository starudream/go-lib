package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Router interface {
	Handler

	Use(middlewares ...Middleware)
	With(middlewares ...Middleware) Router

	Group(fn RouterFunc) Router
	Route(pattern string, fn RouterFunc) Router
	Mount(pattern string, h Handler)

	Handle(pattern string, h Handler)
	HandleFunc(pattern string, h HandlerFunc)
	HandleCtx(pattern string, h HandlerCtx)
}

type RouterFunc = func(r Router)

type Mux struct {
	mux *chi.Mux
}

func NewMux() *Mux {
	return &Mux{chi.NewMux()}
}

var _ Router = (*Mux)(nil)

func (m *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.mux.ServeHTTP(w, r)
}

func (m *Mux) Use(middlewares ...Middleware) {
	m.mux.Use(middlewares...)
}

func (m *Mux) With(middlewares ...Middleware) Router {
	return &Mux{m.mux.With(middlewares...).(*chi.Mux)}
}

func (m *Mux) Group(fn RouterFunc) Router {
	nm := m.With()
	fn(nm)
	return nm
}

func (m *Mux) Route(pattern string, fn RouterFunc) Router {
	nm := NewMux()
	fn(nm)
	m.Mount(pattern, nm)
	return nm
}

func (m *Mux) Mount(pattern string, h Handler) {
	m.mux.Mount(pattern, h)
}

func (m *Mux) Handle(pattern string, h Handler) {
	m.mux.Handle(pattern, h)
}

func (m *Mux) HandleFunc(pattern string, h HandlerFunc) {
	m.mux.Handle(pattern, h)
}

func (m *Mux) HandleCtx(pattern string, h HandlerCtx) {
	m.mux.Handle(pattern, h)
}
