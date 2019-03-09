// +build !debug

package assert

// DebugAssert will panic if the first argument evaluates true and if the game was built with a "debug" tag
func DebugAssert(expr bool, message string) {
}
