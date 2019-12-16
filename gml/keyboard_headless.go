// +build headless

package gml

func KeyboardCheck(key Key) bool {
	return false
}

func KeyboardCheckPressed(key Key) bool {
	return false
}

func keyboardUpdate() {
	// no-op needed here yet
}
