package gml

import (
	"github.com/silbinarywolf/gml-go/gml/internal/object"
)

func RoomInstanceCreate(room *Room) *RoomInstance {
	roomInst := gState.createNewRoomInstance(room)
	return roomInst
}

func RoomInstanceEmptyCreate() *RoomInstance {
	roomInst := gState.createNewRoomInstance(nil)
	return roomInst
}

type RoomInstance struct {
	used            bool
	index           int
	room            *Room
	instanceManager instanceManager
}

func (roomInst *RoomInstance) Index() int {
	return roomInst.index
}

/*func (roomInst *RoomInstance) CreateSnapshot() []byte {
	now := time.Now()
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(roomInst.instanceManager.instances)
	if err != nil {
		panic(err)
	}
	result := buf.Bytes()
	println("Time to encode:", time.Now().Sub(now).String(), ", Size in bytes:", len(result))
	return result
}

func (roomInst *RoomInstance) RestoreSnapshot(data []byte) {
	now := time.Now()
	var buf bytes.Buffer
	buf.Write(data)
	enc := gob.NewDecoder(&buf)
	err := enc.Decode(roomInst.instanceManager.instances)
	if err != nil {
		panic(err)
	}
	println("Time to decode:", time.Now().Sub(now).String())
}*/

func RoomGetInstance(roomInstanceIndex int) *RoomInstance {
	roomInst := &gState.roomInstances[roomInstanceIndex]
	if roomInst.used {
		return roomInst
	}
	return nil
}

func (roomInst *RoomInstance) InstanceCreate(position Vec, objectIndex object.ObjectIndex) object.ObjectType {
	// Create and add to entity list
	manager := &roomInst.instanceManager
	index := len(manager.instances)
	inst := object.NewRawInstance(objectIndex, index, roomInst.Index())
	manager.instances = append(manager.instances, inst)

	// Init and Set position
	inst.Create()
	baseObj := inst.BaseObject()
	baseObj.Vec = position
	return inst
}

func (roomInst *RoomInstance) InstanceDestroy(inst object.ObjectType) {
	manager := &roomInst.instanceManager
	manager.InstanceDestroy(inst)
}

func (roomInst *RoomInstance) update(animationUpdate bool) {
	roomInst.instanceManager.update(animationUpdate)
}

func (roomInst *RoomInstance) draw() {
	roomInst.instanceManager.draw()
}
