// +build js

package gml

import (
	"path/filepath"
	"strings"

	"github.com/gopherjs/gopherjs/js"
)

var location = js.Global.Get("location")

func currentDirectory() string {
	result := location.Get("href").String()
	result = filepath.Dir(result)
	result = strings.TrimPrefix(result, "file:/")
	result = strings.TrimPrefix(result, "http:/")
	result = "http://" + result
	return result
}
