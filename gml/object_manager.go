package gml

import (
	"reflect"
)

var (
	gObjectManager *objectManager = newObjectManager()
)

type objectManager struct {
	idToEntityData []ObjectType
	//objectIndexList []ObjectIndex
	indexToName []string
	nameToID    map[string]ObjectIndex
}

func newObjectManager() *objectManager {
	return &objectManager{
		idToEntityData: nil, // NOTE: This should be initialized in user-code with gml.ObjectInitTypes()
		nameToID:       make(map[string]ObjectIndex),
	}
}

// InitObjectGeneratedData is required to be called so the engine can create game objects
func InitObjectGeneratedData(indexToName []string, nameToIndex map[string]ObjectIndex, objTypes []ObjectType) {
	manager := gObjectManager
	if manager.idToEntityData != nil {
		panic("Cannot call init type function more than once.")
	}
	manager.indexToName = indexToName
	manager.nameToID = nameToIndex
	manager.idToEntityData = objTypes
	debugInitObjectMetaList(objTypes[1:])
}

//func ObjectIndexList() []ObjectIndex {
//	return gObjectManager.objectIndexList
//}

//
// This is used to get an object index by the object name.
//

func moveInstance(inst ObjectType, index int, roomInstanceIndex RoomInstanceIndex, layerIndex int) {
	// Initialize object
	baseObj := inst.BaseObject()
	baseObj.index = index
	baseObj.roomInstanceIndex = roomInstanceIndex
	baseObj.layerInstanceIndex = layerIndex
}

func newRawInstance(objectIndex ObjectIndex, index int, roomInstanceIndex RoomInstanceIndex, layerIndex int) ObjectType {
	valToCopy := gObjectManager.idToEntityData[objectIndex]
	if valToCopy == nil {
		panic("Invalid objectIndex given")
	}
	inst := reflect.New(reflect.ValueOf(valToCopy).Elem().Type()).Interface().(ObjectType)
	moveInstance(inst, index, roomInstanceIndex, layerIndex)
	baseObj := inst.BaseObject()
	baseObj.objectIndex = objectIndex
	baseObj.create()
	return inst
	/*// Create
	valToCopy := gObjectManager.idToEntityData[objectIndex]
	inst := reflect.New(reflect.ValueOf(valToCopy).Elem().Type()).Interface().(ObjectType)

	// Initialize object
	baseObj := inst.BaseObject()
	baseObj.index = index
	baseObj.roomInstanceIndex = roomInstanceIndex
	baseObj.layerInstanceIndex = layerIndex
	// todo(Jake): 2018-07-08
	//
	// Figure out a cleaner way to handle this functionality across
	// the room editor and gamecode.
	//
	// Perhaps force objects to have to be created via an instance manager.
	//
	baseObj.SpaceObject.Init(space, spaceIndex)
	baseObj.create()*/
}

func ObjectGetIndex(name string) (ObjectIndex, bool) {
	res, ok := gObjectManager.nameToID[name]
	return res, ok
}
