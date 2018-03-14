package gml

import (
	"math"

	"github.com/hajimehoshi/ebiten"
)

type EntityID int

type EntityType interface {
	BaseEntity() *Entity
	ID() EntityID
	Create()
	Update()
	Draw()
}

type Entity struct {
	SpriteState               // Sprite (contains SetSprite)
	Vec                       // Position (contains X,Y)
	imageAngleRadians float64 // Image Angle
	imageScale        Vec
}

func (e *Entity) init() {
	e.imageScale.X = 1.0
	e.imageScale.Y = 1.0
}

func (e *Entity) BaseEntity() *Entity        { return e }
func (e *Entity) ImageAngleRadians() float64 { return e.imageAngleRadians }
func (e *Entity) ImageAngle() float64        { return e.imageAngleRadians * (180 / math.Pi) }
func (e *Entity) ImageScale() Vec            { return e.imageScale }

func (e *Entity) SetImageAngle(angleInDegrees float64) {
	e.imageAngleRadians = angleInDegrees * (math.Pi / 180)
}

func (e *Entity) SetImageAngleRadians(angleInRadians float64) {
	e.imageAngleRadians = angleInRadians
}

func (e *Entity) DrawSelf() {
	imageIndex := int(e.SpriteState.imageIndex)
	sprite := e.Sprite().frames[imageIndex]

	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(e.X, e.Y)
	op.GeoM.Scale(e.imageScale.X, e.imageScale.Y)
	g_screen.DrawImage(sprite, &op)
}
