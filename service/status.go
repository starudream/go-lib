package service

import (
	"errors"

	"github.com/kardianos/service"
)

const (
	statusRunning = service.StatusRunning
	statusStopped = service.StatusStopped
)

var errNotInstalled = service.ErrNotInstalled

func getStatus(svc Service) (running, stopped, installed bool) {
	sts, err := svc.Status()
	return err == nil && sts == statusRunning, err == nil && sts == statusStopped, err == nil || !errors.Is(err, errNotInstalled)
}
