// +build js

package paniccatch

func maybeRedirectPanicToLog() {
	// do not redirect logging or do anything special for
	// panic crashes in the browser
}
