// +build darwin freebsd linux windows
// +build !js
// +build !android
// +build !ios

package gml

import (
	"path/filepath"
	"runtime"
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
	_, filename, _, _ := runtime.Caller(2)

	var err error
	programDirectory, err := filepath.Abs(filepath.Dir(filename))
	if err != nil {
		panic(err)
	}
	return programDirectory
}
