package service

import (
	"context"

	"github.com/kardianos/service"
)

type program struct {
	ctx    context.Context
	cancel context.CancelFunc

	Run func(context.Context)
}

var _ service.Interface = (*program)(nil)

func (p *program) Start(Service) error {
	p.ctx, p.cancel = context.WithCancel(context.Background())
	go p.Run(p.ctx)
	return nil
}

func (p *program) Stop(Service) error {
	p.cancel()
	return nil
}
