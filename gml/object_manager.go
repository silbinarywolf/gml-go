package gml

var (
	gObjectManager *objectManager = newObjectManager()
)

type objectManager struct {
	idToEntityData []ObjectType
	indexToName    []string
	nameToID       map[string]ObjectIndex
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
	manager.idToEntityData = objTypes
	manager.indexToName = indexToName
	manager.nameToID = nameToIndex
	if len(objTypes) > 0 {
		debugInitObjectMetaList(objTypes[1:])
	}
}

// UnsafeObjectTypeList provides a copy of the list of object type definition data
// this is to be used by custom tools like a room editor or similar.
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

//
// This is used to get an object index by the object name.
//

/*func moveInstance(inst ObjectType, roomInstanceIndex RoomInstanceIndex, layerIndex int) {
	// Initialize object
	baseObj := inst.BaseObject()
	baseObj.index = index
	baseObj.roomInstanceIndex = roomInstanceIndex
	baseObj.layerInstanceIndex = layerIndex
}*/

/*func newRawInstance(objectIndex ObjectIndex, index int, roomInstanceIndex RoomInstanceIndex) ObjectType {
	inst := objectIndex.New()
	baseObj := inst.BaseObject()
	baseObj.internal.RoomInstanceIndex = roomInstanceIndex
	return inst
}*/

func ObjectGetIndex(name string) (ObjectIndex, bool) {
	res, ok := gObjectManager.nameToID[name]
	return res, ok
}
