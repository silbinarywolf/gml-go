package input

import (
	"github.com/silbinarywolf/gml-go/gml"
)

var (
	mockJumpPressed bool
)

func JumpPressed() bool {
	return mockJumpPressed ||
		gml.MouseCheckPressed(gml.MbLeft) ||
		gml.KeyboardCheckPressed(gml.VkSpace) ||
		gml.KeyboardCheckPressed(gml.VkEnter)
}

// TestResetJumpPressed is used by the tests to reset the state of whether jump is being
// pressed or not
func TestResetJumpPressed() {
	mockJumpPressed = false
}

// SetJumpPressed is used by testing tools to mock pressing of the jump button
func TestSetJumpPressed(pressed bool) {
	mockJumpPressed = pressed
}
