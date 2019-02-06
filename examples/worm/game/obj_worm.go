package game

import (
	"math"

	"github.com/silbinarywolf/gml-go/examples/worm/game/input"
	"github.com/silbinarywolf/gml-go/examples/worm/game/wall"
	"github.com/silbinarywolf/gml-go/gml"
)

const (
	WormStartingBodyParts = 1
	WormMaxBodyParts      = 10
	WormLeapPower         = -21
	WormJumpGravity       = 0.66
	WormFallGravity       = 0.56
	WormDieGravity        = 0.58
	WormSproutSpeed       = 0.05
	WormSinCounterStart   = 9999999 // used in original game. Could just count up instead but kept it
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
	FlapCounter  float64
	WingCount    float64
}

func (self *Worm) Create() {
	self.SetSprite(SprWormHead)
	self.SetDepth(DepthWorm)
	self.Start.X = 304
	self.Start.Y = 528
	self.Score = 0
	self.Vec = self.Start
	self.YLag = self.Y
	self.FlapCounter = 0
	self.SinCounter = WormSinCounterStart

	startPos := self.Vec
	for i := 0; i < WormStartingBodyParts; i++ {
		bodyPart := &self.bodyParts[i]
		bodyPart.X = startPos.X - bodyPart.SeperationWidth()
		bodyPart.Y = startPos.Y
		bodyPart.YLag = bodyPart.Y
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

func (self *Worm) ScoreIncrease() {
	self.Score++
	switch self.Score {
	case 2, 5, 9, 15:
		// Add body
		for i := 0; i < len(self.bodyParts); i++ {
			bodyPart := &self.bodyParts[i]
			if !bodyPart.HasSprouted {
				parentBodyPart := &self.bodyParts[i-1]
				bodyPart.X = parentBodyPart.X
				bodyPart.Y = parentBodyPart.Y
				bodyPart.HasSprouted = true
				break
			}
		}
	case 22, 30, 39, 50, 65:
		// Add Wing
		self.WingCount++

		// Update walls
		switch self.WingCount {
		case 1:
			Global.Notification.SetNotification("You got a wing\n\nEach wing will add an extra jump")

			// Change wall list to be harder and have pipes that require 1 wing
			self.WallList = self.WallList[:0]
			self.WallList = append(self.WallList, wall.WallSetFlatHard...)
			self.WallList = append(self.WallList, wall.WallSetFly1...)
			/*
				with(objController)
				{
					// Remove any easy walls
					ds_list_copy(wall, wall_flat_hard)
					ds_list_append(wall, wall_fly_1)
				}
				// Get Wings
				notification_set("You got a wing!##Each wing will add an extra jump!", 0.3)
				wall_offscreen_kill()
				with(objController)
				{
					alarm[0] = room_speed * 4
				}
			*/
		case 2:
			//ds_list_append(wall, wall_fly_2)
		case 3:
			//ds_list_append(wall, wall_fly_3)
		case 4:
			//ds_list_append(wall, wall_fly_4)
		case 5:
			//ds_list_append(wall, wall_fly_5)
		default:
			panic("invalid wing number")
		}
	}
}

func (self *Worm) Draw() {
	// Draw body parts
	// in reverse so they layer correctly
	for i := len(self.bodyParts) - 1; i >= 0; i-- {
		bodyPart := &self.bodyParts[i]
		if !bodyPart.HasSprouted {
			continue
		}
		gml.DrawSprite(SprWormBody, 0, bodyPart.X, bodyPart.Y)
		if self.WingCount > float64(i) {
			imageIndex := 0.0
			if self.Speed.Y < 0 &&
				self.FlapCounter == float64(i+1) {
				imageIndex = 1
			}
			gml.DrawSprite(SprWing, imageIndex, bodyPart.X-12, bodyPart.Y+24)
		}
	}

	// Draw worm
	self.Object.Draw()
}

func (self *Worm) Update() {
	// Pre Begin Step
	self.WallSpawner.Update(self.RoomInstanceIndex())
	self.Physics.Update(&self.Object)

	// Worm Body Parts
	{
		// Begin Step
		{
			// NOTE(Jake): 2019-02-04 - #82
			// To make the worm body parts feel like the original
			// the YLag body part update needs to happen before Physics
			// update and the "Alarms" need to be after the loop.
			for i := 0; i < len(self.bodyParts); i++ {
				// Begin Step
				bodyPart := &self.bodyParts[i]
				var parentX, parentYLag float64
				if i == 0 {
					parentX = self.X
					parentYLag = self.YLag // self.Vec
				} else {
					parentBodyPart := &self.bodyParts[i-1]
					parentX = parentBodyPart.X
					parentYLag = parentBodyPart.YLag
				}

				if bodyPart.HasSprouted {
					bodyPart.SproutLerp += WormSproutSpeed
					if bodyPart.SproutLerp > 1 {
						bodyPart.SproutLerp = 1
					}
				}

				bodyPart.X = parentX - (bodyPart.SeperationWidth() * bodyPart.SproutLerp)
				bodyPart.Y = parentYLag

			}
		}

		// Alarm 11
		{
			if self.LagTimer.Repeat(2) {
				self.YLag = self.Y
				for i := 0; i < len(self.bodyParts); i++ {
					// Alarm 11
					bodyPart := &self.bodyParts[i]
					bodyPart.YLag = bodyPart.Y
				}
			}
		}
	}

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
			(!self.InAir || self.FlapCounter < self.WingCount) &&
			self.Top() > 0 {
			if !self.InAir {
				self.Speed.Y = WormLeapPower
				self.Y = self.Start.Y
				self.SinCounter = WormSinCounterStart
				self.InAir = true
			} else {
				self.Speed.Y = WormLeapPower / 2
				self.FlapCounter++
			}
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
		self.FlapCounter = 0
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

		self.SinCounter -= 1 * gml.DeltaTime()

		// Taken from: y = ystart + round(sin(alarm[1]*0.15)*21);
		self.Y = self.Start.Y + math.Round(math.Sin(self.SinCounter*0.15)*21)
	} else {
		if self.Speed.Y < 0 {
			self.Gravity = WormJumpGravity
		} else {
			self.Gravity = WormFallGravity
		}
	}

	HandleCollisionForWormOrWormPart(&self.Object, self)
	/*for i, _ := range self.bodyParts {
		bodyPart := &self.bodyParts[i]
		if bodyPart.HasSprouted {
			HandleCollisionForWormOrWormPart(&bodyPart, self)
		}
	}*/

	for _, id := range gml.CollisionRectList(self, self.X, self.Y) {
		inst, ok := id.Get().(*Checkpoint)
		if !ok {
			continue
		}
		self.ScoreIncrease()
		gml.InstanceDestroy(inst)
	}
}
