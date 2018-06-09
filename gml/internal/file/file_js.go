// +build js

package file

import (
	"path/filepath"
	"strings"

	"github.com/gopherjs/gopherjs/js"
)

var (
	ProgramDirectory string = calculateProgramDir()
)

func calculateProgramDir() string {
	// Setup program dir
	location := js.Global.Get("location")
	result := location.Get("href").String()
	result = filepath.Dir(result)
	result = strings.TrimPrefix(result, "file:/")
	if strings.HasPrefix(result, "http:/") {
		result = strings.TrimPrefix(result, "http:/")
		result = "http://" + result
	}
	if strings.HasPrefix(result, "https:/") {
		result = strings.TrimPrefix(result, "https:/")
		result = "https://" + result
	}
	return result
}
