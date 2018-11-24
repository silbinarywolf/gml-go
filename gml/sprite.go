package gml

import (
	"github.com/silbinarywolf/gml-go/gml/internal/sprite"
)

// todo(Jake): 2018-11-24 - Github Issue #2
// Remove Sprite in favour of exposing "SpriteIndex"
type Sprite = sprite.Sprite
type SpriteIndex = sprite.SpriteIndex

type SpriteState = sprite.SpriteState

func SpriteInitializeIndexToName(indexToName []string, nameToIndex map[string]SpriteIndex) {
	sprite.SpriteInitializeIndexToName(indexToName, nameToIndex)
}

func SpriteLoad(index SpriteIndex) *Sprite {
	return sprite.SpriteLoad(index)
}
