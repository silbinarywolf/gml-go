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

func InstanceCreate(position Vec, entityID ObjectIndex) ObjectType {
	if entityID == 0 {
		panic("Cannot pass 0 as 2nd parameter to InstanceCreate(position, entityID)")
	}
	valToCopy := gInstanceManager.idToEntityData[entityID]

	// Create and add to entity list
	inst := reflect.New(reflect.ValueOf(valToCopy).Elem().Type()).Interface().(ObjectType)
	index := len(gInstanceManager.entities)
	gInstanceManager.entities = append(gInstanceManager.entities, inst)

	// Initialize object
	baseObj := inst.BaseObject()
	baseObj.index = index
	baseObj.Create()
	inst.Create()

	// Set at instance create position
	baseObj.Vec = position
	return inst
}

func InstanceDestroy(entity ObjectType) {
	be := entity.BaseObject()

	// Unordered delete
	i := be.index
	manager := gInstanceManager
	lastEntry := manager.entities[len(manager.entities)-1]
	manager.entities[i] = lastEntry
	manager.entities = manager.entities[:len(manager.entities)-1]

	// maybetodo(Jake): 2018-05-27
	//
	// Add func Destroy() to Entity interface and Call e.Destroy()
	//
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
	for _, inst := range manager.entities {
		if inst == nil {
			continue
		}
		inst.Draw()
	}
}
