package gml

import (
	"reflect"

	"github.com/silbinarywolf/gml-go/gml/internal/geom"
)

// Noone is to be used when checking if there is no instance with InstanceIndex type
const Noone InstanceIndex = 0

type InstanceIndex int32

type instanceManager struct {
	instances            []ObjectType
	instanceIndexToIndex map[InstanceIndex]int
	nextInstanceIndex    InstanceIndex
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
	{
		inst := reflect.New(reflect.ValueOf(valToCopy).Elem().Type()).Interface().(ObjectType)
		manager.instances = append(manager.instances, inst)
	}
	slot := len(manager.instances) - 1
	inst := manager.instances[slot]
	baseObj := inst.BaseObject()
	gState.instanceManager.nextInstanceIndex++
	baseObj.instanceIndex = gState.instanceManager.nextInstanceIndex
	manager.instanceIndexToIndex[baseObj.instanceIndex] = slot
	return inst
}

// todo: Jake: 2018-12-16
// Deprecate this in favour of one storage area for all entities
type roomInstanceManager struct {
	instances []ObjectType
}

type instanceObject struct {
	isDestroyed        bool
	instanceIndex      InstanceIndex     // global uuid
	roomInstanceIndex  RoomInstanceIndex // Room Instance Index belongs to
	layerInstanceIndex int               // Layer belongs to
}

func (inst *Object) InstanceIndex() InstanceIndex {
	return inst.instanceIndex
}

func (inst *Object) RoomInstanceIndex() RoomInstanceIndex {
	return inst.roomInstanceIndex
}

func newroomInstanceManager() *roomInstanceManager {
	manager := new(roomInstanceManager)
	manager.reset()
	return manager
}

func (manager *roomInstanceManager) reset() {
	*manager = roomInstanceManager{}
}

// getBaseObject get the base object for an instance
func (index InstanceIndex) getBaseObject() *Object {
	if inst := index.Get(); inst != nil {
		return inst.BaseObject()
	}
	return nil
}

func (index InstanceIndex) Get() ObjectType {
	dataIndex, ok := gState.instanceManager.instanceIndexToIndex[index]
	if !ok {
		return nil
	}
	inst := gState.instanceManager.instances[dataIndex]
	if !InstanceExists(inst) {
		return nil
	}
	return inst
}

func InstanceChangeRoom(inst ObjectType, roomInstanceIndex RoomInstanceIndex) {
	roomInst := &gState.roomInstances[roomInstanceIndex]
	if !roomInst.used {
		return
	}
	// NOTE(Jake): 2018-07-22
	// For now instances default to the last instance layer
	//layerIndex := len(roomInst.instanceLayers) - 1
	//layer := &roomInst.instanceLayers[layerIndex]

	//instanceRemove(inst)
	panic("todo: Update this to remove instance index from one room instance list and add it to another")
	// Move entity to new list
	//index := len(manager.instances)
	//moveInstance(inst, index, roomInstanceIndex, layerIndex)
	//manager.instances = append(manager.instances, inst)
}

func InstanceCreate(x, y float64, roomInstanceIndex RoomInstanceIndex, objectIndex ObjectIndex) ObjectType {
	inst := allocateNewInstance(objectIndex)
	{
		baseObj := inst.BaseObject()
		baseObj.Vec = geom.Vec{x, y}
		baseObj.objectIndex = objectIndex
		baseObj.roomInstanceIndex = roomInstanceIndex
		roomInst := &gState.roomInstances[roomInstanceIndex]
		// NOTE(Jake): 2018-07-22
		// For now instances default to the last instance layer
		layerIndex := len(roomInst.instanceLayers) - 1
		layer := &roomInst.instanceLayers[layerIndex]
		layer.instances = append(layer.instances, baseObj.InstanceIndex())

		baseObj.create()
		inst.Create()
	}

	return inst
	/*roomInst := &gState.roomInstances[roomInstanceIndex]
	// NOTE(Jake): 2018-07-22
	// For now instances default to the last instance layer
	layerIndex := len(roomInst.instanceLayers) - 1
	layer := &roomInst.instanceLayers[layerIndex]
	//manager.instances = append(manager.instances, inst)
	//fmt.Printf("InstanceCreateRoom: Create on layer %d\n", layerIndex)
	return layer.manager.InstanceCreate(position, objectIndex, roomInst.index, layer.index)*/
}

// InstanceExists will return true if an object has not been destroyed and belongs to a room
func InstanceExists(inst ObjectType) bool {
	baseObj := inst.BaseObject()
	roomInst := roomGetInstance(baseObj.RoomInstanceIndex())
	return baseObj != nil &&
		!baseObj.isDestroyed &&
		roomInst != nil
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

// WithAll returns a list of instances in the same room as the provided object
func WithAll(instType collisionObject) []InstanceIndex {
	inst := instType.BaseObject()
	room := roomGetInstance(inst.BaseObject().RoomInstanceIndex())
	if room == nil {
		panic("RoomInstance this object belongs to has been destroyed")
	}
	var list []InstanceIndex
	for i := 0; i < len(room.instanceLayers); i++ {
		for _, otherIndex := range room.instanceLayers[i].instances {
			other := otherIndex.getBaseObject()
			if other == nil {
				continue
			}
			list = append(list, otherIndex)
		}
	}
	if len(list) == 0 {
		return nil
	}
	return list
}

/*func WithObject(instType collisionObject, objectIndex ObjectIndex) []InstanceIndex {
	inst := instType.BaseObject()
	room := roomGetInstance(inst.BaseObject().RoomInstanceIndex())
	if room == nil {
		panic("RoomInstance this object belongs to has been destroyed")
	}
	var list []InstanceIndex
	for i := 0; i < len(room.instanceLayers); i++ {
		for _, otherIndex := range room.instanceLayers[i].instances {
			other := otherIndex.getBaseObject()
			if other == nil ||
				other.ObjectIndex() == objectIndex {
				continue
			}
			list = append(list, otherIndex)
		}
	}
	if len(list) == 0 {
		return nil
	}
	return list
}*/

/*
func instanceRemove(inst ObjectType) {
	baseObj := inst.BaseObject()

	// Get slots
	roomInstanceIndex := baseObj.roomInstanceIndex
	layerIndex := baseObj.layerInstanceIndex
	index := baseObj.index

	// Get manager
	roomInst := &gState.roomInstances[roomInstanceIndex]
	layerInst := &roomInst.instanceLayers[layerIndex]

	if layerInst.instances[index] != inst {
		panic("instanceRemove failed as instance provided has already been removed")
	}
	// Get index
	index := -1
	for i, otherInst := range manager.instances {
		if inst == otherInst {
			index = i
		}
	}
	if index == -1 {
		panic("instanceRemove failed as instance provided has already been removed")
	}

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
}*/

func InstanceDestroy(inst ObjectType) {
	baseObj := inst.BaseObject()
	if baseObj.isDestroyed {
		// NOTE(Jake): 2018-10-07
		// Maybe making this just silently returning will be better / less error
		// prone? For now lets be strict.
		panic("Cannot call InstanceDestroy on an object more than once.")
	}

	// Run user-destroy code
	inst.Destroy()

	// Mark as destroyed
	baseObj.isDestroyed = true

	// NOTE(Jake): 2018-10-07
	// Remove at the end of the frame (gState.update)
	gState.instancesMarkedForDelete = append(gState.instancesMarkedForDelete, baseObj.InstanceIndex())
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
