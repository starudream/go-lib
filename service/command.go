package service

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/starudream/go-lib/cobra/v2"
	"github.com/starudream/go-lib/core/v2/gh"
	"github.com/starudream/go-lib/core/v2/slog"
	"github.com/starudream/go-lib/core/v2/utils/strutil"
)

func AddCommand(cmd *cobra.Command, svc Service) {
	serviceCmd := cobra.NewCommand(func(c *cobra.Command) {
		c.Use = "service"
		c.Short = "Manage service"
		c.PersistentFlags().Bool("user", false, "run service as user")
	})

	serviceCmd.AddCommand(cobra.NewCommand(func(c *cobra.Command) {
		c.Use = "status"
		c.Short = "Show service status"
		c.RunE = func(cmd *cobra.Command, args []string) error {
			sts, err := svc.Status()
			msg := gh.Ternary(sts == statusRunning, "service is running", "service is stopped")
			if err == nil {
				msg += fmt.Sprintf(", %s managed by %s at %q", svc.String(), svc.Platform(), getConfigPath(svc))
			} else if errors.Is(err, errNotInstalled) {
				msg, err = "service is not installed", nil
			}
			return mPrintln(err, "service status unknown", msg)
		}
	}))

	for _, action := range actions {
		serviceCmd.AddCommand(cobra.NewCommand(func(c *cobra.Command) {
			c.Use = action
			c.Short = strutil.ToSnakeCase(action) + " service"
			c.RunE = func(cmd *cobra.Command, args []string) error { return control(svc, cmd.Use) }
		}))
	}

	cmd.AddCommand(serviceCmd)
}

// don't change the order
var actions = []string{"install", "uninstall", "start", "stop", "restart", "reinstall"}

func control(svc Service, action string) (err error) {
	fe, fs := fmt.Sprintf("service %s error", action), fmt.Sprintf("service %s success", action)
	defer func() { err = mPrintln(err, fe, fs) }()
	running, stopped, installed := getStatus(svc)
	if action == actions[0] && installed {
		fs = fmt.Sprintf("service is already installed at %q", getConfigPath(svc))
		return nil
	} else if action != actions[0] && action != actions[5] && !installed {
		fs = "service is not installed"
		return nil
	}
	switch action {
	case actions[0]:
		return svc.Install()
	case actions[1]:
		return svc.Uninstall()
	case actions[2]:
		if running {
			fs = "service is already running"
			return nil
		}
		return svc.Start()
	case actions[3]:
		if stopped {
			fs = "service is already stopped"
			return nil
		}
		return svc.Stop()
	case actions[4]:
		if running {
			err = svc.Stop()
			if err != nil {
				return err
			}
		}
		time.Sleep(100 * time.Millisecond)
		return svc.Start()
	case actions[5]:
		if installed {
			err = svc.Uninstall()
			if err != nil {
				return err
			}
		}
		time.Sleep(100 * time.Millisecond)
		return svc.Install()
	default:
		panic("unreachable")
	}
}

func mPrintln(err error, fe, fs string) error {
	if err != nil {
		esg := fe + ": " + strings.ReplaceAll(strings.TrimSuffix(err.Error(), "\n"), "\n", ", ")
		if strings.Contains(esg, "permission denied") {
			esg += ". try to run with sudo or use --user flag"
		}
		err = errors.New(esg)
		log(slog.LevelError, esg)
	} else {
		log(slog.LevelInfo, fs)
		fmt.Println(fs)
	}
	return err
}
