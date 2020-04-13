package gml

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
	inst := allocateRawInstance(objectIndex)
	manager := &gState.instanceManager
	manager.instances = append(manager.instances, inst)
	slot := len(manager.instances) - 1
	return inst, slot
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

// InstanceExists will return true if an object has not been destroyed and belongs to a room
func InstanceExists(inst ObjectType) bool {
	baseObj := inst.BaseObject()
	return baseObj != nil &&
		!baseObj.internal.IsDestroyed
}

// fastInstanceDestroy exists to quickly destroy instances without removing
// from an array. Used by rooms when they're destroying themselves
func fastInstanceFree(inst ObjectType) {
	// Run free code
	inst.Free()

	// Mark as freed / destroyed
	inst.BaseObject().internal.IsDestroyed = true
}

// InstanceDestroy will make an object call its Destroy then its Free method and
// then it will no longer exist.
func InstanceDestroy(inst ObjectType) {
	baseObj := inst.BaseObject()
	if baseObj.internal.IsDestroyed {
		// NOTE(Jake): 2018-10-07
		// Maybe making this just silently returning will be better / less error
		// prone? For now lets be strict.
		panic("Cannot call InstanceDestroy on an object more than once.")
	}

	// Run user-destroy code
	inst.Destroy()

	fastInstanceFree(inst)

	// NOTE(Jake): 2018-10-07
	// Remove at the end of the frame (gState.update)
	gState.instancesMarkedForDelete = append(gState.instancesMarkedForDelete, baseObj.InstanceIndex())
}

// InstanceFree will make an object call its Free method and
// then it will no longer exist.
func InstanceFree(inst ObjectType) {
	baseObj := inst.BaseObject()
	if baseObj.internal.IsDestroyed {
		// NOTE(Jake): 2019-09-12
		// Maybe making this just silently returning will be better / less error
		// prone? For now lets be strict.
		panic("Cannot call InstanceFree on an object more than once.")
	}

	fastInstanceFree(inst)

	// NOTE(Jake): 2019-09-12
	// Remove at the end of the frame (gState.update)
	gState.instancesMarkedForDelete = append(gState.instancesMarkedForDelete, baseObj.InstanceIndex())
}
