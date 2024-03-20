package ictx

import (
	"context"

	"google.golang.org/grpc/metadata"
)

type Context struct {
	context.Context
}

var _ context.Context = (*Context)(nil)

const srvCtxkey = "server-ctxkey"

func FromContext(ctx context.Context) *Context {
	if ctx == nil {
		ctx = context.Background()
	}
	if c, ok := ctx.(*Context); ok {
		return c
	}
	if c, ok := ctx.Value(srvCtxkey).(*Context); ok {
		return c
	}
	var kvs []string
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		for k, vs := range md {
			for i := len(vs) - 1; i >= 0; i-- {
				kvs = append([]string{k, vs[i]}, kvs...)
			}
		}
	}
	return &Context{Context: metadata.AppendToOutgoingContext(ctx, kvs...)}
}

func (c *Context) Set(kvs ...string) *Context {
	if len(kvs) == 0 {
		return c
	}
	if len(kvs)%2 != 0 {
		panic("kvs must be even")
	}
	md, _ := metadata.FromOutgoingContext(c.Context)
	md = md.Copy()
	for i := 0; i < len(kvs); i += 2 {
		md.Set(kvs[i], kvs[i+1])
	}
	c.Context = metadata.NewOutgoingContext(c.Context, md)
	return c
}

func (c *Context) Append(k string, vs ...string) *Context {
	if k == "" || len(vs) == 0 {
		return c
	}
	md, _ := metadata.FromOutgoingContext(c.Context)
	md = md.Copy()
	for i := 0; i < len(vs); i++ {
		md.Append(k, vs[i])
	}
	c.Context = metadata.NewOutgoingContext(c.Context, md)
	return c
}

func (c *Context) Get(ks ...string) string {
	md, _ := metadata.FromOutgoingContext(c.Context)
	for i := 0; i < len(ks); i++ {
		vs := md.Get(ks[i])
		for j := len(vs) - 1; j >= 0; j-- {
			if vs[j] != "" {
				return vs[j]
			}
		}
	}
	return ""
}
