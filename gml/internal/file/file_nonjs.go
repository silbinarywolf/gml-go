// +build darwin freebsd linux windows
// +build !js
// +build !android
// +build !ios

package file

import (
	"io"
	"os"
	"path/filepath"
)

var (
	ProgramDirectory string = calculateProgramDir()
)

// ReadSeekCloser is io.ReadSeeker and io.Closer.
type readSeekCloser interface {
	io.ReadSeeker
	io.Closer
}

func OpenFile(path string) (readSeekCloser, error) {
	return os.Open(filepath.FromSlash(path))
}

func calculateProgramDir() string {
	exePath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	result, err := filepath.Abs(filepath.Dir(exePath))
	if err != nil {
		panic(err)
	}
	return result
}
