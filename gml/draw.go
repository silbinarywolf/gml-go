package gml

import (
	"github.com/silbinarywolf/gml-go/gml/internal/sprite"
)

func DrawSelf(state *sprite.SpriteState, x, y float64) {
	DrawSpriteScaled(state.SpriteIndex(), state.ImageIndex(), x, y, state.ImageScale)
}
