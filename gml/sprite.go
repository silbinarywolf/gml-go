package gml

import (
	"github.com/silbinarywolf/gml-go/gml/internal/sprite"
)

type SpriteIndex = sprite.SpriteIndex

type SpriteState = sprite.SpriteState

// InitSpriteGeneratedData is used by code generated by gmlgo so you can query a sprite by index or name
func InitSpriteGeneratedData(indexToName []string, nameToIndex map[string]SpriteIndex, indexToPath []string) {
	testInitAssetDir()

	sprite.InitSpriteGeneratedData(indexToName, nameToIndex, indexToPath)
}

// SpriteLoad will ensure the sprite is loaded
func SpriteLoad(index SpriteIndex) SpriteIndex {
	sprite.SpriteLoad(index)
	return index
}
