package gml

import (
	"math"

	m "github.com/silbinarywolf/gml-go/gml/internal/math"
	"github.com/silbinarywolf/gml-go/gml/internal/sprite"
)

type ObjectIndex int32

type ObjectType interface {
	BaseObject() *Object
	ObjectIndex() ObjectIndex
	Create()
	Update()
	Draw()
}

type Object struct {
	sprite.SpriteState       // Sprite (contains SetSprite)
	m.Vec                    // Position (contains X,Y)
	Size               m.Vec // Size (X,Y)
	index              int   // index in the 'entities' array
	room               *RoomInstance
	imageAngleRadians  float64 // Image Angle
	imageScale         m.Vec
}

func (inst *Object) Create() {
	inst.imageScale.X = 1.0
	inst.imageScale.Y = 1.0
}

func (inst *Object) BaseObject() *Object        { return inst }
func (inst *Object) Pos() m.Vec                 { return inst.Vec }
func (inst *Object) ImageAngleRadians() float64 { return inst.imageAngleRadians }
func (inst *Object) ImageAngle() float64        { return inst.imageAngleRadians * (180 / math.Pi) }
func (inst *Object) ImageScale() m.Vec          { return inst.imageScale }

func (inst *Object) SetSprite(sprite *sprite.Sprite) {
	inst.SpriteState.SetSprite(sprite)

	// Infer width and height if they aren't manually set
	// (This might be a bad idea, too magic! But feels like Game Maker, so...)
	if inst.Size.X == 0 {
		inst.Size.X = sprite.Size().X
	}
	if inst.Size.Y == 0 {
		inst.Size.Y = sprite.Size().Y
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
	spr := state.Sprite()
	DrawSprite(spr, state.ImageIndex(), pos)
}
