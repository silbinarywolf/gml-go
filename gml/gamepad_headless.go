// +build headless

package gml

func GamepadCheckPressed(id int, button GamepadButton) bool {
	return false
}

func GamepadCheck(id int, button GamepadButton) bool {
	return false
}

func GamepadAxisValue(id int, axis GamepadAxis) float64 {
	return 0
}

func gamepadUpdate() {
}
