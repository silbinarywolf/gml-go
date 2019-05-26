// +build js

package file

import (
	"path/filepath"
	"strings"

	"syscall/js"
)

// computeProgramDirectory returns the directory or url that the executable is running from
func computeProgramDirectory() string {
	location := js.Global().Get("location")
	url := location.Get("href").String()
	url = filepath.Dir(url)
	url = strings.TrimPrefix(url, "file:/")
	if strings.HasPrefix(url, "http:/") {
		url = strings.TrimPrefix(url, "http:/")
		url = "http://" + url
	}
	if strings.HasPrefix(url, "https:/") {
		url = strings.TrimPrefix(url, "https:/")
		url = "https://" + url
	}
	return url
}
