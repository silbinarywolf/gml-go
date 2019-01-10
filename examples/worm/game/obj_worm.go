package game

import (
	"math"

	"github.com/silbinarywolf/gml-go/gml"
)

const (
	WormStartingBodyParts = 1
	WormLeapPower         = -21
	WormJumpGravity       = 0.66
	WormFallGravity       = 0.56
	WormDieGravity        = 0.58
)

type Worm struct {
	gml.Object
	Physics
	WormDrag
	WallSpawner

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
	self.YDrag = self.Y
	self.WallSpawner.Reset()

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

}

func (self *Worm) TriggerDeath() {
	if !self.Dead {
		self.SetSprite(SprWormHeadDead)
		self.Dead = true

		// Leap into air at death
		self.Speed.Y = WormLeapPower
		self.Gravity = WormDieGravity
	}
}

func (self *Worm) Update() {
	self.WallSpawner.Update(self.RoomInstanceIndex())
	self.Physics.Update(&self.Object)
	self.WormDrag.Update(&self.Object)

	if self.Dead {
		return
	}

	self.SinCounter += 0.5

	// Jump
	{
		hasPressedJumpButton := gml.MouseCheckPressed(gml.MbLeft) ||
			gml.KeyboardCheckPressed(gml.VkSpace)
		if hasPressedJumpButton &&
			!self.InAir &&
			self.Top() > 0 {
			self.Y = self.Start.Y
			self.Speed.Y = WormLeapPower
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
			self.Gravity = WormJumpGravity
		} else {
			self.Gravity = WormFallGravity
		}
	}

	// Collide with Wall
	for _, id := range gml.CollisionRectList(self, self.Pos()) {
		_, ok := gml.InstanceGet(id).(*Wall)
		if !ok {
			continue
		}
		self.TriggerDeath()
		break
	}
}
