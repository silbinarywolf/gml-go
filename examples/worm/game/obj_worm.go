package game

import (
	"math"

	"github.com/silbinarywolf/gml-go/examples/worm/game/input"
	"github.com/silbinarywolf/gml-go/gml"
)

const (
	WormStartingBodyParts = 2
	WormMaxBodyParts      = 10
	WormLeapPower         = -21
	WormJumpGravity       = 0.66
	WormFallGravity       = 0.56
	WormDieGravity        = 0.58
)

type Worm struct {
	gml.Object
	Physics
	WormLag
	WallSpawner
	bodyParts [WormMaxBodyParts]WormBody

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
	self.YLag = self.Y

	startPos := self.Vec
	for i := 0; i < WormStartingBodyParts; i++ {
		bodyPart := &self.bodyParts[i]
		bodyPart.X = startPos.X - bodyPart.SeperationWidth()
		bodyPart.Y = startPos.Y
		bodyPart.HasSprouted = true

		startPos = bodyPart.Vec
	}

	// Create body
	/*parentIndex := self.InstanceIndex()
	for i := 0; i < WormStartingBodyParts; i++ {
		inst := gml.InstanceCreate(self.X, self.Y, self.RoomInstanceIndex(), ObjWormBody).(*WormBody)
		inst.Parent = parentIndex
		inst.Master = self.InstanceIndex()
		inst.Index = inst.Index + 1
		parentIndex = inst.InstanceIndex()
	}
	self.LastBodyPart = parentIndex*/

}

func (self *Worm) TriggerDeath() {
	if !self.Dead {
		SndWormDie.Play()
		self.SetSprite(SprWormHeadDead)
		self.Dead = true

		// Show game over menu
		gml.InstanceCreate(0, 0, self.RoomInstanceIndex(), ObjMenuGameover)

		// Leap into air at death
		self.Speed.Y = WormLeapPower
		self.Gravity = WormDieGravity
	}
}

func (self *Worm) Draw() {
	// Draw body parts
	// in reverse so they layer correctly
	for i := len(self.bodyParts) - 1; i >= 0; i-- {
		bodyPart := self.bodyParts[i]
		if !bodyPart.HasSprouted {
			continue
		}
		gml.DrawSprite(SprWormBody, 0, bodyPart.X, bodyPart.Y)
	}

	// Draw worm
	self.Object.Draw()
}

func (self *Worm) Update() {
	self.WallSpawner.Update(self.RoomInstanceIndex())
	self.Physics.Update(&self.Object)

	{
		// NOTE(Jake): 2019-02-03 - #82
		// We must update ylag timers after the vspeed/physics and
		// we must loop it like this so it feels like the Game Maker
		// version
		if self.LagTimer.Repeat(2) {
			self.YLag = self.Y
		}
		for i := 0; i < len(self.bodyParts); i++ {
			// Alarm 11
			bodyPart := &self.bodyParts[i]
			//if bodyPart.LagTimer.Repeat(2) {
			bodyPart.YLag = bodyPart.Y
			//}
		}
	}
	for i := 0; i < len(self.bodyParts); i++ {
		// Begin Step
		bodyPart := &self.bodyParts[i]
		var parentX, parentYLag float64
		if i == 0 {
			parentX = self.X
			parentYLag = self.YLag // self.Vec
		} else {
			prevBodyPart := &self.bodyParts[i-1]
			parentX = prevBodyPart.X
			parentYLag = prevBodyPart.YLag
		}
		bodyPart.X = parentX - bodyPart.SeperationWidth()
		bodyPart.Y = parentYLag
	}

	/*prevPosY := self.Y
	if self.DragTimer.Repeat(1) {
		bodyPart := &self.bodyParts[0]
		if bodyPart.HasSprouted {
			bodyPart.Y = prevPosY
			prevPosY = bodyPart.Y
		}
	}
	for i := 1; i < WormStartingBodyParts; i++ {
		bodyPart := &self.bodyParts[i]
		if !bodyPart.HasSprouted {
			continue
		}
		if bodyPart.Y != prevPosY {
			bodyPart.Y = prevPosY
			break
		}
		prevPosY = bodyPart.Y
	}*/
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
		inst, ok := id.Get().(*Checkpoint)
		if !ok {
			continue
		}
		self.Score += 1
		gml.InstanceDestroy(inst)
	}

	if self.LagTimer.Repeat(2) {
		self.YLag = self.Y
	}
}
