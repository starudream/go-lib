package task

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"

	"github.com/starudream/go-lib/core/v2/utils/osutil"
)

type Pool struct {
	ctx    context.Context
	cancel context.CancelFunc

	tasks []func(context.Context) error
}

func NewPool(ctx context.Context) *Pool {
	p := &Pool{}
	p.ctx, p.cancel = context.WithCancel(ctx)
	return p
}

func (p *Pool) Add(fn func(ctx context.Context) error) *Pool {
	p.tasks = append(p.tasks, fn)
	return p
}

func (p *Pool) Run(n ...int) error {
	num := runtime.NumCPU()
	if len(n) > 0 && n[0] > 0 {
		num = n[0]
	}

	wg := sync.WaitGroup{}
	wg.Add(len(p.tasks))

	var (
		errVal  atomic.Value
		errOnce sync.Once

		taskCh = make(chan func(context.Context) error)
	)

	for i := 0; i < num; i++ {
		go func() {
			for {
				task, ok := <-taskCh
				if !ok {
					return
				}

				func() {
					defer wg.Done()

					if task == nil || errVal.Load() != nil {
						return
					}

					defer func() {
						if r := recover(); r != nil {
							errOnce.Do(func() {
								errVal.Store(fmt.Errorf("[%s] %v", osutil.CallerString(5), r))
								p.cancel()
							})
						}
					}()

					err := task(p.ctx)
					if err != nil {
						errOnce.Do(func() {
							errVal.Store(err)
							p.cancel()
						})
					}
				}()
			}
		}()
	}

	for i := range p.tasks {
		taskCh <- p.tasks[i]
	}

	wg.Wait()

	close(taskCh)

	//goland:noinspection GoTypeAssertionOnErrors
	err, _ := errVal.Load().(error)
	return err
}
