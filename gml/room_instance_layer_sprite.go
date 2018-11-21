package gml

import (
	"github.com/silbinarywolf/gml-go/gml/internal/geom"
	"github.com/silbinarywolf/gml-go/gml/internal/sprite"
)

type RoomInstanceLayerSpriteObject struct {
	geom.Vec
	Sprite *sprite.Sprite
}

func (record *RoomInstanceLayerSpriteObject) Rect() geom.Rect {
	r := geom.Rect{}
	r.Vec = record.Vec
	r.Size = record.Sprite.Size()
	return r
}

type RoomInstanceLayerSprite struct {
	RoomInstanceLayerDrawBase
	name    string
	sprites []RoomInstanceLayerSpriteObject
	//spaces       space.SpaceBucketArray
	hasCollision bool
}

func (layer *RoomInstanceLayerSprite) order() int32 {
	return layer.drawOrder
}

func (layer *RoomInstanceLayerSprite) draw() {
	//screen := gScreen
	for _, record := range layer.sprites {
		/*position := maybeApplyOffsetByCamera(record.Vec)
		frame := sprite.GetRawFrame(record.Sprite, 0) // int(math.Floor(subimage))
		op := ebiten.DrawImageOptions{}
		op.GeoM.Translate(position.X, position.Y)
		screen.DrawImage(frame, &op)*/
		DrawSprite(record.Sprite, 0, geom.Vec{record.X, record.Y})
	}
}
