package middlewares

import (
	"time"

	"github.com/go-chi/chi/v5/middleware"

	"github.com/starudream/go-lib/core/v2/codec/json"
	"github.com/starudream/go-lib/core/v2/slog"
	"github.com/starudream/go-lib/core/v2/utils/poolutil"
	"github.com/starudream/go-lib/server/v2/http"
	"github.com/starudream/go-lib/server/v2/iconst"
	"github.com/starudream/go-lib/server/v2/jwt"
)

var loggerBuf = poolutil.NewBytesBuffer(1024)

func Logger() http.Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerCtx(func(c *http.Context) error {
			attrs := slog.GetAttrs(c)

			attrs = append(attrs,
				slog.String("method", c.Req.Method),
				slog.String("path", c.Req.URL.Path),
			)

			claims, _ := jwt.FromContext(c)
			if claims != nil {
				attrs = append(attrs,
					slog.String("jwt-issuer", claims.ISS()),
					slog.String("jwt-subject", claims.SUB()),
					slog.String("jwt-audience", claims.AUD()),
				)
			}

			loggerReq(c, attrs)

			rw := middleware.NewWrapResponseWriter(c.Res, c.Req.ProtoMajor)
			c.Res = rw
			buf := loggerBuf.Get()
			defer loggerBuf.Put(buf)
			rw.Tee(buf)

			defer func(start time.Time) {
				attrs = append(attrs, slog.Duration("took", time.Since(start)))
				loggerResp(c, attrs, buf.Bytes())
			}(time.Now())

			c.Req = c.Req.WithContext(slog.WithAttrs(c.Context(), attrs...))

			next.ServeHTTP(c.Res, c.Req)

			return nil
		})
	}
}

func loggerReq(c *http.Context, attrs []slog.Attr) {
	msg := ""
	typ := slog.String("content-type", filterFlags(c.Req.Header.Get(iconst.HeaderContentType)))
	if body := c.MustGetRawBody(); len(body) == 0 {
		msg = "req: <empty>"
	} else if req, err := json.Compact(body); err == nil {
		msg = "req: " + string(req)
	} else {
		msg = "req: <ignore>"
	}
	slog.Info(msg, attrs, typ)
}

func loggerResp(c *http.Context, attrs []slog.Attr, body []byte) {
	msg := ""
	typ := slog.String("content-type", filterFlags(c.Res.Header().Get(iconst.HeaderContentType)))
	if len(body) == 0 {
		msg = "resp: <empty>"
	} else if resp, err := json.Compact(body); err == nil {
		msg = "resp: " + string(resp)
	} else {
		msg = "resp: <ignore>"
	}
	slog.Info(msg, attrs, typ)
}
