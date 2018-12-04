package gml

import (
	"github.com/silbinarywolf/gml-go/gml/internal/geom"
	"github.com/silbinarywolf/gml-go/gml/internal/sprite"
)

func DrawSelf(state *sprite.SpriteState, position geom.Vec) {
	DrawSpriteScaled(state.SpriteIndex(), state.ImageIndex(), position, state.ImageScale)
}
