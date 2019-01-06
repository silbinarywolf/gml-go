package game

import (
	"math"

	"github.com/silbinarywolf/gml-go/gml"
)

const (
	WormStartingBodyParts = 4
	WormLeapPower         = 21
)

type Worm struct {
	gml.Object
	WormDrag

	Speed   gml.Vec
	Gravity float64

	Start        gml.Vec
	SinCounter   float64
	LastBodyPart gml.InstanceIndex
	Dead         bool
	InAir        bool
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
	self.LastBodyPart = parentIndex

	// todo(Jake): 2019-01-06
	// Move this to an alarm
	GameSpawnWall(self.RoomInstanceIndex())
}

func (self *Worm) Update() {
	self.Speed.Y += self.Gravity
	self.Y += self.Speed.Y

	if self.Dead {
		return
	}

	self.WormDrag.Update(self.BaseObject())
	self.SinCounter += 0.5

	// Jump
	{
		hasPressedJumpButton := gml.MouseCheckPressed(gml.MbLeft) ||
			gml.KeyboardCheckPressed(gml.VkSpace)
		if hasPressedJumpButton &&
			!self.InAir &&
			self.Top() > 0 {
			self.Speed.Y = -WormLeapPower
			self.Y = self.Start.Y
			self.InAir = true
		}
	}

	//
	if self.Speed.Y < 0 &&
		!self.InAir {
		self.SetImageIndex(0)
		self.InAir = true
	} else if self.Speed.Y > 0 &&
		self.Y > self.Start.Y {
		self.InAir = false
	}

	//
	if !self.InAir {
		self.Gravity = 0
		self.Speed.Y = 0
		self.Y = self.Start.Y + math.Round(math.Sin(self.SinCounter*0.15)*21)
	} else {
		if self.Speed.Y < 0 {
			self.Gravity = 0.66
		} else {
			self.Gravity = 0.56
		}
	}
}
