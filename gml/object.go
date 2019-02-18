package gml

import (
	"math"

	"github.com/silbinarywolf/gml-go/gml/internal/geom"
	"github.com/silbinarywolf/gml-go/gml/internal/sprite"
)

/*func ObjectGetIndex(name string) (ObjectIndex, bool) {
	res, ok := ObjectGetIndex(name)
	return res, ok
}*/

type ObjectIndex int32

type ObjectType interface {
	BaseObject() *Object // remove this or unexport it once ObjectIndex is used in the engine
	ObjectIndex() ObjectIndex
	ObjectName() string
	Create()
	Destroy()
	Update()
	Draw()
}

type Object struct {
	geom.Rect
	sprite.SpriteState // Sprite (contains SetSprite)
	bboxOffset         geom.Vec
	instanceObject
	objectIndex       ObjectIndex
	depth             int
	solid             bool
	imageAngleRadians float64 // Image Angle
}

func (inst *Object) Create() {}

func (inst *Object) Update() {}

func (inst *Object) Destroy() {}

func (inst *Object) Draw() {
	DrawSprite(inst.SpriteIndex(), inst.ImageIndex(), inst.X, inst.Y)
}

func (inst *Object) Bbox() geom.Rect {
	return geom.Rect{
		Vec: geom.Vec{
			X: inst.X + inst.bboxOffset.X,
			Y: inst.Y + inst.bboxOffset.Y,
		},
		Size: inst.Size,
	}
}

func (inst *Object) bboxAt(x, y float64) geom.Rect {
	return geom.Rect{
		Vec: geom.Vec{
			X: x + inst.bboxOffset.X,
			Y: y + inst.bboxOffset.Y,
		},
		Size: inst.Size,
	}
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
func (inst *Object) ObjectName() string         { return gObjectManager.indexToName[inst.objectIndex] }
func (inst *Object) ObjectIndex() ObjectIndex   { return inst.objectIndex }
func (inst *Object) ImageAngleRadians() float64 { return inst.imageAngleRadians }
func (inst *Object) ImageAngle() float64        { return inst.imageAngleRadians * (180 / math.Pi) }

// Depth will get the draw order of the object
func (inst *Object) Depth() int { return inst.depth }

// SetDepth will change the draw order of the object
func (inst *Object) SetDepth(depth int) {
	inst.depth = depth
}

// SetSprite will change the image used to draw the object
func (inst *Object) SetSprite(spriteIndex sprite.SpriteIndex) {
	var oldSize geom.Vec
	if oldSpriteIndex := inst.SpriteIndex(); oldSpriteIndex != sprite.SprUndefined {
		oldSize = oldSpriteIndex.Size()
	}

	inst.SpriteState.SetSprite(spriteIndex)

	// Infer width and height if they aren't manually set
	// (This might be a bad idea, too magic! But feels like Game Maker, so...)
	size := inst.Size
	if size.X == oldSize.X &&
		size.Y == oldSize.Y {
		rect := sprite.SpriteCollisionMask(spriteIndex)
		inst.bboxOffset = rect.Vec
		inst.Size = rect.Size
	}
}

func (inst *Object) SetImageAngle(angleInDegrees float64) {
	inst.imageAngleRadians = angleInDegrees * (math.Pi / 180)
}

func (inst *Object) SetImageAngleRadians(angleInRadians float64) {
	inst.imageAngleRadians = angleInRadians
}
