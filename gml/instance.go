package gml

import (
	"github.com/silbinarywolf/gml-go/gml/internal/geom"
	"github.com/silbinarywolf/gml-go/gml/internal/object"
	"github.com/silbinarywolf/gml-go/gml/internal/space"
)

/*type InstanceIndex struct {
	layerIndex        int
	roomInstanceIndex int
	instanceIndex     int
	obj               object.ObjectType
}
*/
type instanceManagerResettableData struct {
	instances []object.ObjectType
	spaces    space.SpaceBucketArray
}

func (manager *instanceManager) reset() {
	manager.instanceManagerResettableData = instanceManagerResettableData{}
}

type instanceManager struct {
	instanceManagerResettableData
}

func newInstanceManager() *instanceManager {
	manager := new(instanceManager)
	manager.reset()
	return manager
}

func instanceCreateLayer(position geom.Vec, layer *RoomInstanceLayerInstance, roomInst *RoomInstance, objectIndex object.ObjectIndex) object.ObjectType {
	return layer.manager.InstanceCreate(position, objectIndex, roomInst.Index(), layer.index)
	/*result := InstanceIndex{
		layerIndex:        layer.index,
		roomInstanceIndex: roomInst.Index(),
		instanceIndex:     len(layer.manager.instances),
	}
	result.obj = layer.manager.InstanceCreate(position, objectIndex, roomInst.Index(), layer.index)
	return result.obj*/
}

func InstanceCreateRoom(position geom.Vec, roomInst *RoomInstance, objectIndex object.ObjectIndex) object.ObjectType {
	// NOTE(Jake): 2018-07-22
	//
	// For now instances default to the last instance layer
	//
	layerIndex := len(roomInst.instanceLayers) - 1
	//fmt.Printf("InstanceCreateRoom: Create on layer %d\n", layerIndex)
	layer := &roomInst.instanceLayers[layerIndex]
	return layer.manager.InstanceCreate(position, objectIndex, roomInst.Index(), layer.index)
}

func InstanceExists(inst object.ObjectType) bool {
	baseObj := inst.BaseObject()
	if baseObj == nil {
		return false
	}
	roomInst := RoomGetInstance(object.RoomInstanceIndex(baseObj))
	// todo(Jake): 2018-08-20
	//
	// Check to see if current entity is destroyed
	//
	return roomInst != nil
}

func (manager *instanceManager) InstanceCreate(position geom.Vec, objectIndex object.ObjectIndex, roomInstanceIndex, layerIndex int) object.ObjectType {
	// Create and add to entity list
	index := len(manager.instances)

	// Get Pos/Size part of instance (SpaceObject)
	spaceIndex := manager.spaces.GetNew()
	space := manager.spaces.Get(spaceIndex)

	// Get instance
	inst := object.NewRawInstance(objectIndex, index, roomInstanceIndex, layerIndex, space, spaceIndex)
	manager.instances = append(manager.instances, inst)

	// Init and Set position
	inst.Create()
	inst.BaseObject().Vec = position
	return inst
}

func InstanceDestroy(inst object.ObjectType) {
	// Destroy this
	inst.Destroy()
	cameraInstanceDestroy(inst)

	baseObj := inst.BaseObject()

	// Get slots
	roomInstanceIndex := object.RoomInstanceIndex(baseObj)
	layerIndex := object.LayerInstanceIndex(baseObj)
	index := object.InstanceIndex(baseObj)

	// Get manager
	roomInst := &gState.roomInstances[roomInstanceIndex]
	layerInst := &roomInst.instanceLayers[layerIndex]
	manager := &layerInst.manager

	// Free up SpaceObject slot
	spaceIndex := baseObj.SpaceIndex()
	baseObj.Space = nil
	if spaceIndex > -1 {
		manager.spaces.Remove(spaceIndex)
	}

	// Unordered delete
	lastEntry := manager.instances[len(manager.instances)-1]
	manager.instances[index] = lastEntry
	manager.instances = manager.instances[:len(manager.instances)-1]
}

func (manager *instanceManager) update(animationUpdate bool) {
	{
		instances := manager.instances
		for _, inst := range instances {
			inst.Update()
		}

		if animationUpdate {
			for _, inst := range instances {
				if inst == nil {
					continue
				}
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
