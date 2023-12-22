package http

import (
	"net/http"
)

// header -> code -> body
type responseWriter struct {
	http.ResponseWriter
	status int
	size   int
}

func newResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{ResponseWriter: w, status: http.StatusOK, size: -1}
}

func (w *responseWriter) Status() int {
	return w.status
}

func (w *responseWriter) Size() int {
	return w.size
}

func (w *responseWriter) Written() bool {
	return w.Size() >= 0
}

var _ http.ResponseWriter = (*responseWriter)(nil)

func (w *responseWriter) Header() http.Header {
	return w.ResponseWriter.Header()
}

func (w *responseWriter) WriteHeader(status int) {
	w.status = status
}

func (w *responseWriter) WriteHeaderNow() {
	if !w.Written() {
		w.size = 0
		w.ResponseWriter.WriteHeader(w.status)
	}
}

func (w *responseWriter) Write(bs []byte) (int, error) {
	w.WriteHeaderNow()
	n, err := w.ResponseWriter.Write(bs)
	w.size += n
	return n, err
}
