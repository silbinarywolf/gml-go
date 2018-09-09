package object

type instanceObject struct {
	index              int // index in the 'entities' array
	roomInstanceIndex  int // Room Instance Index belongs to
	layerInstanceIndex int // Layer belongs to
}

func InstanceIndex(inst *Object) int {
	return inst.index
}

func RoomInstanceIndex(inst *Object) int {
	return inst.roomInstanceIndex
}

func LayerInstanceIndex(inst *Object) int {
	return inst.layerInstanceIndex
}
