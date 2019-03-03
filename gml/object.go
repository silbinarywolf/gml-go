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

/*type objectInteral struct {
	objectIndex       ObjectIndex
	depth             int
	solid             bool
}*/

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

func (inst Object) MarshalBinary() ([]byte, error) {
	w := objectSerialize{
		Rect:              inst.Rect,
		SpriteState:       inst.SpriteState,
		BboxOffset:        inst.bboxOffset,
		InstanceIndex:     inst.instanceIndex,
		RoomInstanceIndex: inst.roomInstanceIndex,
		ObjectIndex:       inst.objectIndex,
		Depth:             inst.depth,
		Solid:             inst.solid,
		ImageAngleRadians: inst.imageAngleRadians,
	}
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(w); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (inst *Object) UnmarshalBinary(data []byte) error {
	w := objectSerialize{}
	reader := bytes.NewReader(data)
	dec := gob.NewDecoder(reader)
	if err := dec.Decode(&w); err != nil {
		return err
	}
	inst.Rect = w.Rect
	inst.SpriteState = w.SpriteState
	inst.bboxOffset = w.BboxOffset
	inst.instanceIndex = w.InstanceIndex
	inst.roomInstanceIndex = w.RoomInstanceIndex
	inst.objectIndex = w.ObjectIndex
	inst.depth = w.Depth
	inst.solid = w.Solid
	inst.imageAngleRadians = w.ImageAngleRadians
	return nil
}
