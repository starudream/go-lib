package selfupdate

import (
	"bytes"
	"errors"

	"github.com/Masterminds/semver/v3"

	"github.com/starudream/go-lib/core/v2/slog"
)

type Source interface {
	Latest() (string, error)
	Binary(version string) ([]byte, error)
}

var (
	ErrInvalidLocalVersion  = errors.New("invalid local version")
	ErrInvalidRemoteVersion = errors.New("invalid remote version")
	ErrInvalidRemoteFile    = errors.New("invalid remote file")
)

func Update(source Source, confirm func() bool) error {
	local := CurrentVersion()
	if local == nil {
		return ErrInvalidLocalVersion
	}

	v, err := source.Latest()
	if err != nil {
		return err
	}
	remote, err := semver.NewVersion(v)
	if err != nil {
		return ErrInvalidRemoteVersion
	}

	if remote.Equal(local) {
		slog.Info("up to date", slog.String("remote", remote.String()))
		return nil
	} else if remote.LessThan(local) {
		slog.Info("no need to update", slog.String("local", local.String()), slog.String("remote", remote.String()))
		return nil
	}

	slog.Info("update available", slog.String("local", local.String()), slog.String("remote", remote.String()))

	if !confirm() {
		return nil
	}

	binary, err := source.Binary(v)
	if err != nil {
		return err
	}

	err = Apply(bytes.NewReader(binary))
	if err != nil {
		return err
	}

	current := CurrentVersion()
	if current == nil {
		return ErrInvalidRemoteFile
	}

	slog.Info("updated", slog.String("before", local.String()), slog.String("after", current.String()))

	return nil
}
