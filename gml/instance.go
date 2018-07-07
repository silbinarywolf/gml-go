package gml

import "github.com/silbinarywolf/gml-go/gml/internal/object"

type instanceManagerResettableData struct {
	instances      []object.ObjectType
	instanceSpaces []object.Space
}

func (manager *instanceManager) reset() {
	manager.instanceManagerResettableData = instanceManagerResettableData{}
}

type instanceManager struct {
	instanceManagerResettableData

	// todo(Jake): 2018-06-02
	//
	// Move these to an objectManager struct
	//
	//idToEntityData []ObjectType
	//nameToID       map[string]ObjectIndex
}

func newInstanceManager() *instanceManager {
	manager := new(instanceManager)
	manager.reset()
	return manager
}

func (manager *instanceManager) InstanceCreate(position Vec, objectIndex object.ObjectIndex, roomInstanceIndex int) object.ObjectType {
	// Create and add to entity list
	index := len(manager.instances)
	inst := object.NewRawInstance(objectIndex, index, roomInstanceIndex)

	// NOTE(Jake): 2018-07-07
	//
	// These needs to be in-sync
	//
	manager.instanceSpaces = append(manager.instanceSpaces, object.Space{})
	manager.instances = append(manager.instances, inst)

	// Attach
	baseObj := inst.BaseObject()
	baseObj.Space = &manager.instanceSpaces[index]

	// Init and Set position
	inst.Create()
	baseObj.Vec = position
	return inst
}

func (manager *instanceManager) InstanceDestroy(inst object.ObjectType) {
	be := inst.BaseObject()

	// Unordered delete
	i := be.Index()
	lastEntry := manager.instances[len(manager.instances)-1]
	manager.instances[i] = lastEntry
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
	for i := 0; i < len(cameraList); i++ {
		cam := &cameraList[i]
		if !cam.enabled {
			continue
		}
		cam.update()
		cameraSetActive(i)
		for _, inst := range manager.instances {
			if inst == nil {
				continue
			}
			inst.Draw()
		}
	}
	cameraClearActive()
}
