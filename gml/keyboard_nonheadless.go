// +build !headless

package gml

import (
	"github.com/hajimehoshi/ebiten"
)

var (
	pressingKeyLastFrame [VkSize]bool
)

func KeyboardCheck(key Key) bool {
	return ebiten.IsKeyPressed(keyboardVkToEbiten[key])
}

func KeyboardCheckPressed(key Key) bool {
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
