// +build !js,!windows

package timeprec

import (
	"time"
)

func Now() int64 {
	return time.Now().UnixNano()
}
