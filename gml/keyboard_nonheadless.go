// +build !headless

package gml

import (
	"github.com/hajimehoshi/ebiten"
)

var (
	pressingKeyLastFrame [vkSize]bool
)

func KeyboardCheck(key int16) bool {
	return ebiten.IsKeyPressed(keyboardVkToEbiten[key])
}

func KeyboardCheckPressed(key int16) bool {
	isHeld := KeyboardCheck(key)
	if !isHeld {
		pressingKeyLastFrame[key] = false
	}
	if pressingKeyLastFrame[key] {
		return false
	}
	if isHeld {
		pressingKeyLastFrame[key] = true
	}
	return isHeld
}
