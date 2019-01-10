package game

import (
	"github.com/silbinarywolf/gml-go/gml"
)

type Physics struct {
	Speed   gml.Vec
	Gravity float64
}

func (phys *Physics) Update(self *gml.Object) {
	phys.Speed.Y += phys.Gravity
	self.Y += phys.Speed.Y
}
