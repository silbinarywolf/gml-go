package gml

import "reflect"

type instanceManagerResettableData struct {
	entities []ObjectType
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
	idToEntityData []ObjectType
	nameToID       map[string]ObjectIndex
}

var (
	gInstanceManager *instanceManager = newInstanceManager()
)

func newInstanceManager() *instanceManager {
	manager := new(instanceManager)
	manager.reset()
	return manager
}

func ObjectGetIndex(name string) (ObjectIndex, bool) {
	res, ok := gInstanceManager.nameToID[name]
	return res, ok
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

func (manager *instanceManager) InstanceDestroy(inst ObjectType) {
	be := inst.BaseObject()

	// Unordered delete
	i := be.index
	lastEntry := manager.entities[len(manager.entities)-1]
	manager.entities[i] = lastEntry
	manager.entities = manager.entities[:len(manager.entities)-1]
}

// NOTE(Jake): 2018-06-02
//
// Kinda hacky way to get width of instances to calculate room bounds
//
func newInstance(entityID ObjectIndex) ObjectType {
	// Create
	valToCopy := gInstanceManager.idToEntityData[entityID]
	inst := reflect.New(reflect.ValueOf(valToCopy).Elem().Type()).Interface().(ObjectType)

	// Initialize object
	baseObj := inst.BaseObject()
	baseObj.Create()

	return inst
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
		baseObj.SpriteState.imageUpdate()
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
