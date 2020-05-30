// +build !js,!debug

package paniclog

import (
	"log"
	"os"
)

func initPanicCatch() {
	// NOTE(Jae): 2019-12-21
	// We dont utilize file.ProgramDirectory here as it won't
	// be computed yet, so we just log to the relative "logs"
	// directory for now.
	logDir := "logs"
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		if err := os.Mkdir(logDir, 0700); err != nil {
			panic("Failed to create logs folder")
		}
	}
	logFile := logDir + "/error.log"
	if err := os.Remove(logFile); err != nil {
		// todo(Jae): 2019-12-21
		// - store current datetime at app start
		// - write "error_TIMESTAMP.log" file
		// - get list of error files, keep only the current 3-10 runs? (make configurable)
	}
	f, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v\n", err)
	}
	redirectStderr(f)
}
