package gml

import (
	"github.com/hajimehoshi/ebiten"
)

func KeyboardCheck(key int16) bool {
	return ebiten.IsKeyPressed(keyboardVkToEbiten[key])
}
