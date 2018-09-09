package gml

type RoomInstanceLayerInstance struct {
	RoomInstanceLayerDrawBase
	index   int
	name    string
	manager instanceManager
	//_parent *RoomInstance
}

//func (layer *RoomInstanceLayerInstance) parent() *RoomInstance {
//	return layer._parent
//}

func (layer *RoomInstanceLayerInstance) update(animationUpdate bool) {
	layer.manager.update(animationUpdate)
}

func (layer *RoomInstanceLayerInstance) draw() {
	layer.manager.draw()
}
