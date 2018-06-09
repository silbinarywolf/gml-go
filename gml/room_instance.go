package gml

import (
	"github.com/silbinarywolf/gml-go/gml/internal/object"
)

var (
	// NOTE(Jake): 2018-06-09
	//
	// Starting middle param at "1" so that the 0th item always returns nil
	//
	roomInstances []*RoomInstance = make([]*RoomInstance, 1, 10)
)

func (room *Room) Create() *RoomInstance {
	roomInst := &RoomInstance{
		room: room,
	}
	index := len(roomInstances)
	roomInst.index = index
	roomInstances = append(roomInstances, roomInst)

	// Instantiate instances for this room
	for _, obj := range room.Instances {
		roomInst.InstanceCreate(V(float64(obj.X), float64(obj.Y)), object.ObjectIndex(obj.ObjectIndex))
	}
	return roomInst
}

type RoomInstance struct {
	index           int
	room            *Room
	instanceManager instanceManager
}

func (roomInst *RoomInstance) Index() int {
	return roomInst.index
}

func RoomGetInstance(roomInstanceIndex int) *RoomInstance {
	return roomInstances[roomInstanceIndex]
}

func (roomInst *RoomInstance) InstanceCreate(position Vec, objectIndex object.ObjectIndex) object.ObjectType {
	// Create and add to entity list
	manager := &roomInst.instanceManager
	index := len(manager.entities)
	inst := object.NewRawInstance(objectIndex, index, roomInst.Index())
	manager.entities = append(manager.entities, inst)

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

func (roomInst *RoomInstance) update() {
	roomInst.instanceManager.update()
}

func (roomInst *RoomInstance) draw() {
	roomInst.instanceManager.draw()
}
