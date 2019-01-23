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

func DrawSprite(spr sprite.SpriteIndex, subimage float64, x, y float64) {
}

func DrawSpriteScaled(spr sprite.SpriteIndex, subimage float64, x, y float64, scale Vec) {
}

func DrawSpriteExt(spr sprite.SpriteIndex, subimage float64, x, y float64, scale Vec, alpha float64) {
}

func DrawRectangle(x, y, w, h float64, col color.Color) {
}

func DrawRectangleBorder(x, y, w, h float64, color color.Color, borderSize float64, borderColor color.Color) {
}

func DrawText(x, y float64, message string) {
}

func DrawTextColor(x, y float64, message string, col color.Color) {
}

func DrawTextF(x, y float64, message string, args ...interface{}) {
}
