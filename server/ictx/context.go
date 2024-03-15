package ictx

import (
	"context"
)

type Context struct {
	context.Context
}

var _ context.Context = (*Context)(nil)
