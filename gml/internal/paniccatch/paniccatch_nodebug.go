// +build !debug

package paniccatch

func init() {
	maybeRedirectPanicToLog()
}
