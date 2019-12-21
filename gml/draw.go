package gml

import (
	"github.com/silbinarywolf/gml-go/gml/internal/sprite"
)

type SpriteFrame struct {
	SpriteIndex sprite.SpriteIndex
	ImageIndex  float64
}

type DrawSpriteMaskOptions struct {
	Masks []SpriteFrame
}

func DrawSelf(state *sprite.SpriteState, x, y float64) {
	DrawSpriteScaled(state.SpriteIndex(), state.ImageIndex(), x, y, state.ImageScale)
}
