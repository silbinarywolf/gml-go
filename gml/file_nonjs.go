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
	// NOTE(Jake): 2018-06-03
	//
	// Using runtime.Caller(2) to get program directory.
	// We do this so `go test` just works OOTB.
	//
	// Source: https://stackoverflow.com/questions/18537257/how-to-get-the-directory-of-the-currently-running-file#comment59595931_18537792
	//
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
