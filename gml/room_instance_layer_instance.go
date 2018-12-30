package gml

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
	for _, instanceIndex := range layer.instances {
		if inst := InstanceGet(instanceIndex); inst != nil {
			inst.Draw()
		}
	}
}
