package selfupdate

import (
	"bytes"
	"crypto"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/starudream/go-lib/core/v2/utils/optionutil"
)

type Options struct {
	// TargetPath defines the path to the file to update.
	// The empty string means 'the executable file of the running program'.
	TargetPath string

	// Create TargetPath replacement with this file mode. If zero, defaults to 0755.
	TargetMode os.FileMode

	// Checksum of the new binary to verify against. If nil, no checksum or signature verification is done.
	Checksum []byte

	// Use this hash function to generate the checksum. If not set, crypto.SHA256 is used.
	Hash crypto.Hash
}

func newOptions(options ...Option) *Options {
	return optionutil.Build(&Options{
		TargetMode: 0755,
		Hash:       crypto.SHA256,
	}, options...)
}

type Option = optionutil.I[Options]

func WithTargetPath(targetPath string) Option {
	return optionutil.New(func(t *Options) {
		t.TargetPath = targetPath
	})
}

func WithTargetMode(targetMode os.FileMode) Option {
	return optionutil.New(func(t *Options) {
		t.TargetMode = targetMode
	})
}

func WithChecksum(hash crypto.Hash, checksum []byte) Option {
	return optionutil.New(func(t *Options) {
		t.Hash = hash
		t.Checksum = checksum
	})
}

// CheckPermissions determines whether the process has the correct permissions to
// perform the requested update. If the update can proceed, it returns nil, otherwise
// it returns the error that would occur if an update were attempted.
func (o *Options) CheckPermissions() error {
	// get the directory the file exists in
	path, err := o.getPath()
	if err != nil {
		return err
	}

	fileDir := filepath.Dir(path)
	fileName := filepath.Base(path)

	// attempt to open a file in the file's directory
	newPath := filepath.Join(fileDir, fmt.Sprintf(".%s.new", fileName))

	newFile, err := os.OpenFile(newPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, o.TargetMode)
	if err != nil {
		return err
	}
	newFile.Close()

	_ = os.Remove(newPath)
	return nil
}

func (o *Options) getPath() (string, error) {
	if o.TargetPath != "" {
		return o.TargetPath, nil
	}
	return os.Executable()
}

func (o *Options) verifyChecksum(updated []byte) error {
	checksum, err := checksumFor(o.Hash, updated)
	if err != nil {
		return err
	}
	if !bytes.Equal(o.Checksum, checksum) {
		return fmt.Errorf("updated file has wrong checksum. expected: %x, actual: %x", o.Checksum, checksum)
	}
	return nil
}

func checksumFor(h crypto.Hash, payload []byte) ([]byte, error) {
	if !h.Available() {
		return nil, errors.New("requested hash function not available")
	}
	hash := h.New()
	hash.Write(payload) // guaranteed not to error
	return hash.Sum([]byte{}), nil
}
