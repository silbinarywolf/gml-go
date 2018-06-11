// +build headless

package gml

func KeyboardCheck(key int16) bool {
	return false
}

func KeyboardCheckPressed(key int16) bool {
	return false
}

func keyboardUpdate() {
	// no-op needed here yet
}
