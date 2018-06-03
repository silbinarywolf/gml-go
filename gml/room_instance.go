package gml

type RoomInstance struct {
	room            *Room
	instanceManager instanceManager
}

func (roomInst *RoomInstance) InstanceCreate(position Vec, objectIndex ObjectIndex) ObjectType {
	// Create and add to entity list
	manager := &roomInst.instanceManager
	inst := newInstance(objectIndex)
	baseObj := inst.BaseObject()
	baseObj.index = len(manager.entities)
	baseObj.room = roomInst
	manager.entities = append(manager.entities, inst)

	// Init and Set position
	inst.Create()
	baseObj.Vec = position
	return inst
}

func (roomInst *RoomInstance) update() {
	roomInst.instanceManager.update()
}

func (roomInst *RoomInstance) draw() {
	roomInst.instanceManager.draw()
}
