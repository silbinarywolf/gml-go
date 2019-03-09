package gml

import (
	"bytes"
	"encoding/gob"
	"math"

	"github.com/silbinarywolf/gml-go/gml/internal/assert"
	"github.com/silbinarywolf/gml-go/gml/internal/geom"
	"github.com/silbinarywolf/gml-go/gml/internal/sprite"
)

/*func ObjectGetIndex(name string) (ObjectIndex, bool) {
	res, ok := ObjectGetIndex(name)
	return res, ok
}*/

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

type objectExternal struct {
	geom.Rect
	sprite.SpriteState // Sprite (contains SetSprite)
}

type objectInternal struct {
	BboxOffset        geom.Vec
	IsDestroyed       bool
	InstanceIndex     InstanceIndex     // global uuid
	RoomInstanceIndex RoomInstanceIndex // Room Instance Index belongs to
	ObjectIndex       ObjectIndex
	Depth             int
	Solid             bool
	ImageAngleRadians float64 // Image Angle
}

type Object struct {
	objectExternal
	internal objectInternal
}

type objectSerialize struct {
	Rect        geom.Rect
	SpriteState sprite.SpriteState
	Internal    objectInternal
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
			X: inst.X + inst.internal.BboxOffset.X,
			Y: inst.Y + inst.internal.BboxOffset.Y,
		},
		Size: inst.Size,
	}
}

func (inst *Object) bboxAt(x, y float64) geom.Rect {
	return geom.Rect{
		Vec: geom.Vec{
			X: x + inst.internal.BboxOffset.X,
			Y: y + inst.internal.BboxOffset.Y,
		},
		Size: inst.Size,
	}
}

func (inst *Object) create() {
	inst.ImageScale.X = 1.0
	inst.ImageScale.Y = 1.0
}

func (inst *Object) SetSolid(isSolid bool) {
	inst.internal.Solid = isSolid
}

func (inst *Object) Solid() bool                { return inst.internal.Solid }
func (inst *Object) BaseObject() *Object        { return inst }
func (inst *Object) ObjectName() string         { return gObjectManager.indexToName[inst.internal.ObjectIndex] }
func (inst *Object) ObjectIndex() ObjectIndex   { return inst.internal.ObjectIndex }
func (inst *Object) ImageAngleRadians() float64 { return inst.internal.ImageAngleRadians }
func (inst *Object) ImageAngle() float64        { return inst.internal.ImageAngleRadians * (180 / math.Pi) }

// Depth will get the draw order of the object
func (inst *Object) Depth() int { return inst.internal.Depth }

// SetDepth will change the draw order of the object
func (inst *Object) SetDepth(depth int) {
	inst.internal.Depth = depth
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
		inst.internal.BboxOffset = rect.Vec
		inst.Size = rect.Size
	}
}

func (inst *Object) SetImageAngle(angleInDegrees float64) {
	inst.internal.ImageAngleRadians = angleInDegrees * (math.Pi / 180)
}

func (inst *Object) SetImageAngleRadians(angleInRadians float64) {
	inst.internal.ImageAngleRadians = angleInRadians
}

func (inst Object) MarshalBinaryObject() ([]byte, error) {
	if inst.internal.RoomInstanceIndex == 0 {
		panic("RoomInstanceIndex cannot be 0")
	}
	w := objectSerialize{
		Rect:        inst.Rect,
		SpriteState: inst.SpriteState,
		Internal:    inst.internal,
	}
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(w); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (inst *Object) UnmarshalBinaryObject(data []byte) error {
	w := objectSerialize{}
	reader := bytes.NewReader(data)
	dec := gob.NewDecoder(reader)
	if err := dec.Decode(&w); err != nil {
		return err
	}
	prevRoomIndex := inst.internal.RoomInstanceIndex

	inst.Rect = w.Rect
	inst.SpriteState = w.SpriteState
	inst.internal = w.Internal

	assert.DebugAssert(inst.X == 0, "X cannot be 0")
	assert.DebugAssert(inst.Y == 0, "Y cannot be 0")
	assert.DebugAssert(inst.internal.InstanceIndex == 0, "InstanceIndex cannot be 0")
	assert.DebugAssert(inst.internal.RoomInstanceIndex == 0, "RoomInstanceIndex cannot be 0")

	// NOTE: Jake: 2019-03-09
	// This is incorrect behaviour and a hack.
	// InstanceRestore should be updating the room that the player is in based on
	// byte data. I need to do this in the near future.
	// Add to room
	if prevRoomIndex == 0 {
		roomInst := &roomInstanceState.roomInstances[inst.internal.RoomInstanceIndex]
		roomInst.instances = append(roomInst.instances, inst.internal.InstanceIndex)
	}

	return nil
}
