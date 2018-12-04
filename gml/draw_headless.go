// +build headless

package gml

import (
	"image/color"

	"github.com/silbinarywolf/gml-go/gml/internal/sprite"
)

func DrawGetGUI() bool {
	return false
}

func DrawSetGUI(guiMode bool) {
}

func DrawSprite(spr sprite.SpriteIndex, subimage float64, position Vec) {
}

func DrawSpriteScaled(spr sprite.SpriteIndex, subimage float64, position Vec, scale Vec) {
}

func DrawSpriteExt(spr sprite.SpriteIndex, subimage float64, position Vec, scale Vec, alpha float64) {
}

func DrawRectangle(pos Vec, size Vec, col color.Color) {
}

func DrawRectangleBorder(position Vec, size Vec, color color.Color, borderSize float64, borderColor color.Color) {
}

func DrawText(position Vec, message string) {
}

func DrawTextF(position Vec, message string, args ...interface{}) {
}
