package game

import (
	"math"

	"github.com/silbinarywolf/gml-go/gml"
)

type Physics struct {
	Speed   gml.Vec
	Gravity float64
}

func (phys *Physics) Update(self *gml.Object) {
	// Formula taken from:
	// https://web.archive.org/web/20120614044757/http://www.niksula.hut.fi/~hkankaan/Homepages/gravity.html
	dt := gml.DeltaTime()
	accel := phys.Gravity * dt
	phys.Speed.Y += accel
	if dt == 1 {
		self.Y += phys.Speed.Y
	} else {
		// Experimental higher fps
		//self.Y += (phys.Speed.Y + accel/2) * dt
		self.Y += dt*phys.Speed.Y + math.Pow(dt, 2)*accel
	}
}
