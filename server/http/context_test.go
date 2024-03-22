package http_test

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net"
	stdhttp "net/http"
	"strings"
	"testing"
	"time"

	"github.com/starudream/go-lib/core/v2/gh"
	"github.com/starudream/go-lib/core/v2/utils/osutil"
	"github.com/starudream/go-lib/core/v2/utils/testutil"
	"github.com/starudream/go-lib/resty/v2"
	"github.com/starudream/go-lib/server/v2/http"
)

func TestContext_Req(t *testing.T) {
	t.Run("param", With(
		func(t *testing.T, s *http.Server) {
			s.Get("/user/{name}", func(c *http.Context) error {
				return c.JSON(200, c.GetParams())
			})
			s.Get("/version/{other}/latest", func(c *http.Context) error {
				return c.JSON(200, c.GetParams())
			})
			s.Get("/page/*", func(c *http.Context) error {
				return c.JSON(200, c.GetParams())
			})
		},
		func(t *testing.T, c *resty.Client) {
			t.Run("match", func(t *testing.T) {
				resp, err := c.R().Get("/user/admin")
				testutil.Nil(t, err)
				testutil.Equal(t, 200, resp.StatusCode())
				testutil.Equal(t, `{"name":"admin"}`, string(resp.Body()))
			})
			t.Run("not_match", func(t *testing.T) {
				resp, err := c.R().Get("/version/foo/1")
				testutil.Nil(t, err)
				testutil.Equal(t, 404, resp.StatusCode())
			})
			t.Run("wildcard", func(t *testing.T) {
				resp, err := c.R().SetQueryParam("foo", "bar").Get("/page/a/b/c")
				testutil.Nil(t, err)
				testutil.Equal(t, 200, resp.StatusCode())
				testutil.Equal(t, `{"*":"a/b/c"}`, string(resp.Body()))
			})
		},
	))

	t.Run("query", With(
		func(t *testing.T, s *http.Server) {
			s.Get("/query", func(c *http.Context) error {
				return c.TEXT(200, c.GetQuery("foo"))
			})
		},
		func(t *testing.T, c *resty.Client) {
			resp, err := c.R().SetQueryParam("foo", "bar").Get("/query")
			testutil.Nil(t, err)
			testutil.Equal(t, 200, resp.StatusCode())
			testutil.Equal(t, "bar", string(resp.Body()))
		},
	))
}

func TestContext_Bind(t *testing.T) {
	t.Run("query", With(
		func(t *testing.T, s *http.Server) {
			s.Get("/query", func(c *http.Context) error {
				type V struct {
					Q1 string `query:"q1"`
					Q2 int    `query:"q2"`
				}
				v := &V{}
				e := c.BindQuery(v)
				testutil.LogNoErr(t, e, v)
				testutil.Equal(t, "query", v.Q1)
				testutil.Equal(t, 123, v.Q2)
				return c.TEXT(200, "ok")
			})
		},
		func(t *testing.T, c *resty.Client) {
			resp, err := c.R().SetQueryParams(gh.MS{"q1": "query", "q2": "123"}).Get("/query")
			testutil.Equal(t, 200, resp.StatusCode(), err)
		},
	))

	t.Run("header", With(
		func(t *testing.T, s *http.Server) {
			s.Get("/header", func(c *http.Context) error {
				type V struct {
					H1 string `header:"h1"`
					H2 int    `header:"h2"`
				}
				v := &V{}
				e := c.BindHeader(v)
				testutil.LogNoErr(t, e, v)
				testutil.Equal(t, "header", v.H1)
				testutil.Equal(t, 123, v.H2)
				return c.TEXT(200, "ok")
			})
		},
		func(t *testing.T, c *resty.Client) {
			resp, err := c.R().SetHeaders(gh.MS{"h1": "header", "h2": "123"}).Get("/header")
			testutil.Equal(t, 200, resp.StatusCode(), err)
		},
	))

	t.Run("form", With(
		func(t *testing.T, s *http.Server) {
			s.Post("/form", func(c *http.Context) error {
				type V struct {
					Q1 string `form:"q1"`
					Q2 int    `form:"q2"`
					F1 string `form:"f1"`
					F2 int    `form:"f2"`
				}
				v1 := &V{}
				e1 := c.BindForm(v1)
				testutil.LogNoErr(t, e1, v1)
				testutil.Equal(t, "query", v1.Q1)
				testutil.Equal(t, 123, v1.Q2)
				testutil.Equal(t, "form", v1.F1)
				testutil.Equal(t, 456, v1.F2)
				v2 := &V{}
				e2 := c.BindPostForm(v2)
				testutil.LogNoErr(t, e2, v2)
				testutil.Equal(t, "", v2.Q1)
				testutil.Equal(t, 0, v2.Q2)
				testutil.Equal(t, "form", v1.F1)
				testutil.Equal(t, 456, v1.F2)
				return c.TEXT(200, "ok")
			})
		},
		func(t *testing.T, c *resty.Client) {
			resp, err := c.R().SetQueryParams(gh.MS{"q1": "query", "q2": "123"}).SetFormData(gh.MS{"f1": "form", "f2": "456"}).Post("/form")
			testutil.Equal(t, 200, resp.StatusCode(), err)
		},
	))

	t.Run("multi_form", With(
		func(t *testing.T, s *http.Server) {
			s.Post("/multi_form", func(c *http.Context) error {
				type V struct {
					F1         string                `form:"f1"`
					F2         int                   `form:"f2"`
					FileHeader *multipart.FileHeader `form:"file"`
					File       multipart.File        `form:"file"`
				}
				v := &V{}
				e1 := c.BindMultiForm(v)
				testutil.LogNoErr(t, e1, v)
				bs, e2 := io.ReadAll(v.File)
				testutil.LogNoErr(t, e2, string(bs))
				testutil.Equal(t, "data", string(bs))
				return c.TEXT(200, "ok")
			})
		},
		func(t *testing.T, c *resty.Client) {
			resp, err := c.R().SetMultipartFormData(gh.MS{"f1": "form1", "f2": "123"}).SetFileReader("file", "test.txt", strings.NewReader("data")).Post("/multi_form")
			testutil.Equal(t, 200, resp.StatusCode(), err)
		},
	))
}

func TestContext_Render(t *testing.T) {
	t.Run("json", With(
		func(t *testing.T, s *http.Server) {
			s.Post("/json", func(c *http.Context) error {
				return c.JSON(200, gh.M{"foo": "bar", "num": 123})
			})
		},
		func(t *testing.T, c *resty.Client) {
			resp, err := c.R().Post("/json")
			testutil.Nil(t, err)
			testutil.Equal(t, 200, resp.StatusCode())
			testutil.Equal(t, `{"foo":"bar","num":123}`, string(resp.Body()))
		},
	))
}

func With(sfn func(t *testing.T, s *http.Server), cfn func(t *testing.T, c *resty.Client)) func(*testing.T) {
	return func(t *testing.T) {
		s := http.NewServer()
		ln := osutil.Must1(net.Listen("tcp", ":54444"))
		go func() {
			if err := s.Start(ln); err != nil && !errors.Is(err, stdhttp.ErrServerClosed) {
				osutil.PanicErr(err)
			}
		}()
		sfn(t, s)
		c := resty.New().SetBaseURL(fmt.Sprintf("http://localhost:%d", ln.Addr().(*net.TCPAddr).Port))
		cfn(t, c)
		s.Stop(time.Second)
	}
}
