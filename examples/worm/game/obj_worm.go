package game

import (
	"math"

	"github.com/silbinarywolf/gml-go/examples/worm/game/input"
	"github.com/silbinarywolf/gml-go/gml"
)

const (
	WormStartingBodyParts = 2
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
	dirtCreateTimer    gml.Alarm
	inputDisabledTimer gml.Alarm

	Start        gml.Vec
	Score        float64
	SinCounter   float64
	LastBodyPart gml.InstanceIndex
	Dead         bool
	InAir        bool
}

func (self *Worm) Create() {
	self.SetSprite(SprWormHead)
	self.SetDepth(DepthWorm)
	self.Start.X = 304
	self.Start.Y = 528
	self.Score = 0
	self.Vec = self.Start
	self.YDrag = self.Y

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

		// Sounds
		SndSunnyFields.Stop()
		SndWormDie.Play()
		SndGameover.Play()

		// Leap into air at death
		self.Speed.Y = WormLeapPower
		self.Gravity = WormDieGravity
	}
}

func (self *Worm) Update() {
	self.WallSpawner.Update(self.RoomInstanceIndex())
	self.Physics.Update(&self.Object)
	if self.DragTimer.Repeat(1) {
		self.YDrag = self.Y
	}
	//if self.inputDisabledTimer.Tick() {
	//	self.DisableInput = false
	//}

	if self.Dead {
		return
	}

	if self.dirtCreateTimer.Repeat(2) {
		if !self.InAir {
			gml.InstanceCreate(self.X, self.Y, self.RoomInstanceIndex(), ObjWormHole)
		}
	}

	// Jump
	{
		if input.JumpPressed() &&
			!self.InAir &&
			self.Top() > 0 {
			self.Y = self.Start.Y
			self.Speed.Y = WormLeapPower
			self.SinCounter = 0
			self.InAir = true
			SndPlay.Play()
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
		// NOTE(Jake): 2019-01-25
		// The original game gave the worm an image speed of 0.21 when
		// in the ground. To recreate this just add 0.08. The 0.13 should come
		// from the default processing of the animation.
		self.SetImageIndex(self.ImageIndex() + (0.08 * gml.DeltaTime()))
		self.Gravity = 0
		self.Speed.Y = 0

		self.SinCounter += 1 * gml.DeltaTime()
		self.Y = self.Start.Y + math.Round(math.Sin(self.SinCounter*0.15)*21)
	} else {
		if self.Speed.Y < 0 {
			self.Gravity = WormJumpGravity
		} else {
			self.Gravity = WormFallGravity
		}
	}

	HandleCollisionForWormOrWormPart(&self.Object, self)
	for _, id := range gml.CollisionRectList(self, self.X, self.Y) {
		inst, ok := gml.InstanceGet(id).(*Checkpoint)
		if !ok {
			continue
		}
		self.Score += 1
		gml.InstanceDestroy(inst)
	}
}
