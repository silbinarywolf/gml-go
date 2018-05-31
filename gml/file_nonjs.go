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
	workingDirectory string
)

func currentDirectory() string {
	return workingDirectory
}

func WorkingDirectory() string {
	return workingDirectory
}

func init() {
	var err error
	workingDirectory, err = filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
}
