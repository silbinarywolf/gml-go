package game

import (
	"math"

	"github.com/silbinarywolf/gml-go/gml"
)

type Worm struct {
	gml.Object
	Start      gml.Vec
	SinCounter float64
}

func (self *Worm) Create() {
	self.SetSprite(SprWormHead)

	self.Start.X = 304
	self.Start.Y = 528

	self.Vec = self.Start
}

func (self *Worm) Update() {
	self.SinCounter += 0.5

	self.Y = self.Start.Y + math.Round(math.Sin(self.SinCounter*0.15)*21) // y = ystart + round(sin(alarm[1]*0.15)*21); // 0.15, 21
}
