// +build js

package file

import (
	"path/filepath"
	"strings"

	"github.com/gopherjs/gopherwasm/js"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

func OpenFile(path string) (readSeekCloser, error) {
	return ebitenutil.OpenFile(path)
}

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
