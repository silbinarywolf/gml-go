package gml

import (
	"sort"
)

type roomInstanceLayerInstance struct {
	roomInstanceLayerDrawBase
	index     int
	name      string
	instances []InstanceIndex
	//manager roomInstanceManager
	//_parent *RoomInstance
}

//func (layer *RoomInstanceLayerInstance) parent() *RoomInstance {
//	return layer._parent
//}

//func (layer *roomInstanceLayerInstance) update(animationUpdate bool) {
//	for _, room := range layer.instances {
//
//	}
//}

func (layer *roomInstanceLayerInstance) draw() {
	// Sort by order
	sort.SliceStable(layer.instances, func(i, j int) bool {
		a := InstanceGet(layer.instances[i])
		if a == nil {
			return false
		}
		b := InstanceGet(layer.instances[j])
		if b == nil {
			return false
		}
		return a.BaseObject().Depth() > b.BaseObject().Depth()
	})
	//log.Printf("Stable sort count: %d, cap: %d\n", len(layer.instances), cap(layer.instances))

	for _, instanceIndex := range layer.instances {
		inst := InstanceGet(instanceIndex)
		if inst == nil {
			panic("instance index not removed from draw list when destroyed")
		}
		inst.Draw()
	}
}
