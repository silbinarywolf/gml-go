package game

import (
	"github.com/silbinarywolf/gml-go/gml"
)

const (
	WormHoleSpeed = 7
)

type WormHole struct {
	gml.Object
}

func (self *WormHole) Create() {
	self.SetSprite(SprWormHole)
}

func (self *WormHole) Update() {
	self.X -= WormHoleSpeed
	if self.X+self.Size.X < 0 {
		gml.InstanceDestroy(self)
	}
}
