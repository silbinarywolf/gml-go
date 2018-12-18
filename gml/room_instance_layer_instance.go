package gml

type roomInstanceLayerInstance struct {
	roomInstanceLayerDrawBase
	index   int
	name    string
	manager roomInstanceManager
	//_parent *RoomInstance
}

//func (layer *RoomInstanceLayerInstance) parent() *RoomInstance {
//	return layer._parent
//}

func (layer *roomInstanceLayerInstance) update(animationUpdate bool) {
	layer.manager.update(animationUpdate)
}

func (layer *roomInstanceLayerInstance) draw() {
	layer.manager.draw()
}
