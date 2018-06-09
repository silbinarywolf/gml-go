// +build darwin freebsd linux windows
// +build !js
// +build !android
// +build !ios

package file

import (
	"os"
	"path/filepath"
)

var (
	ProgramDirectory string = calculateProgramDir()
)

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
