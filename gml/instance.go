package gml

import (
	"github.com/silbinarywolf/gml-go/gml/internal/geom"
	"github.com/silbinarywolf/gml-go/gml/internal/object"
)

type instanceManager struct {
	instances []object.ObjectType
}

func newInstanceManager() *instanceManager {
	manager := new(instanceManager)
	manager.reset()
	return manager
}

func (manager *instanceManager) reset() {
	*manager = instanceManager{}
}

func instanceCreateLayer(position geom.Vec, layer *roomInstanceLayerInstance, roomInst *RoomInstance, objectIndex object.ObjectIndex) object.ObjectType {
	return layer.manager.InstanceCreate(position, objectIndex, roomInst.Index(), layer.index)
}

func InstanceGet(index object.ObjectIndex) object.ObjectType {
	panic("todo: Implement InstanceGet()")
	return nil
}

func InstanceChangeRoom(inst object.ObjectType, roomInstanceIndex int) {
	roomInst := &gState.roomInstances[roomInstanceIndex]
	if !roomInst.used {
		return
	}
	// NOTE(Jake): 2018-07-22
	// For now instances default to the last instance layer
	layerIndex := len(roomInst.instanceLayers) - 1
	layer := &roomInst.instanceLayers[layerIndex]

	instanceRemove(inst)
	layer.manager.instanceAdd(inst, roomInst.Index(), layer.index)
}

func InstanceCreateRoom(position geom.Vec, roomInstanceIndex int, objectIndex object.ObjectIndex) object.ObjectType {
	roomInst := &gState.roomInstances[roomInstanceIndex]
	// NOTE(Jake): 2018-07-22
	// For now instances default to the last instance layer
	layerIndex := len(roomInst.instanceLayers) - 1
	layer := &roomInst.instanceLayers[layerIndex]
	//fmt.Printf("InstanceCreateRoom: Create on layer %d\n", layerIndex)
	return layer.manager.InstanceCreate(position, objectIndex, roomInst.Index(), layer.index)
}

func InstanceExists(inst object.ObjectType) bool {
	baseObj := inst.BaseObject()
	if baseObj == nil {
		return false
	}
	roomInst := roomGetInstance(baseObj.RoomInstanceIndex())
	// todo(Jake): 2018-08-20
	//
	// Check to see if current entity is destroyed
	//
	return roomInst != nil
}

func (manager *instanceManager) instanceAdd(inst object.ObjectType, roomInstanceIndex, layerIndex int) {
	// Move entity to new list
	index := len(manager.instances)
	object.MoveInstance(inst, index, roomInstanceIndex, layerIndex)
	manager.instances = append(manager.instances, inst)
}

func (manager *instanceManager) InstanceCreate(position geom.Vec, objectIndex object.ObjectIndex, roomInstanceIndex, layerIndex int) object.ObjectType {
	// Create and add to entity list
	index := len(manager.instances)

	// Get instance
	inst := object.NewRawInstance(objectIndex, index, roomInstanceIndex, layerIndex)
	manager.instances = append(manager.instances, inst)

	// Init and Set position
	inst.Create()
	inst.BaseObject().Vec = position
	return inst
}

func instanceRemove(inst object.ObjectType) {
	baseObj := inst.BaseObject()

	// Get slots
	roomInstanceIndex := baseObj.RoomInstanceIndex()
	layerIndex := object.LayerInstanceIndex(baseObj)
	index := object.InstanceIndex(baseObj)

	// Get manager
	roomInst := &gState.roomInstances[roomInstanceIndex]
	layerInst := &roomInst.instanceLayers[layerIndex]
	manager := &layerInst.manager

	if manager.instances[index] != inst {
		panic("instanceRemove failed as instance provided has already been removed")
	}
	// Get index
	/*index := -1
	for i, otherInst := range manager.instances {
		if inst == otherInst {
			index = i
		}
	}
	if index == -1 {
		panic("instanceRemove failed as instance provided has already been removed")
	}*/

	// Unordered delete
	// NOTE(Jake): 2018-09-15
	// Im aware this sometimes causes the server to crash...
	// but I also don't want to fix this yet as I might store each type of an
	// entity in its own bucket array soon...
	//
	// At the very least I should maybe make this a "mark as deleted"
	// system where it cleans up the entity list at the end of the frame.
	//
	lastEntry := manager.instances[len(manager.instances)-1]
	manager.instances[index] = lastEntry
	object.SetInstanceIndex(lastEntry.BaseObject(), index)
	manager.instances = manager.instances[:len(manager.instances)-1]
}

func InstanceDestroy(inst object.ObjectType) {
	baseObj := inst.BaseObject()
	if object.IsDestroyed(baseObj) {
		// NOTE(Jake): 2018-10-07
		// Maybe making this just silently returning will be better / less error
		// prone? For now lets be strict.
		panic("Cannot call InstanceDestroy on an object more than once.")
		return
	}

	// Run user-destroy code
	inst.Destroy()

	// Mark as destroyed
	object.MarkAsDestroyed(baseObj)

	// NOTE(Jake): 2018-10-07
	// Remove at the end of the frame (gState.update)
	gState.instancesMarkedForDelete = append(gState.instancesMarkedForDelete, inst)
}

func (manager *instanceManager) update(animationUpdate bool) {
	{
		instances := manager.instances
		for _, inst := range instances {
			if inst == nil {
				continue
			}
			inst.Update()
		}

		if animationUpdate {
			for _, inst := range instances {
				baseObj := inst.BaseObject()
				baseObj.SpriteState.ImageUpdate()
			}
		}
	}
}

func (manager *instanceManager) draw() {
	for _, inst := range manager.instances {
		if inst == nil {
			continue
		}
		inst.Draw()
	}
}
