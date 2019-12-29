// +build headless

package gml

func GamepadCheckPressed(id int, button GamepadButton) bool {
	return false
}

func GamepadCheck(id int, button GamepadButton) bool {
	return false
}

func GamepadButtonCount(id int) GamepadButton {
	return 0
}

func GamepadAxisCount(id int) GamepadAxis {
	return 0
}

func GamepadAxisValue(id int, axis GamepadAxis) float64 {
	return 0
}

func GamepadGetDescription(id int) string {
	return ""
}

func gamepadUpdate() {
}
