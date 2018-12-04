package gml

//type RoomInstanceLayer interface {
//	update(animationUpdate bool)
//	draw()
//}

type roomInstanceLayerDrawBase struct {
	drawOrder int32
}

func (layer *roomInstanceLayerDrawBase) order() int32 {
	return layer.drawOrder
}

type roomInstanceLayerUpdate interface {
	update(animationUpdate bool)
}

type roomInstanceLayerDraw interface {
	draw()
	order() int32
}
