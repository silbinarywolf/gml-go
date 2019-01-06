package game

import (
	"math"

	"github.com/silbinarywolf/gml-go/gml"
)

const (
	WormStartingBodyParts = 1
)

type Worm struct {
	gml.Object
	WormDrag

	Start      gml.Vec
	SinCounter float64
	Last       gml.InstanceIndex
	Dead       bool
}

func (self *Worm) Create() {
	self.SetSprite(SprWormHead)

	self.Start.X = 304
	self.Start.Y = 528

	self.Vec = self.Start

	// Create body
	parentIndex := self.InstanceIndex()
	for i := 0; i < WormStartingBodyParts; i++ {
		inst := gml.InstanceCreate(self.X, self.Y, self.RoomInstanceIndex(), ObjWormBody).(*WormBody)
		inst.Parent = parentIndex
		inst.Master = self.InstanceIndex()
		inst.Index = inst.Index + 1
		parentIndex = inst.InstanceIndex()
	}
	self.Last = parentIndex
}

func (self *Worm) Update() {
	self.WormDrag.Update(self.BaseObject())
	self.SinCounter += 0.5
	self.Y = self.Start.Y + math.Round(math.Sin(self.SinCounter*0.15)*21) // y = ystart + round(sin(alarm[1]*0.15)*21); // 0.15, 21

	//gml.InstanceCreate(self.X, self.Y, self.RoomInstanceIndex(), ObjWormHole)
}
