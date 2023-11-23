package selfupdate

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Apply performs an update of the current executable (or opts.TargetFile, if set) with the contents of the given io.Reader.
//
// Apply performs the following actions to ensure a safe cross-platform update:
//
// 1. If configured, applies the contents of the update io.Reader as a binary patch.
//
// 2. If configured, computes the checksum of the new executable and verifies it matches.
//
// 3. If configured, verifies the signature with a public key.
//
// 4. Creates a new file, /path/to/.target.new with the TargetMode with the contents of the updated file
//
// 5. Renames /path/to/target to /path/to/.target.old
//
// 6. Renames /path/to/.target.new to /path/to/target
//
// 7. If the final rename is successful, deletes /path/to/.target.old, returns no error. On Windows,
// the removal of /path/to/target.old always fails, so instead Apply hides the old file instead.
//
// 8. If the final rename fails, attempts to roll back by renaming /path/to/.target.old
// back to /path/to/target.
//
// If the rollback operation fails, the file system is left in an inconsistent state (between steps 5 and 6) where
// there is no new executable file and the old executable file could not be moved to its original location. In this
// case you should notify the user of the bad news and ask them to recover manually. Applications can determine whether
// the rollback failed by calling RollbackError, see the documentation on that function for additional detail.
func Apply(update io.Reader, options ...Option) error {
	opts := newOptions(options...)

	// get target path
	var err error
	opts.TargetPath, err = opts.getPath()
	if err != nil {
		return err
	}

	// no patch to apply, go on through
	newBytes, err := io.ReadAll(update)
	if err != nil {
		return err
	}

	// verify checksum if requested
	if opts.Checksum != nil {
		if err = opts.verifyChecksum(newBytes); err != nil {
			return err
		}
	}

	// get the directory the executable exists in
	updateDir := filepath.Dir(opts.TargetPath)
	filename := filepath.Base(opts.TargetPath)

	// Copy the contents of new binary to a new executable file
	newPath := filepath.Join(updateDir, fmt.Sprintf(".%s.new", filename))

	newFile, err := os.OpenFile(newPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, opts.TargetMode)
	if err != nil {
		return err
	}
	defer newFile.Close()

	_, err = io.Copy(newFile, bytes.NewReader(newBytes))
	if err != nil {
		return err
	}

	// if we don't call newFile.Close(), windows won't let us move the new executable
	// because the file will still be "in use"
	newFile.Close()

	// this is where we'll move the executable to so that we can swap in the updated replacement
	oldPath := filepath.Join(updateDir, fmt.Sprintf(".%s.old", filename))

	// delete any existing old exec file - this is necessary on Windows for two reasons:
	// 1. after a successful update, Windows can't remove the .old file because the process is still running
	// 2. windows rename operations fail if the destination file already exists
	_ = os.Remove(oldPath)

	// move the existing executable to a new file in the same directory
	err = os.Rename(opts.TargetPath, oldPath)
	if err != nil {
		return err
	}

	// move the new executable in to become the new program
	err = os.Rename(newPath, opts.TargetPath)

	if err != nil {
		// move unsuccessful
		//
		// The filesystem is now in a bad state. We have successfully
		// moved the existing binary to a new location, but we couldn't move the new
		// binary to take its place. That means there is no file where the current executable binary
		// used to be!
		// Try to rollback by restoring the old binary to its original path.
		e := os.Rename(oldPath, opts.TargetPath)
		if e != nil {
			return &rollbackErr{err, e}
		}

		return err
	}

	// move successful, remove the old binary if needed
	errRemove := os.Remove(oldPath)
	// windows has trouble with removing old binaries, so hide it instead
	if errRemove != nil {
		_ = hideFile(oldPath)
	}

	return nil
}
