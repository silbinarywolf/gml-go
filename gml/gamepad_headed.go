// +build !headless

package gml

import "github.com/hajimehoshi/ebiten"

var gamepadToEbiten = []ebiten.GamepadButton{
	GpFace1:      ebiten.GamepadButton0,
	GpFace2:      ebiten.GamepadButton1,
	GpFace3:      ebiten.GamepadButton2,
	GpFace4:      ebiten.GamepadButton3,
	GpShoulderLB: ebiten.GamepadButton4,
	GpShoulderRB: ebiten.GamepadButton5,
	GpShoulderLT: 0, // Uses Axis 4
	GpShoulderRT: 0, // Uses Axis 5
	GpSelect:     ebiten.GamepadButton6,
	GpStart:      ebiten.GamepadButton7,
	GpStickLeft:  ebiten.GamepadButton8,
	GpStickRight: ebiten.GamepadButton9,
	GpPadUp:      ebiten.GamepadButton10,
	GpPadRight:   ebiten.GamepadButton11,
	GpPadDown:    ebiten.GamepadButton12,
	GpPadLeft:    ebiten.GamepadButton13,
	/*GamepadButton0:  ebiten.GamepadButton0,
	GamepadButton1:  ebiten.GamepadButton1,
	GamepadButton2:  ebiten.GamepadButton2,
	GamepadButton3:  ebiten.GamepadButton3,
	GamepadButton4:  ebiten.GamepadButton4,
	GamepadButton5:  ebiten.GamepadButton5,
	GamepadButton6:  ebiten.GamepadButton6,
	GamepadButton7:  ebiten.GamepadButton7,
	GamepadButton8:  ebiten.GamepadButton8,
	GamepadButton9:  ebiten.GamepadButton9,
	GamepadButton10: ebiten.GamepadButton10,
	GamepadButton11: ebiten.GamepadButton11,
	GamepadButton12: ebiten.GamepadButton12,
	GamepadButton13: ebiten.GamepadButton13,
	GamepadButton14: ebiten.GamepadButton14,
	GamepadButton15: ebiten.GamepadButton15,
	GamepadButton16: ebiten.GamepadButton16,
	GamepadButton17: ebiten.GamepadButton17,
	GamepadButton18: ebiten.GamepadButton18,
	GamepadButton19: ebiten.GamepadButton19,
	GamepadButton20: ebiten.GamepadButton20,
	GamepadButton21: ebiten.GamepadButton21,
	GamepadButton22: ebiten.GamepadButton22,
	GamepadButton23: ebiten.GamepadButton23,
	GamepadButton24: ebiten.GamepadButton24,
	GamepadButton25: ebiten.GamepadButton25,
	GamepadButton26: ebiten.GamepadButton26,
	GamepadButton27: ebiten.GamepadButton27,
	GamepadButton28: ebiten.GamepadButton28,
	GamepadButton29: ebiten.GamepadButton29,
	GamepadButton30: ebiten.GamepadButton30,
	GamepadButton31: ebiten.GamepadButton31,*/
}

var (
	pressingButtonLastFrame [maxGamepads][gpSize]bool
)

func GamepadCheck(id int, button gamepadButton) bool {
	switch button {
	case GpShoulderLT:
		return ebiten.GamepadAxis(id, 4) > -0.5
	case GpShoulderRT:
		return ebiten.GamepadAxis(id, 5) > -0.5
	}
	return ebiten.IsGamepadButtonPressed(id, gamepadToEbiten[button])
}

func GamepadCheckPressed(id int, button gamepadButton) bool {
	isHeld := GamepadCheck(id, button)
	if !isHeld {
		pressingButtonLastFrame[id][button] = false
	}
	if pressingButtonLastFrame[id][button] {
		return false
	}
	if isHeld {
		pressingButtonLastFrame[id][button] = true
	}
	return isHeld
}
