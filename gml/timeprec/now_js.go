// +build js

package timeprec

import (
	"time"

	"github.com/gopherjs/gopherwasm/js"
)

func Now() int64 {
	// time.Now() is not reliable until GopherJS supports performance.now().
	return int64(js.Global().Get("performance").Call("now").Float() * float64(time.Millisecond))
}
