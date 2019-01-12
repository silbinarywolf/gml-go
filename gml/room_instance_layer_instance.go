package gml

import "sort"

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

	for _, instanceIndex := range layer.instances {
		if inst := InstanceGet(instanceIndex); inst != nil {
			inst.Draw()
		}
	}
}
