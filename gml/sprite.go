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

func DrawSprite(spr *sprite.Sprite, subimage float64, position Vec) {
	screen := gScreen
	cameraPos := currentCamera.Vec

	frame := spr.GetFrame(int(subimage))
	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(position.X-cameraPos.X, position.Y-cameraPos.Y)
	screen.DrawImage(frame, &op)
}
