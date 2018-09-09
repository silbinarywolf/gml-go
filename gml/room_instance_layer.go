package gml

//type RoomInstanceLayer interface {
//	update(animationUpdate bool)
//	draw()
//}

type RoomInstanceLayerDrawBase struct {
	drawOrder int32
}

func (layer *RoomInstanceLayerDrawBase) order() int32 {
	return layer.drawOrder
}

type RoomInstanceLayerUpdate interface {
	update(animationUpdate bool)
}

type RoomInstanceLayerDraw interface {
	draw()
	order() int32
}
