// +build headless

package gml

import (
	"image/color"

	"github.com/silbinarywolf/gml-go/gml/internal/sprite"
)

func DrawSprite(spr *sprite.Sprite, subimage float64, position Vec) {
}

func DrawSpriteExt(spr *sprite.Sprite, subimage float64, position Vec, scale Vec) {
}

func DrawRectangle(pos Vec, size Vec, col color.RGBA) {
}

func DrawText(position Vec, message string) {
}
