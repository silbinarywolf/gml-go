// +build !js,!windows

package monotime

import (
	"time"
)

func now() int64 {
	return time.Now().UnixNano()
}
