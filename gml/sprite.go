package gml

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/silbinarywolf/gml-go/gml/internal/sprite"
)

type Sprite = sprite.Sprite

type SpriteState = sprite.SpriteState

func LoadSprite(name string) *sprite.Sprite {
	return sprite.LoadSprite(name)
}

func DrawSelf(state *sprite.SpriteState, position Vec) {
	DrawSprite(state.Sprite(), state.ImageIndex(), position)
}

func DrawSprite(spr *sprite.Sprite, subimage float64, position Vec) {
	screen := gScreen
	camPos := cameraGetActive().Vec

	frame := spr.GetFrame(int(subimage))
	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(position.X-camPos.X, position.Y-camPos.Y)
	screen.DrawImage(frame, &op)
}
