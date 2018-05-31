package gml

import (
	"math"
)

type ObjectIndex int32

type ObjectType interface {
	BaseObject() *Object
	ID() ObjectIndex
	Create()
	Update()
	Draw()
}

type Object struct {
	SpriteState               // Sprite (contains SetSprite)
	Vec                       // Position (contains X,Y)
	Size              Vec     // Size (X,Y)
	index             int     // index in the 'entities' array
	imageAngleRadians float64 // Image Angle
	imageScale        Vec
}

func (inst *Object) Create() {
	inst.imageScale.X = 1.0
	inst.imageScale.Y = 1.0
}

func (inst *Object) BaseObject() *Object        { return inst }
func (inst *Object) ImageAngleRadians() float64 { return inst.imageAngleRadians }
func (inst *Object) ImageAngle() float64        { return inst.imageAngleRadians * (180 / math.Pi) }
func (inst *Object) ImageScale() Vec            { return inst.imageScale }

func (inst *Object) SetSprite(sprite *Sprite) {
	inst.SpriteState.SetSprite(sprite)

	// Infer width and height if they aren't manually set
	// (This might be a bad idea, too magic! But feels like Game Maker, so...)
	if inst.Size.X == 0 {
		inst.Size.X = sprite.size.X
	}
	if inst.Size.Y == 0 {
		inst.Size.Y = sprite.size.Y
	}
}

func (inst *Object) SetImageAngle(angleInDegrees float64) {
	inst.imageAngleRadians = angleInDegrees * (math.Pi / 180)
}

func (inst *Object) SetImageAngleRadians(angleInRadians float64) {
	inst.imageAngleRadians = angleInRadians
}

func (inst *Object) DrawSelf() {
	pos := inst.Vec
	state := inst.SpriteState
	state.sprite.DrawSprite(state.imageIndex, pos)
}
