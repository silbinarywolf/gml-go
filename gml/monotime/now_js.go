// +build js

package monotime

import (
	"time"

	"syscall/js"
)

func now() int64 {
	// time.Now() is not reliable until GopherJS supports performance.now().
	return int64(js.Global().Get("performance").Call("now").Float() * float64(time.Millisecond))
}
