package gml

import (
	"github.com/silbinarywolf/gml-go/gml/internal/sprite"
)

type Sprite = sprite.Sprite

type SpriteState = sprite.SpriteState

func LoadSprite(name string) *sprite.Sprite {
	return sprite.LoadSprite(name)
}
