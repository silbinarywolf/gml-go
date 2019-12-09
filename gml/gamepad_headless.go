// +build headless

package gml

type gamepadButton int

func GamepadCheckPressed(id int, button gamepadButton) bool {
	return false
}

func GamepadCheck(id int, button gamepadButton) bool {
	return false
}
