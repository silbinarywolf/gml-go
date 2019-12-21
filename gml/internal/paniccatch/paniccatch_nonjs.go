// +build !js

package paniccatch

import (
	"log"
	"os"
)

func init() {
	f, err := os.OpenFile("error.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v\n", err)
	}
	redirectStderr(f)
}
