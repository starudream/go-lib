package service

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/kardianos/service"
	"github.com/spf13/pflag"

	"github.com/starudream/go-lib/core/v2/config/global"
	"github.com/starudream/go-lib/core/v2/utils/maputil"
	"github.com/starudream/go-lib/core/v2/utils/osutil"
)

type Service = service.Service

var _services maputil.SyncMap[string, Service]

func New(name string, run func(context.Context), options ...Option) Service {
	svc, exists := _services.Load(name)
	if !exists {
		svc = osutil.Must1(newService(name, run, options...))
		_services.Store(name, svc)
	}
	return svc
}

func newService(name string, run func(context.Context), options ...Option) (Service, error) {
	if name == "" {
		return nil, fmt.Errorf("service name is empty")
	}
	if run == nil {
		return nil, fmt.Errorf("service run is nil")
	}

	opts := newOptions(options...)

	svc, err := service.New(&program{Run: run}, &service.Config{
		Name:             name,
		DisplayName:      opts.displayName,
		Description:      opts.description,
		Arguments:        opts.arguments,
		WorkingDirectory: osutil.WorkDir(),
		Option:           defaultServiceOption(),
		EnvVars:          opts.envVars,
	})
	if err != nil {
		return nil, err
	}

	return svc, nil
}

func defaultServiceOption() map[string]any {
	return map[string]any{
		"UserService": func() (user bool) {
			fs := pflag.NewFlagSet("", pflag.ContinueOnError)
			fs.SetOutput(io.Discard)
			fs.BoolVar(&user, "user", false, "")
			_ = fs.Parse(os.Args[1:])
			return
		}(),
		"LogDirectory": filepath.Dir(global.C().LogFileFilename),
	}
}
