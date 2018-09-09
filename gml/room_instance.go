package gml

import (
	"github.com/silbinarywolf/gml-go/gml/internal/object"
	"github.com/silbinarywolf/gml-go/gml/internal/room"
)

type RoomInstance struct {
	used  bool
	index int
	room  *Room

	instanceLayers []RoomInstanceLayerInstance
	spriteLayers   []RoomInstanceLayerSprite
	drawLayers     []RoomInstanceLayerDraw
}

func RoomInstanceCreate(room *Room) *RoomInstance {
	roomInst := gState.createNewRoomInstance(room)
	return roomInst
}

func RoomInstanceDestroy(roomInst *RoomInstance) {
	gState.deleteRoomInstance(roomInst)
}

func RoomInstanceEmptyCreate() *RoomInstance {
	roomInst := gState.createNewRoomInstance(nil)
	return roomInst
}

func (roomInst *RoomInstance) Index() int {
	return roomInst.index
}

func (roomInst *RoomInstance) Room() *room.Room {
	return roomInst.room
}

type roomInstanceObject interface {
	BaseObject() *object.Object
}

func RoomInstanceInstances(inst roomInstanceObject) []object.ObjectType {
	roomInstanceIndex := object.RoomInstanceIndex(inst.BaseObject())
	roomInst := RoomGetInstance(roomInstanceIndex)
	if roomInst == nil {
		return nil
	}
	instanceLayer := &roomInst.instanceLayers[len(roomInst.instanceLayers)-1]
	return instanceLayer.manager.instances
}

// NOTE(Jake):2018-08-19
//
// I might want to make this private so a user
// can only manipulate a room instance via functions
//
func RoomGetInstance(roomInstanceIndex int) *RoomInstance {
	roomInst := &gState.roomInstances[roomInstanceIndex]
	if roomInst.used {
		return roomInst
	}
	return nil
}

// todo(Jake): 2018-07-22
// Figure out this
/*func (roomInst *RoomInstance) InstanceCreateLayer(position Vec, layer *RoomInstanceLayerInstance, objectIndex object.ObjectIndex) object.ObjectType {

}

func (roomInst *RoomInstance) InstanceCreate(position Vec, objectIndex object.ObjectIndex) object.ObjectType {
	return roomInst.instanceManager.InstanceCreate(position, objectIndex, roomInst.Index())
}

func (roomInst *RoomInstance) InstanceDestroy(inst object.ObjectType) {
	manager := &roomInst.instanceManager
	manager.InstanceDestroy(inst)
}*/

func (roomInst *RoomInstance) update(animationUpdate bool) {
	for _, layer := range roomInst.instanceLayers {
		layer.update(animationUpdate)
	}
}

func (roomInst *RoomInstance) draw() {
	for _, layer := range roomInst.drawLayers {
		layer.draw()
	}
}
