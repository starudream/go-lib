package ntfy

import (
	"context"
)

type Interface interface {
	Name() string
	Notify(ctx context.Context, text string) error
}
