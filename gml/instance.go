package gml

import (
	"reflect"

	"github.com/silbinarywolf/gml-go/gml/internal/geom"
)

type instanceManager struct {
	instances   []ObjectType
	indexToData map[ObjectIndex]int
}

func allocateNewInstance(objectIndex ObjectIndex) ObjectType {
	manager := &gState.instanceManager

	// Allocate new instance
	valToCopy := gObjectManager.idToEntityData[objectIndex]
	if valToCopy == nil {
		panic("Invalid objectIndex")
	}
	// todo(Jake): 2018-12-18
	// Explore allocating from a large fixed-size pool
	inst := reflect.New(reflect.ValueOf(valToCopy).Elem().Type()).Interface().(ObjectType)
	baseObj := inst.BaseObject()
	manager.instances = append(manager.instances, inst)
	slot := len(manager.instances) - 1
	baseObj.index = slot
	return manager.instances[slot]
}

func (manager *instanceManager) Get(index ObjectIndex) ObjectType {
	dataIndex, ok := manager.indexToData[index]
	if !ok {
		return nil
	}
	inst := manager.instances[dataIndex]
	if inst.BaseObject().isDestroyed {
		return nil
	}
	return inst
}

// todo: Jake: 2018-12-16
// Deprecate this in favour of one storage area for all entities
type roomInstanceManager struct {
	instances []ObjectType
}

type instanceObject struct {
	isDestroyed        bool
	index              int               // index in the 'entities' array
	roomInstanceIndex  RoomInstanceIndex // Room Instance Index belongs to
	layerInstanceIndex int               // Layer belongs to
}

func (inst *Object) RoomInstanceIndex() RoomInstanceIndex {
	return inst.roomInstanceIndex
}

func IsDestroyed(inst *Object) bool {
	return inst.isDestroyed
}

func MarkAsDestroyed(inst *Object) {
	inst.isDestroyed = true
}

func SetInstanceIndex(inst *Object, index int) {
	inst.index = index
}

func InstanceIndex(inst *Object) int {
	return inst.index
}

func LayerInstanceIndex(inst *Object) int {
	return inst.layerInstanceIndex
}

func newroomInstanceManager() *roomInstanceManager {
	manager := new(roomInstanceManager)
	manager.reset()
	return manager
}

func (manager *roomInstanceManager) reset() {
	*manager = roomInstanceManager{}
}

func instanceCreateLayer(position geom.Vec, layer *roomInstanceLayerInstance, roomInst *roomInstance, objectIndex ObjectIndex) ObjectType {
	return layer.manager.InstanceCreate(position, objectIndex, roomInst.index, layer.index)
}

func InstanceGet(index ObjectIndex) ObjectType {
	panic("todo: Implement InstanceGet()")
	return nil
}

func InstanceChangeRoom(inst ObjectType, roomInstanceIndex int) {
	roomInst := &gState.roomInstances[roomInstanceIndex]
	if !roomInst.used {
		return
	}
	// NOTE(Jake): 2018-07-22
	// For now instances default to the last instance layer
	layerIndex := len(roomInst.instanceLayers) - 1
	layer := &roomInst.instanceLayers[layerIndex]

	instanceRemove(inst)
	layer.manager.instanceAdd(inst, roomInst.index, layer.index)
}

func InstanceCreateRoom(position geom.Vec, roomInstanceIndex RoomInstanceIndex, objectIndex ObjectIndex) ObjectType {
	inst := allocateNewInstance(objectIndex)
	{
		baseObj := inst.BaseObject()
		baseObj.Vec = position
		baseObj.objectIndex = objectIndex
		baseObj.roomInstanceIndex = roomInstanceIndex
		inst.Create()
	}

	return inst
	/*roomInst := &gState.roomInstances[roomInstanceIndex]
	// NOTE(Jake): 2018-07-22
	// For now instances default to the last instance layer
	layerIndex := len(roomInst.instanceLayers) - 1
	layer := &roomInst.instanceLayers[layerIndex]
	layer.manager.instances =
	//manager.instances = append(manager.instances, inst)
	//fmt.Printf("InstanceCreateRoom: Create on layer %d\n", layerIndex)
	return layer.manager.InstanceCreate(position, objectIndex, roomInst.index, layer.index)*/
}

func InstanceExists(inst ObjectType) bool {
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

func (manager *roomInstanceManager) instanceAdd(inst ObjectType, roomInstanceIndex RoomInstanceIndex, layerIndex int) {
	// Move entity to new list
	index := len(manager.instances)
	moveInstance(inst, index, roomInstanceIndex, layerIndex)
	manager.instances = append(manager.instances, inst)
}

func (manager *roomInstanceManager) InstanceCreate(position geom.Vec, objectIndex ObjectIndex, roomInstanceIndex RoomInstanceIndex, layerIndex int) ObjectType {

	// Create and add to entity list
	index := len(manager.instances)

	// Get instance
	inst := newRawInstance(objectIndex, index, roomInstanceIndex, layerIndex)
	manager.instances = append(manager.instances, inst)

	// Init and Set position
	inst.Create()
	inst.BaseObject().Vec = position
	return inst
}

func instanceRemove(inst ObjectType) {
	baseObj := inst.BaseObject()

	// Get slots
	roomInstanceIndex := baseObj.RoomInstanceIndex()
	layerIndex := LayerInstanceIndex(baseObj)
	index := InstanceIndex(baseObj)

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
	SetInstanceIndex(lastEntry.BaseObject(), index)
	manager.instances = manager.instances[:len(manager.instances)-1]
}

func InstanceDestroy(inst ObjectType) {
	baseObj := inst.BaseObject()
	if IsDestroyed(baseObj) {
		// NOTE(Jake): 2018-10-07
		// Maybe making this just silently returning will be better / less error
		// prone? For now lets be strict.
		panic("Cannot call InstanceDestroy on an object more than once.")
		return
	}

	// Run user-destroy code
	inst.Destroy()

	// Mark as destroyed
	MarkAsDestroyed(baseObj)

	// NOTE(Jake): 2018-10-07
	// Remove at the end of the frame (gState.update)
	gState.instancesMarkedForDelete = append(gState.instancesMarkedForDelete, inst)
}

func (manager *roomInstanceManager) update(animationUpdate bool) {
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

func (manager *roomInstanceManager) draw() {
	for _, inst := range manager.instances {
		if inst == nil {
			continue
		}
		inst.Draw()
	}
}
