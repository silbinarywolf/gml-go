package game

import (
	"github.com/silbinarywolf/gml-go/gml"
)

type Checkpoint struct {
	gml.Object
}

func (self *Checkpoint) Create() {
	self.Size.X = 32
	self.Size.Y = WindowHeight
}

func (self *Checkpoint) Update() {
	self.X -= WallSpeed()
	if self.X+self.Size.X < 0 {
		gml.InstanceDestroy(self)
	}
}
