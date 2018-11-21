package object

import (
	"fmt"
	"reflect"
)

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
		otherID, used := manager.nameToID[name]
		if used {
			otherType := manager.idToEntityData[otherID]
			panic(fmt.Sprintf("You cannot have two objects with the same object name.\n- %T::ObjectName() == %s\n- %T::ObjectName() == %s", objType, objType.ObjectName(), otherType, otherType.ObjectName()))
		}
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

func MoveInstance(inst ObjectType, index int, roomInstanceIndex int, layerIndex int) {
	// Initialize object
	baseObj := inst.BaseObject()
	baseObj.index = index
	baseObj.roomInstanceIndex = roomInstanceIndex
	baseObj.layerInstanceIndex = layerIndex
}

func NewRawInstance(objectIndex ObjectIndex, index int, roomInstanceIndex int, layerIndex int) ObjectType {
	valToCopy := gObjectManager.idToEntityData[objectIndex]
	inst := reflect.New(reflect.ValueOf(valToCopy).Elem().Type()).Interface().(ObjectType)
	MoveInstance(inst, index, roomInstanceIndex, layerIndex)
	baseObj := inst.BaseObject()
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
	nameToID := NameToID()
	res, ok := nameToID[name]
	return res, ok
}
