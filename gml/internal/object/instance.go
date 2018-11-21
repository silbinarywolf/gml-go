package object

type instanceObject struct {
	isDestroyed        bool
	index              int // index in the 'entities' array
	roomInstanceIndex  int // Room Instance Index belongs to
	layerInstanceIndex int // Layer belongs to
}

func (inst *Object) RoomInstanceIndex() int {
	return inst.roomInstanceIndex
}

func IsDestroyed(inst *Object) bool {
	return inst.isDestroyed
}

func MarkAsDestroyed(inst *Object) {
	inst.isDestroyed = true
}

func SetInstanceIndex(inst *Object, index int) {
	inst.index = index
}

func InstanceIndex(inst *Object) int {
	return inst.index
}

func LayerInstanceIndex(inst *Object) int {
	return inst.layerInstanceIndex
}
