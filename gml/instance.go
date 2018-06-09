package gml

import "github.com/silbinarywolf/gml-go/gml/internal/object"

type instanceManagerResettableData struct {
	entities []object.ObjectType
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

var (
	gInstanceManager *instanceManager = newInstanceManager()
)

func newInstanceManager() *instanceManager {
	manager := new(instanceManager)
	manager.reset()
	return manager
}

//func GetAll() []ObjectType {
//	return gInstanceManager.entities
//}

//func InstanceCreate(position Vec, entityID ObjectIndex) ObjectType {
//	return gInstanceManager.InstanceCreate(position, entityID)
//}

//func InstanceDestroy(entity ObjectType) {
//	gInstanceManager.InstanceDestroy(entity)
//}

/*func (manager *instanceManager) InstanceCreate(position Vec, objectIndex ObjectIndex) ObjectType {
	// Create and add to entity list
	inst := newInstance(objectIndex)
	baseObj := inst.BaseObject()
	baseObj.index = len(manager.entities)
	manager.entities = append(manager.entities, inst)

	// Init and set position
	inst.Create()
	baseObj.Vec = position
	return inst
}*/

func (manager *instanceManager) InstanceDestroy(inst object.ObjectType) {
	be := inst.BaseObject()

	// Unordered delete
	i := be.Index()
	lastEntry := manager.entities[len(manager.entities)-1]
	manager.entities[i] = lastEntry
	manager.entities = manager.entities[:len(manager.entities)-1]
}

func (manager *instanceManager) update() {
	entities := manager.entities
	for _, inst := range entities {
		inst.Update()
	}

	for _, inst := range entities {
		if inst == nil {
			continue
		}
		baseObj := inst.BaseObject()
		baseObj.SpriteState.ImageUpdate()
	}
}

func (manager *instanceManager) draw() {
	for i := 0; i < len(cameraList); i++ {
		currentCamera = &cameraList[i]
		if !currentCamera.enabled {
			continue
		}
		currentCamera.update()
		for _, inst := range manager.entities {
			if inst == nil {
				continue
			}
			inst.Draw()
		}
	}
	currentCamera = nil
}
