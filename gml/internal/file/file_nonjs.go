// +build darwin freebsd linux windows
// +build !js
// +build !android
// +build !ios

package file

import (
	"os"
	"path/filepath"
)

func OpenFile(path string) (readSeekCloser, error) {
	return os.Open(filepath.FromSlash(path))
}

// computeProgramDirectory returns the directory or url that the executable is running from
func computeProgramDirectory() string {
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
