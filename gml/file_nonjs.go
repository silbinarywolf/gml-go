// +build darwin freebsd linux windows
// +build !js
// +build !android
// +build !ios

package gml

import (
	"os"
	"path/filepath"
)

func currentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	return dir
}
