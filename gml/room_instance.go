package gml

import "reflect"

type RoomInstance struct {
	room            *Room
	instanceManager instanceManager
}

func (roomInst *RoomInstance) InstanceCreate(position Vec, entityID ObjectIndex) ObjectType {
	if entityID == 0 {
		panic("Cannot pass 0 as 2nd parameter to InstanceCreate(position, entityID)")
	}
	valToCopy := gInstanceManager.idToEntityData[entityID]

	// Create and add to entity list
	manager := &roomInst.instanceManager
	inst := reflect.New(reflect.ValueOf(valToCopy).Elem().Type()).Interface().(ObjectType)
	index := len(manager.entities)
	manager.entities = append(manager.entities, inst)

	// Initialize object
	baseObj := inst.BaseObject()
	baseObj.index = index
	baseObj.room = roomInst
	baseObj.Create()
	inst.Create()

	// Set at instance create position
	baseObj.Vec = position
	return inst
}

func (roomInst *RoomInstance) update() {
	roomInst.instanceManager.update()
}

func (roomInst *RoomInstance) draw() {
	roomInst.instanceManager.draw()
}
