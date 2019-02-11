package game

import (
	"github.com/silbinarywolf/gml-go/gml"
)

type Physics struct {
	Speed   gml.Vec
	Gravity float64
}

func (phys *Physics) Update(self *gml.Object) {
	// Formula taken from:
	// https://web.archive.org/web/20120614044757/http://www.niksula.hut.fi/~hkankaan/Homepages/gravity.html
	accel := phys.Gravity * gml.DeltaTime()
	phys.Speed.Y += accel
	self.Y += (phys.Speed.Y + accel/2) * gml.DeltaTime()
}
