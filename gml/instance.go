package gml

import (
	"fmt"
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

// InstanceRestore re-creates an object using a previously used instance index
// and object index. This is used to bring old objects back with serialization.
func InstanceRestore(oldInstanceIndex InstanceIndex, objectIndex ObjectIndex) ObjectType {
	inst := oldInstanceIndex.Get()
	if inst != nil {
		panic("Cannot call InstanceRestore if instance still exists.")
	}
	return instanceCreate(0, 0, objectIndex, func(inst *Object) {
		inst.internal.instanceIndex = oldInstanceIndex
		//roomInst := &roomInstanceState.roomInstances[roomInstanceIndex]
		//roomInst.instances = append(roomInst.instances, inst.InstanceIndex())
	}, false)
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
	return inst
}

// todo: Jake: 2018-12-16
// Deprecate this in favour of one storage area for all entities
type roomInstanceManager struct {
	instances []ObjectType
}

type instanceObject struct {
	isDestroyed       bool
	instanceIndex     InstanceIndex     // global uuid
	roomInstanceIndex RoomInstanceIndex // Room Instance Index belongs to
	//layerInstanceIndex int               // Layer belongs to
}

func (inst *Object) InstanceIndex() InstanceIndex {
	return inst.internal.instanceIndex
}

func (inst *Object) RoomInstanceIndex() RoomInstanceIndex {
	return inst.internal.roomInstanceIndex
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

func instanceCreate(x, y float64, objectIndex ObjectIndex, callback func(inst *Object), assignNewInstanceIndex bool) ObjectType {
	inst := allocateNewInstance(objectIndex)
	{
		baseObj := inst.BaseObject()
		baseObj.internal.objectIndex = objectIndex
		if assignNewInstanceIndex {
			// Get next instance index
			gState.instanceManager.nextInstanceIndex++
			baseObj.internal.instanceIndex = gState.instanceManager.nextInstanceIndex
		}
		baseObj.Vec = geom.Vec{x, y}

		callback(baseObj)

		if baseObj.internal.instanceIndex == 0 {
			panic("Instance index cannot be 0")
		}

		// NOTE: Jake: 2019-03-08
		// This is sensitive as it expects InstanceRestore to set a instanceIndex
		slot := len(gState.instanceManager.instances) - 1
		fmt.Printf("Instance Index: %d\n", baseObj.internal.instanceIndex)
		gState.instanceManager.instanceIndexToIndex[baseObj.internal.instanceIndex] = slot

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
	roomInst := roomGetInstance(baseObj.RoomInstanceIndex())
	return baseObj != nil &&
		!baseObj.internal.isDestroyed &&
		roomInst != nil
}

func InstanceDestroy(inst ObjectType) {
	baseObj := inst.BaseObject()
	if baseObj.internal.isDestroyed {
		// NOTE(Jake): 2018-10-07
		// Maybe making this just silently returning will be better / less error
		// prone? For now lets be strict.
		panic("Cannot call InstanceDestroy on an object more than once.")
	}

	// Run user-destroy code
	inst.Destroy()

	// Mark as destroyed
	baseObj.internal.isDestroyed = true

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
