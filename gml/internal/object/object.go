package object

import (
	"math"

	"github.com/silbinarywolf/gml-go/gml/internal/geom"
	"github.com/silbinarywolf/gml-go/gml/internal/sprite"
)

type ObjectIndex int32

type ObjectType interface {
	BaseObject() *Object
	ObjectIndex() ObjectIndex
	ObjectName() string
	Create()
	Destroy()
	Update()
	Draw()
}

type Object struct {
	sprite.SpriteState // Sprite (contains SetSprite)
	geom.Rect
	instanceObject
	objectIndex       ObjectIndex
	solid             bool
	imageAngleRadians float64 // Image Angle
}

func (inst *Object) create() {
	inst.ImageScale.X = 1.0
	inst.ImageScale.Y = 1.0
}

func (inst *Object) SetSolid(isSolid bool) {
	inst.solid = isSolid
}

func (inst *Object) Solid() bool                { return inst.solid }
func (inst *Object) BaseObject() *Object        { return inst }
func (inst *Object) ObjectIndex() ObjectIndex   { return inst.objectIndex }
func (inst *Object) ImageAngleRadians() float64 { return inst.imageAngleRadians }
func (inst *Object) ImageAngle() float64        { return inst.imageAngleRadians * (180 / math.Pi) }

//func (inst *Object) ImageScale() geom.Vec          { return inst.imageScale }

func (inst *Object) SetSprite(spriteIndex sprite.SpriteIndex) {
	inst.SpriteState.SetSprite(spriteIndex)

	// Infer width and height if they aren't manually set
	// (This might be a bad idea, too magic! But feels like Game Maker, so...)
	size := spriteIndex.Size()
	if inst.Size.X == 0 {
		inst.Size.X = size.X
	}
	if inst.Size.Y == 0 {
		inst.Size.Y = size.Y
	}
}

func (inst *Object) SetImageAngle(angleInDegrees float64) {
	inst.imageAngleRadians = angleInDegrees * (math.Pi / 180)
}

func (inst *Object) SetImageAngleRadians(angleInRadians float64) {
	inst.imageAngleRadians = angleInRadians
}

func (inst *Object) CollisionInstance(otherInst ObjectType) bool {
	return inst.Rect.CollisionRectangle(otherInst.BaseObject().Rect)
}
