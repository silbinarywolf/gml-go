package object

import "reflect"

var (
	gObjectManager *objectManager = newObjectManager()
)

type objectManager struct {
	idToEntityData []ObjectType
	nameToID       map[string]ObjectIndex
}

func newObjectManager() *objectManager {
	return &objectManager{}
}

func Init(idToEntityData []ObjectType, nameToID map[string]ObjectIndex) {
	gObjectManager.idToEntityData = idToEntityData
	gObjectManager.nameToID = nameToID
}

func IDToEntityData() []ObjectType {
	return gObjectManager.idToEntityData
}

func NameToID() map[string]ObjectIndex {
	return gObjectManager.nameToID
}

func NewRawInstance(objectIndex ObjectIndex, index int, roomInstanceIndex int) ObjectType {
	// Create
	valToCopy := gObjectManager.idToEntityData[objectIndex]
	inst := reflect.New(reflect.ValueOf(valToCopy).Elem().Type()).Interface().(ObjectType)

	// Initialize object
	baseObj := inst.BaseObject()
	baseObj.index = index
	baseObj.roomInstanceIndex = roomInstanceIndex
	baseObj.Create()

	return inst
}
