// +build !headless

package gml

import (
	"github.com/hajimehoshi/ebiten"
)

type gamepadState int32

const (
	gamepadNotHeld gamepadState = 0 + iota
	gamepadPressed
	gamepadHeld
)

var gamepadToEbiten = []ebiten.GamepadButton{
	// Generic Gamepad
	GpButton0:  ebiten.GamepadButton0,
	GpButton1:  ebiten.GamepadButton1,
	GpButton2:  ebiten.GamepadButton2,
	GpButton3:  ebiten.GamepadButton3,
	GpButton4:  ebiten.GamepadButton4,
	GpButton5:  ebiten.GamepadButton5,
	GpButton6:  ebiten.GamepadButton6,
	GpButton7:  ebiten.GamepadButton7,
	GpButton8:  ebiten.GamepadButton8,
	GpButton9:  ebiten.GamepadButton9,
	GpButton10: ebiten.GamepadButton10,
	GpButton11: ebiten.GamepadButton11,
	GpButton12: ebiten.GamepadButton12,
	GpButton13: ebiten.GamepadButton13,
	GpButton14: ebiten.GamepadButton14,
	GpButton15: ebiten.GamepadButton15,
	GpButton16: ebiten.GamepadButton16,
	GpButton17: ebiten.GamepadButton17,
	GpButton18: ebiten.GamepadButton18,
	GpButton19: ebiten.GamepadButton19,
	GpButton20: ebiten.GamepadButton20,
	GpButton21: ebiten.GamepadButton21,
	GpButton22: ebiten.GamepadButton22,
	GpButton23: ebiten.GamepadButton23,
	GpButton24: ebiten.GamepadButton24,
	GpButton25: ebiten.GamepadButton25,
	GpButton26: ebiten.GamepadButton26,
	GpButton27: ebiten.GamepadButton27,
	GpButton28: ebiten.GamepadButton28,
	GpButton29: ebiten.GamepadButton29,
	GpButton30: ebiten.GamepadButton30,
	GpButton31: ebiten.GamepadButton31,
	// Special Key Handle for Xbox 360 / PS4 controllers
	GpShoulderLT: 0, // Uses Axis 4
	GpShoulderRT: 0, // Uses Axis 5
}

var (
	gamepadStateList [maxGamepads][GpSize]gamepadState
)

func GamepadButtonCount(id int) GamepadButton {
	return GamepadButton(ebiten.GamepadButtonNum(id))
}

//func GamepadGetDeviceCount() int {
//	return len(ebiten.GamepadIDs())
//}

func GamepadCheck(id int, button GamepadButton) bool {
	switch button {
	case GpShoulderLT:
		if ebiten.GamepadAxisNum(id) > 4 {
			return ebiten.GamepadAxis(id, 4) > -0.5
		}
		return false
	case GpShoulderRT:
		if ebiten.GamepadAxisNum(id) > 5 {
			return ebiten.GamepadAxis(id, 5) > -0.5
		}
		return false
	}
	return ebiten.IsGamepadButtonPressed(id, gamepadToEbiten[button])
}

func GamepadCheckPressed(id int, button GamepadButton) bool {
	return gamepadStateList[id][button] == gamepadPressed
}

func GamepadAxisCount(id int) GamepadAxis {
	return GamepadAxis(ebiten.GamepadAxisNum(id))
}

func GamepadAxisValue(id int, axis GamepadAxis) float64 {
	return ebiten.GamepadAxis(id, int(axis-1))
}

func gamepadUpdate() {
	for deviceId := 0; deviceId < maxGamepads; deviceId++ {
		for i, _ := range gamepadStateList[deviceId] {
			ebitenKey := gamepadToEbiten[i]
			if ebitenKey <= 0 {
				// Ignore empty or special key codes
				continue
			}
			if !ebiten.IsGamepadButtonPressed(deviceId, ebitenKey) {
				gamepadStateList[deviceId][i] = gamepadNotHeld
				continue
			}
			switch gamepadStateList[deviceId][i] {
			case gamepadNotHeld:
				gamepadStateList[deviceId][i] = gamepadPressed
			case gamepadPressed:
				gamepadStateList[deviceId][i] = gamepadHeld
			}
		}
	}
}
