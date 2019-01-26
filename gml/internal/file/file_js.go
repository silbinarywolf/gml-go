// +build js

package file

import (
	"bytes"
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/gopherjs/gopherwasm/js"
)

func OpenFile(path string) (readSeekCloser, error) {
	// NOTE(Jake): 2019-01-26
	// Copied this from ebitenutil.OpenFile().
	// We want to avoid importing ebitenutil as it contains debug assets
	// that I don't want such as text/font.
	var err error
	var content js.Value
	ch := make(chan struct{})
	req := js.Global().Get("XMLHttpRequest").New()
	req.Call("open", "GET", path, true)
	req.Set("responseType", "arraybuffer")
	loadCallback := js.NewCallback(func([]js.Value) {
		defer close(ch)
		status := req.Get("status").Int()
		if 200 <= status && status < 400 {
			content = req.Get("response")
			return
		}
		err = errors.New(fmt.Sprintf("http error: %d", status))
	})
	defer loadCallback.Release()
	req.Call("addEventListener", "load", loadCallback)
	errorCallback := js.NewCallback(func([]js.Value) {
		defer close(ch)
		err = errors.New(fmt.Sprintf("XMLHttpRequest error: %s", req.Get("statusText").String()))
	})
	req.Call("addEventListener", "error", errorCallback)
	defer errorCallback.Release()
	req.Call("send")
	<-ch
	if err != nil {
		return nil, err
	}

	uint8contentWrapper := js.Global().Get("Uint8Array").New(content)
	data := make([]byte, uint8contentWrapper.Get("byteLength").Int())
	arr := js.TypedArrayOf(data)
	arr.Call("set", uint8contentWrapper)
	arr.Release()
	f := &file{bytes.NewReader(data)}
	return f, nil
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
