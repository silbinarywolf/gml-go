package gml

import (
	"reflect"
	"sort"
	"strconv"
)

var gObjectManager = &objectManager{
	idToEntityData: make(map[ObjectIndex]ObjectType),
}

type objectManager struct {
	idToEntityData map[ObjectIndex]ObjectType
}

// UnsafeAddObjectGeneratedData is used by code generation to tie the object index to the data
func UnsafeAddObjectGeneratedData(objectIndex ObjectIndex, data ObjectType) {
	gObjectManager.idToEntityData[objectIndex] = data
}

// UnsafeObjectTypeList provides a copy of the list of object type definition data
// this is to be used by custom tools like a room editor or similar.
// Backwards compatibility usage with this function is not guaranteed.
func UnsafeObjectTypeList() []ObjectType {
	if len(gObjectManager.idToEntityData) == 0 {
		panic("UnsafeObjectTypeList is not initialized yet")
	}
	r := make([]ObjectType, 0, len(gObjectManager.idToEntityData))
	for id, _ := range gObjectManager.idToEntityData {
		inst := allocateRawInstance(id)
		r = append(r, inst)
	}
	// Sort alphabetically
	sort.Slice(r[:], func(i, j int) bool {
		return r[i].ObjectName() < r[j].ObjectName()
	})
	return r
}

// allocateRawInstance is used by allocateNewInstance and UnsafeObjectTypeList
func allocateRawInstance(objectIndex ObjectIndex) ObjectType {
	var inst ObjectType
	valToCopy := gObjectManager.idToEntityData[objectIndex]
	if valToCopy == nil {
		panic("Invalid objectIndex given: " + strconv.Itoa(int(objectIndex)))
	}
	inst = reflect.New(reflect.ValueOf(valToCopy).Elem().Type()).Interface().(ObjectType)
	baseObj := inst.BaseObject()
	baseObj.reset()
	baseObj.internal.ObjectIndex = objectIndex
	return inst
}
