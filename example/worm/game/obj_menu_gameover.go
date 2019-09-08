package game

import (
	"github.com/silbinarywolf/gml-go/example/worm/asset"
	"github.com/silbinarywolf/gml-go/gml"
	"github.com/silbinarywolf/gml-go/gml/alarm"
)

const (
	MenuGameOverAccelerationSpeed = 3
)

type MenuGameover struct {
	gml.Object
	Physics
	RetryButton             gml.Rect
	IsHoveringOnMenu        bool
	DisplayScore            GameScore
	MedalDisplayUpdateTimer alarm.Alarm
}

func (self *MenuGameover) Create() {
	self.SetDepth(DepthMenu)

	if Global.MusicPlaying != 0 {
		Global.MusicPlaying.Stop()
	}
	Global.MusicPlaying = asset.MusGameover
	Global.MusicPlaying.Play()

	//
	self.MedalDisplayUpdateTimer.Set(60.0 * 1.0)
	self.DisplayScore = Global.PreviousRound

	// Gameover
	{
		self.SetSprite(asset.SprGameover)
		screenSize := gml.CameraGetViewSize(0)
		self.Size = self.SpriteIndex().Size()
		self.X = (screenSize.X / 2) - (self.Size.X / 2)
		self.Y = -self.Size.Y //(screenSize.Y / 2) - (self.Size.Y / 2)
	}

	// Retry Button
	self.RetryButton.X = 153
	self.RetryButton.Y = 95
	self.RetryButton.Size = asset.SprRetryButton.Size()
}

func (self *MenuGameover) Update() {
	// Animate menu and snap down to the center
	{
		screenSize := gml.CameraGetViewSize(0)
		yCenter := (screenSize.Y / 2) - (self.Size.Y / 2)
		if self.Y != yCenter {
			self.Speed.Y += MenuGameOverAccelerationSpeed
			self.Y += self.Speed.Y
			if self.Y > yCenter {
				self.Y = yCenter
				self.Speed.Y = 0
			}
		}
	}

	// Update medals
	if self.MedalDisplayUpdateTimer.Tick() {
		if self.DisplayScore != Global.CurrentRound {
			if !Global.SoundDisabled {
				asset.SndMedalObtained.Play()
			}
			self.DisplayScore = Global.CurrentRound
		}
	}

	//
	retryButton := self.RetryButton
	retryButton.X += self.X
	retryButton.Y += self.Y
	self.IsHoveringOnMenu = retryButton.CollisionPoint(gml.MousePosition())

	if self.IsHoveringOnMenu &&
		gml.MouseCheckPressed(gml.MbLeft) {
		gml.InstanceDestroy(self)
	}
}

func (self *MenuGameover) Destroy() {
	Global.GameReset()
	Global.MusicRandomizeTrack()
}

func (self *MenuGameover) Draw() {
	// Draw background
	gml.DrawSprite(self.SpriteIndex(), 0, self.X, self.Y)

	// Draw Retry button
	{
		imageIndex := 0.0
		if self.IsHoveringOnMenu {
			imageIndex = 1
		}
		gml.DrawSprite(asset.SprRetryButton, imageIndex, self.X+self.RetryButton.X, self.Y+self.RetryButton.Y)
	}

	// Draw worm medal and flight medal
	gml.DrawSprite(asset.SprMedalWorm, float64(self.DisplayScore.MedalWorm), self.X+41, self.Y+73)
	gml.DrawSprite(asset.SprMedalWing, float64(self.DisplayScore.MedalWing), self.X+290, self.Y+84)
}
