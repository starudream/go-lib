package task

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	p := NewPool(context.Background())
	p.Add(func(ctx context.Context) error {
		time.Sleep(time.Second)
		panic("1")
	})
	p.Add(func(ctx context.Context) error {
		time.Sleep(time.Second)
		fmt.Println("hello")
		return nil
	})
	p.Add(func(ctx context.Context) error {
		return fmt.Errorf("empty")
	})
	t.Log(p.Run(1))
}
