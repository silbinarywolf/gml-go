// +build darwin freebsd linux windows
// +build !js
// +build !android
// +build !ios

package gml

import (
	"os"
	"path/filepath"
)

var (
	programDirectory string = calculateProgramDir()
)

func ProgramDirectory() string {
	return programDirectory
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
