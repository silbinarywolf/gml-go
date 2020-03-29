package gml

var gObjectManager = &objectManager{
	idToEntityData: make(map[ObjectIndex]ObjectType),
}

type objectManager struct {
	idToEntityData map[ObjectIndex]ObjectType
}

// UnsafeAddObjectGeneratedData is used by code generation to tie the object index to the data
func UnsafeAddObjectGeneratedData(objectIndex ObjectIndex, data ObjectType) {
	manager := gObjectManager
	manager.idToEntityData[objectIndex] = data
}

// UnsafeObjectTypeList provides a copy of the list of object type definition data
// this is to be used by custom tools like a room editor or similar.
// Backwards compatibility usage with this function is not guaranteed.
func UnsafeObjectTypeList() []ObjectType {
	if len(gObjectManager.idToEntityData) == 0 {
		panic("UnsafeObjectTypeList is not initialized yet")
	}
	r := make([]ObjectType, 0, len(gObjectManager.idToEntityData))
	for id := 1; id < len(gObjectManager.idToEntityData); id++ {
		inst := ObjectIndex(id).new()
		r = append(r, inst)
	}
	return r
}
