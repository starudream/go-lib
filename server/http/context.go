package http

import (
	"bytes"
	"context"
	"encoding/xml"
	"errors"
	"html/template"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/go-chi/chi/v5"

	"github.com/starudream/go-lib/core/v2/codec/json"
	"github.com/starudream/go-lib/server/v2/iconst"
)

type Context struct {
	mux *Mux

	Req *http.Request
	Res http.ResponseWriter

	chi   *chi.Context
	kvs   map[any]any
	body  []byte
	query url.Values
}

func NewContext(m *Mux, w http.ResponseWriter, r *http.Request) *Context {
	c := &Context{
		mux: m,
		Req: r,
		Res: newResponseWriter(w),
		chi: chi.RouteContext(r.Context()),
		kvs: map[any]any{},
	}
	return c
}

// --- Req

func (c *Context) GetParam(key string) string {
	return c.chi.URLParam(key)
}

func (c *Context) GetParams() map[string]string {
	m := map[string]string{}
	for i, k := range c.chi.URLParams.Keys {
		m[k] = c.chi.URLParams.Values[i]
	}
	return m
}

func (c *Context) GetQuery(key string) string {
	if c.query == nil {
		c.query = c.Req.URL.Query()
	}
	return c.query.Get(key)
}

func (c *Context) GetQueries(key string) []string {
	if c.query == nil {
		c.query = c.Req.URL.Query()
	}
	return c.query[key]
}

func (c *Context) GetHeader(key string) string {
	return c.Req.Header.Get(key)
}

func (c *Context) GetHeaders(key string) []string {
	return c.Req.Header.Values(key)
}

func (c *Context) GetContentType() string {
	return filterFlags(c.GetHeader(iconst.HeaderContentType))
}

func (c *Context) GetRawBody() ([]byte, error) {
	if c.body == nil {
		b, err := io.ReadAll(c.Req.Body)
		if err != nil {
			return nil, err
		}
		c.Req.Body = io.NopCloser(bytes.NewReader(b))
		c.body = b
	}
	return c.body, nil
}

// --- Req Bind

func (c *Context) BindQuery(obj any) error {
	if c.query == nil {
		c.query = c.Req.URL.Query()
	}
	return setAny(obj, "query", c.query)
}

func (c *Context) BindHeader(obj any) error {
	return setAny(obj, "header", c.Req.Header)
}

const defaultMaxMemory = 32 << 20 // 32 MB

func (c *Context) BindForm(obj any) error {
	if err := c.Req.ParseMultipartForm(defaultMaxMemory); err != nil && !errors.Is(err, http.ErrNotMultipart) {
		return err
	}
	return setAny(obj, "form", c.Req.Form)
}

func (c *Context) BindPostForm(obj any) error {
	if err := c.Req.ParseMultipartForm(defaultMaxMemory); err != nil && !errors.Is(err, http.ErrNotMultipart) {
		return err
	}
	return setAny(obj, "form", c.Req.PostForm)
}

func (c *Context) BindMultiForm(obj any) error {
	if err := c.Req.ParseMultipartForm(defaultMaxMemory); err != nil {
		return err
	}
	return setMulti(obj, "form", c.Req.MultipartForm)
}

func (c *Context) BindJSON(obj any) error {
	body, err := c.GetRawBody()
	if err != nil {
		return err
	}
	return json.Unmarshal(body, obj)
}

func (c *Context) BindXML(obj any) error {
	body, err := c.GetRawBody()
	if err != nil {
		return err
	}
	return xml.Unmarshal(body, obj)
}

// --- Resp

func (c *Context) Status(status int) {
	c.Res.WriteHeader(status)
}

func (c *Context) Header(key, value string) {
	if value == "" {
		c.Res.Header().Del(key)
	}
	c.Res.Header().Set(key, value)
}

// --- Resp Render

func (c *Context) Render(status int, contentType string, bs []byte) error {
	c.Header(iconst.HeaderContentType, contentType)
	c.Status(status)
	_, err := c.Res.Write(bs)
	return err
}

func (c *Context) HTML(status int, tpl template.Template, data any, name ...string) error {
	c.Header(iconst.HeaderContentType, iconst.MIMETextHTMLCharsetUTF8)
	c.Status(status)
	if len(name) > 0 {
		return tpl.ExecuteTemplate(c.Res, name[0], data)
	}
	return tpl.Execute(c.Res, data)
}

func (c *Context) TEXT(status int, text string) error {
	return c.Render(status, iconst.MIMETextPlainCharsetUTF8, []byte(text))
}

func (c *Context) JSON(status int, v any) error {
	bs, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return c.Render(status, iconst.MIMEApplicationJSONCharsetUTF8, bs)
}

func (c *Context) XML(status int, v any) error {
	bs, err := xml.Marshal(v)
	if err != nil {
		return err
	}
	return c.Render(status, iconst.MIMEApplicationXMLCharsetUTF8, bs)
}

func (c *Context) Redirect(status int, url string) {
	http.Redirect(c.Res, c.Req, url, status)
}

// --- Context

var _ context.Context = (*Context)(nil)

func (c *Context) Deadline() (deadline time.Time, ok bool) {
	return c.Req.Context().Deadline()
}

func (c *Context) Done() <-chan struct{} {
	return c.Req.Context().Done()
}

func (c *Context) Err() error {
	return c.Req.Context().Err()
}

func (c *Context) Value(key any) any {
	if v, ok := c.kvs[key]; ok {
		return v
	}
	return c.Req.Context().Value(key)
}

// --- Values

func (c *Context) Get(key any) (any, bool) {
	v, ok := c.kvs[key]
	return v, ok
}

func (c *Context) MustGet(key any) any {
	v, ok := c.kvs[key]
	if ok {
		return v
	}
	panic("key not found")
}

func (c *Context) Set(key, value any) {
	c.kvs[key] = value
}
