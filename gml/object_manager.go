package gml

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

func ObjectGetIndex(name string) (ObjectIndex, bool) {
	res, ok := gObjectManager.nameToID[name]
	return res, ok
}
