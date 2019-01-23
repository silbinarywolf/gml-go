package gml

import (
	"github.com/silbinarywolf/gml-go/gml/internal/geom"
	"github.com/silbinarywolf/gml-go/gml/internal/sprite"
)

type roomInstanceLayerSpriteObject struct {
	geom.Vec
	sprite sprite.SpriteIndex
}

func (record *roomInstanceLayerSpriteObject) Rect() geom.Rect {
	r := geom.Rect{}
	r.Vec = record.Vec
	r.Size = record.sprite.Size()
	return r
}

type roomInstanceLayerSprite struct {
	roomInstanceLayerDrawBase
	name    string
	sprites []roomInstanceLayerSpriteObject
	//spaces       space.SpaceBucketArray
	hasCollision bool
}

func (layer *roomInstanceLayerSprite) order() int32 {
	return layer.drawOrder
}

func (layer *roomInstanceLayerSprite) draw() {
	//screen := gScreen
	for _, record := range layer.sprites {
		/*position := maybeApplyOffsetByCamera(record.Vec)
		frame := sprite.GetRawFrame(record.Sprite, 0) // int(math.Floor(subimage))
		op := ebiten.DrawImageOptions{}
		op.GeoM.Translate(position.X, position.Y)
		screen.DrawImage(frame, &op)*/
		DrawSprite(record.sprite, 0, record.X, record.Y)
	}
}
