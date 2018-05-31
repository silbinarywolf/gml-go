// +build js

package gml

import (
	"path/filepath"
	"strings"

	"github.com/gopherjs/gopherjs/js"
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
	// Setup working directory
	location := js.Global.Get("location")
	result := location.Get("href").String()
	result = filepath.Dir(result)
	result = strings.TrimPrefix(result, "file:/")
	// todo(Jake): 2018-05-31
	//
	// Detect https and account for it
	//
	result = strings.TrimPrefix(result, "http:/")
	result = "http://" + result
	workingDirectory = result
}
