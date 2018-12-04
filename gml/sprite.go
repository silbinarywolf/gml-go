package gml

import (
	"github.com/silbinarywolf/gml-go/gml/internal/sprite"
)

type SpriteIndex = sprite.SpriteIndex

type SpriteState = sprite.SpriteState

func SpriteInitializeIndexToName(indexToName []string, nameToIndex map[string]SpriteIndex) {
	sprite.SpriteInitializeIndexToName(indexToName, nameToIndex)
}

func SpriteLoad(index SpriteIndex) SpriteIndex {
	sprite.SpriteLoad(index)
	return index
}
