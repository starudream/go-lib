package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/starudream/go-lib/core/v2/slog"
)

type (
	Handler     = http.Handler
	HandlerFunc func(c *Context) error

	Middleware = func(next Handler) Handler
)

type Router interface {
	Handler

	Use(middlewares ...Middleware)
	With(middlewares ...Middleware) Router
	Group(fn func(r Router)) Router
	Route(pattern string, fn func(r Router)) Router
	Mount(pattern string, h Handler)

	Handle(pattern string, h Handler)
	HandleFunc(pattern string, h HandlerFunc)
	Method(method, pattern string, h Handler)
	MethodFunc(method, pattern string, h HandlerFunc)
	Connect(pattern string, h HandlerFunc)
	Delete(pattern string, h HandlerFunc)
	Get(pattern string, h HandlerFunc)
	Head(pattern string, h HandlerFunc)
	Options(pattern string, h HandlerFunc)
	Patch(pattern string, h HandlerFunc)
	Post(pattern string, h HandlerFunc)
	Put(pattern string, h HandlerFunc)
	Trace(pattern string, h HandlerFunc)
}

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

func (m *Mux) Group(fn func(r Router)) Router {
	t := m.With()
	fn(t)
	return t
}

func (m *Mux) Route(pattern string, fn func(r Router)) Router {
	t := NewMux()
	fn(t)
	m.Mount(pattern, t)
	return t
}

func (m *Mux) Mount(pattern string, h Handler) {
	m.mux.Mount(pattern, h)
}

func (m *Mux) Handle(pattern string, h Handler) {
	m.mux.Handle(pattern, h)
}

func (m *Mux) HandleFunc(pattern string, h HandlerFunc) {
	m.mux.HandleFunc(pattern, m.hf(h))
}

func (m *Mux) Method(method, pattern string, h Handler) {
	m.mux.Method(method, pattern, h)
}

func (m *Mux) MethodFunc(method, pattern string, h HandlerFunc) {
	m.mux.MethodFunc(method, pattern, m.hf(h))
}

func (m *Mux) Connect(pattern string, h HandlerFunc) {
	m.mux.Connect(pattern, m.hf(h))
}

func (m *Mux) Delete(pattern string, h HandlerFunc) {
	m.mux.Delete(pattern, m.hf(h))
}

func (m *Mux) Get(pattern string, h HandlerFunc) {
	m.mux.Get(pattern, m.hf(h))
}

func (m *Mux) Head(pattern string, h HandlerFunc) {
	m.mux.Head(pattern, m.hf(h))
}

func (m *Mux) Options(pattern string, h HandlerFunc) {
	m.mux.Options(pattern, m.hf(h))
}

func (m *Mux) Patch(pattern string, h HandlerFunc) {
	m.mux.Patch(pattern, m.hf(h))
}

func (m *Mux) Post(pattern string, h HandlerFunc) {
	m.mux.Post(pattern, m.hf(h))
}

func (m *Mux) Put(pattern string, h HandlerFunc) {
	m.mux.Put(pattern, m.hf(h))
}

func (m *Mux) Trace(pattern string, h HandlerFunc) {
	m.mux.Trace(pattern, m.hf(h))
}

var ErrorHandler func(c *Context, err error)

func (m *Mux) hf(h HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := NewContext(m, w, r)
		err := h(c)
		if err != nil {
			if ErrorHandler == nil {
				slog.Warn("http handler error: %v", err)
				w.Header().Set("x-server-message", err.Error())
				w.WriteHeader(500)
			}
			ErrorHandler(c, err)
		}
	}
}
