package gml

import (
	"github.com/silbinarywolf/gml-go/gml/internal/sprite"
)

func DrawSelf(state *sprite.SpriteState, position Vec) {
	DrawSpriteExt(state.Sprite(), state.ImageIndex(), position, state.ImageScale)
}
