package ioutil

import (
	"context"
	"io"
)

type writer struct {
	ctx context.Context
	w   io.Writer
}

type copier struct {
	writer
}

func NewWriter(ctx context.Context, w io.Writer) io.Writer {
	if w, ok := w.(*copier); ok && ctx == w.ctx {
		return w
	}
	return &copier{writer{ctx: ctx, w: w}}
}

func (w *writer) Write(p []byte) (n int, err error) {
	select {
	case <-w.ctx.Done():
		return 0, w.ctx.Err()
	default:
		return w.w.Write(p)
	}
}

type reader struct {
	ctx context.Context
	r   io.Reader
}

func NewReader(ctx context.Context, r io.Reader) io.Reader {
	if r, ok := r.(*reader); ok && ctx == r.ctx {
		return r
	}
	return &reader{ctx: ctx, r: r}
}

func (r *reader) Read(p []byte) (n int, err error) {
	select {
	case <-r.ctx.Done():
		return 0, r.ctx.Err()
	default:
		return r.r.Read(p)
	}
}

func (w *copier) ReadFrom(r io.Reader) (n int64, err error) {
	if _, ok := w.w.(io.ReaderFrom); ok {
		return io.Copy(w.writer.w, &reader{ctx: w.ctx, r: r})
	}
	select {
	case <-w.ctx.Done():
		return 0, w.ctx.Err()
	default:
		return io.Copy(&w.writer, r)
	}
}

func NewCloser(ctx context.Context, c io.Closer) io.Closer {
	return &closer{ctx: ctx, c: c}
}

type closer struct {
	ctx context.Context
	c   io.Closer
}

func (c *closer) Close() error {
	select {
	case <-c.ctx.Done():
		return c.ctx.Err()
	default:
		return c.c.Close()
	}
}
