// +build !js,!windows

package timegml

import (
	"time"
)

func Now() int64 {
	return time.Now().UnixNano()
}
