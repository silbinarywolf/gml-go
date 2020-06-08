package game

import (
	"math"

	"github.com/silbinarywolf/gml-go/example/worm/asset"
	"github.com/silbinarywolf/gml-go/example/worm/game/input"
	"github.com/silbinarywolf/gml-go/example/worm/game/wall"
	"github.com/silbinarywolf/gml-go/gml"
	"github.com/silbinarywolf/gml-go/gml/alarm"
)

const (
	WormStartingBodyParts = 1
	WormMaxBodyParts      = 10
	WormStartX            = 304
	WormStartY            = 528
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
	BodyParts [WormMaxBodyParts]WormBody

	sinTimer           alarm.Alarm
	dirtCreateTimer    alarm.Alarm
	bodyTimer          alarm.Alarm
	inputDisabledTimer alarm.Alarm

	Start       gml.Vec
	Score       float64
	Dead        bool
	InAir       bool
	FlapCounter float64
	WingCount   float64
}

func (self *Worm) Create() {
	self.SetSprite(asset.SprWormHead)
	self.SetDepth(DepthWorm)

	self.Start.X = WormStartX
	self.Start.Y = WormStartY
	self.YLag = self.Y
	self.Vec = self.Start

	self.Reset()
	self.WallSpawner.SpawnWallTimer.Set(0)
}

func (self *Worm) Reset() {
	self.WallSpawner.Reset()

	self.Score = 0
	self.FlapCounter = 0
	self.WingCount = 0
	self.SetStartingBodyParts(WormStartingBodyParts)

	// DEBUG: Test
	//for i := 0; i < 23; i++ {
	//	self.ScoreIncrease()
	//}
}

// SetStartingBodyParts allows the Create() event and tests
// setup how many body parts trail the worm
func (self *Worm) SetStartingBodyParts(bodyParts int) {
	prevPos := self.Vec
	for i, _ := range self.BodyParts {
		bodyPart := &self.BodyParts[i]
		*bodyPart = WormBody{}
		bodyPart.X = prevPos.X
		bodyPart.Y = prevPos.Y
		bodyPart.YLag = bodyPart.Y

		prevPos = bodyPart.Vec
	}
	for i := 0; i < bodyParts; i++ {
		bodyPart := &self.BodyParts[i]
		bodyPart.SproutLerp = 1
		bodyPart.HasSprouted = true
	}
}

func (self *Worm) TriggerDeath() {
	if !self.Dead {
		if !Global.SoundDisabled {
			asset.SndWormDie.Play()
		}
		self.SetSprite(asset.SprWormHeadDead)
		self.Dead = true

		// Show game over menu
		self.CalcMedals()
		self.RoomIndex().InstanceCreate(0, 0, ObjMenuGameover)

		// Leap into air at death
		self.Speed.Y = WormLeapPower
		self.Gravity = WormDieGravity
	}
}

func (self *Worm) CalcMedals() {
	wormMedal := MedalNone
	flightMedal := MedalNone

	// Calculate body parts
	bodyPartCount := 0
	for _, bodyPart := range self.BodyParts {
		if !bodyPart.HasSprouted {
			break
		}
		bodyPartCount++
	}

	// Worm Body Medal
	if bodyPartCount >= 3 {
		wormMedal = MedalBronze
		if bodyPartCount >= 5 {
			wormMedal = MedalSilver
		}
	}

	// Worm Flight Medal
	if self.WingCount >= 2 {
		flightMedal = MedalBronze
		if self.WingCount >= 4 {
			flightMedal = MedalSilver
		}
	}

	//
	Global.PreviousRound = Global.CurrentRound
	if wormMedal > Global.CurrentRound.MedalWorm {
		Global.CurrentRound.MedalWorm = wormMedal
	}
	if flightMedal > Global.CurrentRound.MedalWing {
		Global.CurrentRound.MedalWing = flightMedal
	}
}

func (self *Worm) ScoreIncrease() {
	self.Score++
	switch self.Score {
	case 2, 5, 9, 15:
		// Add body
		for i := 0; i < len(self.BodyParts); i++ {
			bodyPart := &self.BodyParts[i]
			if !bodyPart.HasSprouted {
				parentBodyPart := &self.BodyParts[i-1]
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
			// Change wall list to be harder and have pipes that require 1 wing
			self.WallList = self.WallList[:0]
			self.WallList = append(self.WallList, wall.WallSetFlatHard...)
			self.WallList = append(self.WallList, wall.WallSetFly1...)

			// Get wings
			Global.Notification.SetNotification("You got a wing\n\nEach wing will add an extra jump")
			self.WallSpawner.Reset()
			self.SpawnWallTimer.Set(DesignedMaxTPS * 4)
		case 2:
			self.WallList = append(self.WallList, wall.WallSetFly2...)
		case 3:
			self.WallList = append(self.WallList, wall.WallSetFly3...)
		case 4:
			self.WallList = append(self.WallList, wall.WallSetFly4...)
		case 5:
			self.WallList = append(self.WallList, wall.WallSetFly5...)
		default:
			panic("invalid wing number")
		}
	}
}

func (self *Worm) Draw() {
	// Draw body parts
	// in reverse so they layer correctly
	for i := len(self.BodyParts) - 1; i >= 0; i-- {
		bodyPart := &self.BodyParts[i]
		if !bodyPart.HasSprouted {
			continue
		}
		gml.DrawSprite(asset.SprWormBody, 0, bodyPart.X, bodyPart.Y)
		if self.WingCount > float64(i) {
			imageIndex := 0.0
			if self.Speed.Y < 0 &&
				self.FlapCounter == float64(i+1) {
				imageIndex = 1
			}
			gml.DrawSprite(asset.SprWing, imageIndex, bodyPart.X-12, bodyPart.Y+24)
		}
	}

	// Draw worm
	self.Object.Draw()
}

func (self *Worm) Update() {
	// Pre Begin Step
	self.WallSpawner.Update(self.RoomIndex())

	defer func() {
		self.Physics.Update(&self.Object)
	}()

	// Worm Body Parts
	{
		// Begin Step
		{
			// NOTE(Jake): 2019-02-04 - #82
			// To make the worm body parts feel like the original
			// the YLag body part update needs to happen before Physics
			// update and the "Alarms" need to be after the loop.
			if self.bodyTimer.Repeat(1) {
				for i := 0; i < len(self.BodyParts); i++ {
					// Begin Step
					bodyPart := &self.BodyParts[i]
					var parentX, parentYLag float64
					if i == 0 {
						parentX = self.X
						parentYLag = self.YLag // self.Vec
					} else {
						parentBodyPart := &self.BodyParts[i-1]
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
		}

		// Alarm 11
		{
			if self.LagTimer.Repeat(2) {
				self.YLag = self.Y
				for i := 0; i < len(self.BodyParts); i++ {
					// Alarm 11
					bodyPart := &self.BodyParts[i]
					bodyPart.YLag = bodyPart.Y
				}
			}
		}
	}

	self.sinTimer.Repeat(WormSinCounterStart)

	if self.Dead {
		return
	}

	if self.dirtCreateTimer.Repeat(2) {
		if !self.InAir {
			self.RoomIndex().InstanceCreate(self.X, self.Y, ObjWormHole)
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
				//self.sinTimer.Set(WormSinCounterStart - 1)
				self.InAir = true
			} else {
				self.Speed.Y = WormLeapPower / 2
				self.FlapCounter++
			}
			if !Global.SoundDisabled {
				asset.SndJump.Play()
			}
		}
	}

	//
	if self.Speed.Y < 0 &&
		!self.InAir {
		self.SetImageIndex(0)
		self.InAir = true
	} else if self.Speed.Y > 0 &&
		self.Y > self.Start.Y {
		self.sinTimer.Set(WormSinCounterStart)
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

		// Taken from: y = ystart + round(sin(alarm[1]*0.15)*21);
		self.Y = self.Start.Y + math.Round(math.Sin(self.sinTimer.Get()*0.15)*21)
	} else {
		if self.Speed.Y < 0 {
			self.Gravity = WormJumpGravity
		} else {
			self.Gravity = WormFallGravity
		}
	}

	// Handle collision for worm + body parts
	HandleCollisionForWormOrWormPart(self, self.X, self.Y)
	for i, _ := range self.BodyParts {
		bodyPart := &self.BodyParts[i]
		if bodyPart.HasSprouted {
			HandleCollisionForWormOrWormPart(self, bodyPart.X, bodyPart.Y)
		}
	}

	// Get score
	for _, id := range gml.CollisionRectList(self, self.X, self.Y) {
		inst, ok := id.Get().(*Checkpoint)
		if !ok {
			continue
		}
		self.ScoreIncrease()
		gml.InstanceDestroy(inst)
	}
}
