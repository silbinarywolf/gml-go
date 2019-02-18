package gml

import (
	"github.com/silbinarywolf/gml-go/gml/internal/sprite"
)

type roomInstanceLayerBackground struct {
	roomInstanceLayerDrawBase
	name      string
	sprite    sprite.SpriteIndex
	x, y      float64
	roomLeft  float64
	roomRight float64
}

func (layer *roomInstanceLayerBackground) order() int32 {
	return layer.drawOrder
}

func (layer *roomInstanceLayerBackground) draw() {
	sprite := layer.sprite
	width := float64(sprite.Size().X)
	x := layer.x
	y := layer.y
	DrawSprite(sprite, 0, x, y)
	{
		// Tile left
		x := x
		for x > float64(layer.roomLeft) {
			x -= width
			DrawSprite(sprite, 0, x, y)
		}
	}
	{
		// Tile left
		x := x
		for x < float64(layer.roomRight) {
			x += width
			DrawSprite(sprite, 0, x, y)
		}
	}
}
