// +build headless

package gml

import (
	"image/color"

	"github.com/silbinarywolf/gml-go/gml/internal/sprite"
)

func DrawSetGUI(guiMode bool) {
}

func DrawSprite(spr *sprite.Sprite, subimage float64, position Vec) {
}

func DrawSpriteScaled(spr *sprite.Sprite, subimage float64, position Vec, scale Vec) {
}

func DrawSpriteExt(spr *sprite.Sprite, subimage float64, position Vec, scale Vec, alpha float64) {
}

func DrawRectangle(pos Vec, size Vec, col color.Color) {
}

func DrawRectangleBorder(position Vec, size Vec, color color.Color, borderSize float64, borderColor color.Color) {
}

func DrawText(position Vec, message string) {
}
