package selfupdate

import (
	"context"
	"errors"
	"os/exec"
	"strings"
	"time"

	"github.com/Masterminds/semver/v3"

	"github.com/starudream/go-lib/core/v2/config/version"
	"github.com/starudream/go-lib/core/v2/slog"
	"github.com/starudream/go-lib/core/v2/utils/osutil"
)

type Version = semver.Version

func CurrentVersion() *Version {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	go func() {
		<-ctx.Done()
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			slog.Warn("get current version timeout")
		}
	}()

	output, err := exec.CommandContext(ctx, osutil.ExeFull(), "-v").Output()
	if err == nil {
		gv := version.Parse(string(output)).GitVersion
		if gv != "" {
			idx := strings.LastIndex(gv, "/")
			if idx >= 0 {
				gv = gv[idx+1:]
			}
			sv, _ := semver.NewVersion(gv)
			return sv
		}
	}

	return nil
}
