package gml

import (
	"reflect"

	"github.com/silbinarywolf/gml-go/gml/internal/assert"
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

// InstanceRestore re-creates an object using a previously used instance index
// and object index. This is used to bring old objects back with serialization.
func InstanceRestore(oldInstanceIndex InstanceIndex, objectIndex ObjectIndex) ObjectType {
	if inst := oldInstanceIndex.Get(); inst != nil {
		panic("Cannot call InstanceRestore if instance still exists.")
	}
	inst, slot := allocateNewInstance(objectIndex)
	baseObj := inst.BaseObject()
	baseObj.internal.InstanceIndex = oldInstanceIndex
	gState.instanceManager.instanceIndexToIndex[baseObj.internal.InstanceIndex] = slot
	//assert.DebugAssert(baseObj.internal.RoomInstanceIndex == 0, "Room Instance Index cannot be 0")
	return inst
}

func allocateNewInstance(objectIndex ObjectIndex) (ObjectType, int) {
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
	inst.BaseObject().internal.ObjectIndex = objectIndex
	return inst, slot
}

func (inst *Object) InstanceIndex() InstanceIndex {
	return inst.internal.InstanceIndex
}

func (inst *Object) RoomInstanceIndex() RoomInstanceIndex {
	return inst.internal.RoomInstanceIndex
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

func instanceCreate(x, y float64, objectIndex ObjectIndex, callback func(inst *Object), assignNewInstanceIndex bool) ObjectType {
	inst, slot := allocateNewInstance(objectIndex)
	{
		baseObj := inst.BaseObject()
		if assignNewInstanceIndex {
			// Get next instance index
			gState.instanceManager.nextInstanceIndex++
			baseObj.internal.InstanceIndex = gState.instanceManager.nextInstanceIndex
		}
		baseObj.Vec = geom.Vec{x, y}

		callback(baseObj)

		assert.DebugAssert(baseObj.internal.RoomInstanceIndex == 0, "Instance index cannot be 0")
		gState.instanceManager.instanceIndexToIndex[baseObj.internal.InstanceIndex] = slot
		assert.DebugAssert(baseObj.internal.RoomInstanceIndex == 0, "Room Instance Index cannot be 0")
		if assignNewInstanceIndex {
			baseObj.create()
			inst.Create()
		}
	}

	return inst
}

// InstanceExists will return true if an object has not been destroyed and belongs to a room
func InstanceExists(inst ObjectType) bool {
	baseObj := inst.BaseObject()
	return baseObj != nil &&
		!baseObj.internal.IsDestroyed
}

// fastInstanceDestroy exists to quickly destroy instances without removing
// from an array. Used by rooms when they're destroying themselves
func fastInstanceDestroy(inst ObjectType) {
	// Run user-destroy code
	inst.Destroy()

	// Mark as destroyed
	inst.BaseObject().internal.IsDestroyed = true
}

func InstanceDestroy(inst ObjectType) {
	baseObj := inst.BaseObject()
	if baseObj.internal.IsDestroyed {
		// NOTE(Jake): 2018-10-07
		// Maybe making this just silently returning will be better / less error
		// prone? For now lets be strict.
		panic("Cannot call InstanceDestroy on an object more than once.")
	}

	fastInstanceDestroy(inst)

	// NOTE(Jake): 2018-10-07
	// Remove at the end of the frame (gState.update)
	gState.instancesMarkedForDelete = append(gState.instancesMarkedForDelete, baseObj.InstanceIndex())
}
