package gml

import (
	"github.com/silbinarywolf/gml-go/gml/internal/object"
	"github.com/silbinarywolf/gml-go/gml/internal/room"
)

type RoomInstance struct {
	used  bool
	index int
	room  *room.Room

	instanceLayers []roomInstanceLayerInstance
	spriteLayers   []roomInstanceLayerSprite
	drawLayers     []roomInstanceLayerDraw
}

// todo(Jake): 2018-12-01: Remove this if it feels unnecessary or goes unused
// RoomInstanceName get the name of the room used by the room instance
/*func RoomInstanceName(roomInstanceIndex int) string {
	roomInst := &gState.roomInstances[roomInstanceIndex]
	if !roomInst.used {
		return ""
	}
	return roomInst.room.Config.UUID
}*/

// RoomInstanceNew create a new empty room instance programmatically
func RoomInstanceNew() int {
	roomInst := gState.createNewRoomInstance(nil)
	return roomInst.index
}

// RoomInstanceCreate will create a new instance of the room given
// todo(Jake): 2018-12-01: #6: Change *Room to be gml.RoomIndex
func todo__roomInstanceCreate(room *room.Room) int {
	roomInst := gState.createNewRoomInstance(room)
	return roomInst.index
}

// RoomInstanceDestroy destroys a room instance
func RoomInstanceDestroy(roomInstanceIndex int) {
	roomInst := &gState.roomInstances[roomInstanceIndex]
	gState.deleteRoomInstance(roomInst)
}

// todo(Jake): 2018-12-01: Github #19: Remove this
func (roomInst *RoomInstance) Index() int {
	return roomInst.index
}

type roomInstanceObject interface {
	BaseObject() *object.Object
}

/*func RoomInstanceInstances(inst roomInstanceObject) []object.ObjectType {
	roomInstanceIndex := object.RoomInstanceIndex(inst.BaseObject())
	roomInst := RoomGetInstance(roomInstanceIndex)
	if roomInst == nil {
		return nil
	}
	instanceLayer := &roomInst.instanceLayers[len(roomInst.instanceLayers)-1]
	return instanceLayer.manager.instances
}*/

func roomGetInstance(roomInstanceIndex int) *RoomInstance {
	roomInst := &gState.roomInstances[roomInstanceIndex]
	if roomInst.used {
		return roomInst
	}
	return nil
}

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
