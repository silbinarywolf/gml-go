package gml

import (
	"bytes"
	"encoding/binary"
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
	BaseObject() *Object
	ObjectIndex() ObjectIndex
	ObjectName() string
	Create()
	Destroy()
	Update()
	Draw()
}

// ObjectSerialize hints to the code generator to generate serialization functions
type ObjectSerialize struct {
}

type objectExternal struct {
	geom.Rect
	sprite.SpriteState // Sprite (contains SetSprite)
}

type objectInternal struct {
	IsDestroyed       bool
	Solid             bool
	BboxOffset        geom.Vec
	InstanceIndex     InstanceIndex     // global uuid
	RoomInstanceIndex RoomInstanceIndex // Room Instance Index belongs to
	ObjectIndex       ObjectIndex
	Depth             int
	ImageAngleRadians float64 // Image Angle
}

type Object struct {
	internal objectInternal
	objectExternal
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

func (inst Object) UnsafeSnapshotMarshalBinary(buf *bytes.Buffer) error {
	if inst.internal.RoomInstanceIndex == 0 {
		panic("RoomInstanceIndex cannot be 0")
	}
	if err := binary.Write(buf, binary.LittleEndian, inst.internal.InstanceIndex); err != nil {
		return err
	}
	if err := binary.Write(buf, binary.LittleEndian, inst.internal.ObjectIndex); err != nil {
		return err
	}
	if err := binary.Write(buf, binary.LittleEndian, inst.internal.RoomInstanceIndex); err != nil {
		return err
	}
	if err := binary.Write(buf, binary.LittleEndian, inst.objectExternal.Rect); err != nil {
		return err
	}
	if err := inst.objectExternal.SpriteState.UnsafeSnapshotMarshalBinary(buf); err != nil {
		return err
	}
	if err := binary.Write(buf, binary.LittleEndian, inst.internal.BboxOffset); err != nil {
		return err
	}
	if err := binary.Write(buf, binary.LittleEndian, int32(inst.internal.Depth)); err != nil {
		return err
	}
	if err := binary.Write(buf, binary.LittleEndian, inst.internal.ImageAngleRadians); err != nil {
		return err
	}
	if err := binary.Write(buf, binary.LittleEndian, inst.internal.IsDestroyed); err != nil {
		return err
	}
	if err := binary.Write(buf, binary.LittleEndian, inst.internal.Solid); err != nil {
		return err
	}
	return nil
}

func (inst *Object) UnsafeSnapshotUnmarshalBinary(buf *bytes.Buffer) error {
	if err := binary.Read(buf, binary.LittleEndian, &inst.internal.InstanceIndex); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &inst.internal.ObjectIndex); err != nil {
		return err
	}
	var roomInstanceIndex RoomInstanceIndex
	if err := binary.Read(buf, binary.LittleEndian, &roomInstanceIndex); err != nil {
		return err
	}
	roomInstanceIndex.RoomInstanceChangeRoom(inst)
	if err := binary.Read(buf, binary.LittleEndian, &inst.objectExternal.Rect); err != nil {
		return err
	}
	if err := inst.objectExternal.SpriteState.UnsafeSnapshotUnmarshalBinary(buf); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &inst.internal.BboxOffset); err != nil {
		return err
	}
	var d int32
	if err := binary.Read(buf, binary.LittleEndian, &d); err != nil {
		return err
	}
	inst.internal.Depth = int(d)
	if err := binary.Read(buf, binary.LittleEndian, &inst.internal.ImageAngleRadians); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &inst.internal.IsDestroyed); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &inst.internal.Solid); err != nil {
		return err
	}
	return nil
}
