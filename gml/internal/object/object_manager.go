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
	return &objectManager{
		idToEntityData: nil, // NOTE: This should be initialized in user-code with gml.ObjectInitTypes()
		nameToID:       make(map[string]ObjectIndex),
	}
}

func InitTypes(objTypes []ObjectType) {
	manager := gObjectManager
	if manager.idToEntityData != nil {
		panic("Cannot call init type function more than once.")
	}
	manager.idToEntityData = objTypes
	for _, objType := range objTypes {
		if objType == nil {
			continue
		}
		name := objType.ObjectName()
		objectIndex := objType.ObjectIndex()
		manager.nameToID[name] = objectIndex
	}
}

func IDToEntityData() []ObjectType {
	return gObjectManager.idToEntityData
}

//
// This is used to get an object index by the object name.
//
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
