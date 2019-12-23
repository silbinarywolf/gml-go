// +build !headless

package gml

import (
	"github.com/hajimehoshi/ebiten"
)

type keyState int

const (
	keyNotHeld keyState = 0 + iota
	keyPressed
	keyHeld
)

var (
	keyStateList [VkSize]keyState
)

func KeyboardCheck(key Key) bool {
	return ebiten.IsKeyPressed(keyboardVkToEbiten[key])
}

func KeyboardCheckPressed(key Key) bool {
	return keyStateList[key] == keyPressed
}

func keyboardUpdate() {
	for i, _ := range keyStateList {
		ebitenKey := keyboardVkToEbiten[i]
		if ebitenKey <= 0 {
			// Ignore empty or special key codes
			continue
		}
		if !ebiten.IsKeyPressed(ebitenKey) {
			keyStateList[i] = keyNotHeld
			continue
		}
		switch keyStateList[i] {
		case keyNotHeld:
			keyStateList[i] = keyPressed
		case keyPressed:
			keyStateList[i] = keyHeld
		}
	}
}
