package gml

import (
	"bytes"
	"encoding/gob"
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

type objectInternal struct {
	bboxOffset geom.Vec
	instanceObject
	objectIndex       ObjectIndex
	depth             int
	solid             bool
	imageAngleRadians float64 // Image Angle
}

type objectSerialize struct {
	Rect              geom.Rect
	SpriteState       sprite.SpriteState
	BboxOffset        geom.Vec
	InstanceIndex     InstanceIndex
	RoomInstanceIndex RoomInstanceIndex
	ObjectIndex       ObjectIndex
	Depth             int
	Solid             bool
	ImageAngleRadians float64
}

type Object struct {
	geom.Rect
	sprite.SpriteState // Sprite (contains SetSprite)
	internal           objectInternal
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
			X: inst.X + inst.internal.bboxOffset.X,
			Y: inst.Y + inst.internal.bboxOffset.Y,
		},
		Size: inst.Size,
	}
}

func (inst *Object) bboxAt(x, y float64) geom.Rect {
	return geom.Rect{
		Vec: geom.Vec{
			X: x + inst.internal.bboxOffset.X,
			Y: y + inst.internal.bboxOffset.Y,
		},
		Size: inst.Size,
	}
}

func (inst *Object) create() {
	inst.ImageScale.X = 1.0
	inst.ImageScale.Y = 1.0
}

func (inst *Object) SetSolid(isSolid bool) {
	inst.internal.solid = isSolid
}

func (inst *Object) Solid() bool                { return inst.internal.solid }
func (inst *Object) BaseObject() *Object        { return inst }
func (inst *Object) ObjectName() string         { return gObjectManager.indexToName[inst.internal.objectIndex] }
func (inst *Object) ObjectIndex() ObjectIndex   { return inst.internal.objectIndex }
func (inst *Object) ImageAngleRadians() float64 { return inst.internal.imageAngleRadians }
func (inst *Object) ImageAngle() float64        { return inst.internal.imageAngleRadians * (180 / math.Pi) }

// Depth will get the draw order of the object
func (inst *Object) Depth() int { return inst.internal.depth }

// SetDepth will change the draw order of the object
func (inst *Object) SetDepth(depth int) {
	inst.internal.depth = depth
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
		inst.internal.bboxOffset = rect.Vec
		inst.Size = rect.Size
	}
}

func (inst *Object) SetImageAngle(angleInDegrees float64) {
	inst.internal.imageAngleRadians = angleInDegrees * (math.Pi / 180)
}

func (inst *Object) SetImageAngleRadians(angleInRadians float64) {
	inst.internal.imageAngleRadians = angleInRadians
}

func (inst Object) MarshalBinaryField() ([]byte, error) {
	w := objectSerialize{
		Rect:              inst.Rect,
		SpriteState:       inst.SpriteState,
		BboxOffset:        inst.internal.bboxOffset,
		InstanceIndex:     inst.internal.instanceIndex,
		RoomInstanceIndex: inst.internal.roomInstanceIndex,
		ObjectIndex:       inst.internal.objectIndex,
		Depth:             inst.internal.depth,
		Solid:             inst.internal.solid,
		ImageAngleRadians: inst.internal.imageAngleRadians,
	}
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(w); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (inst *Object) UnmarshalBinaryField(data []byte) error {
	w := objectSerialize{}
	reader := bytes.NewReader(data)
	dec := gob.NewDecoder(reader)
	if err := dec.Decode(&w); err != nil {
		return err
	}
	inst.Rect = w.Rect
	inst.SpriteState = w.SpriteState
	inst.internal.bboxOffset = w.BboxOffset
	inst.internal.instanceIndex = w.InstanceIndex
	inst.internal.roomInstanceIndex = w.RoomInstanceIndex
	inst.internal.objectIndex = w.ObjectIndex
	inst.internal.depth = w.Depth
	inst.internal.solid = w.Solid
	inst.internal.imageAngleRadians = w.ImageAngleRadians
	return nil
}
