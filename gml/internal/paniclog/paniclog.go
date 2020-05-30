package paniclog

// Init will redirect crash logging to "logs/*.log" file if not running
// with debug tags.
func Init() {
	initPanicCatch()
}
