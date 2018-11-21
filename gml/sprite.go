package gml

import (
	"github.com/silbinarywolf/gml-go/gml/internal/sprite"
)

// todo(Jake): 2018-10-27
// Consider changing this to `type SpriteIndex int` and exposing
// Sprite functions by accessing the assetList.
type Sprite = sprite.Sprite

type SpriteState = sprite.SpriteState

func SpriteLoad(name string) *Sprite {
	return sprite.LoadSprite(name)
}
