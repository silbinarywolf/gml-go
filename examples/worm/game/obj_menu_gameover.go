package game

import (
	"github.com/silbinarywolf/gml-go/gml"
)

const (
	MenuGameOverAccelerationSpeed = 3
)

type MenuGameover struct {
	gml.Object
	Physics
	RetryButton      gml.Rect
	IsHoveringOnMenu bool
}

func (self *MenuGameover) Create() {
	self.SetDepth(DepthMenu)

	Global.MusicPlaying.Stop()
	Global.MusicPlaying = SndGameover
	Global.MusicPlaying.Play()

	// Gameover
	{
		self.SetSprite(SprGameover)
		screenSize := gml.CameraGetViewSize(0)
		self.Size = self.SpriteIndex().Size()
		self.X = (screenSize.X / 2) - (self.Size.X / 2)
		self.Y = -self.Size.Y //(screenSize.Y / 2) - (self.Size.Y / 2)
	}

	// Retry Button
	self.RetryButton.X = 153
	self.RetryButton.Y = 95
	self.RetryButton.Size = SprRetryButton.Size()
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
		gml.DrawSprite(SprRetryButton, imageIndex, self.X+self.RetryButton.X, self.Y+self.RetryButton.Y)
	}
	//screenSize := gml.CameraGetViewSize(0)
	//x := 16.0
	//y := 16.0

	//draw_background(bgGameOver, x, y);
	//draw_sprite(sprRetry, retry_img_index, x + retry_x, y + retry_y);

	//draw_sprite(sprMedalWorm, global.best_medals[MEDAL_WORM], x + 41, y + 73);
	//draw_sprite(sprMedalWing, global.best_medals[MEDAL_FLIGHT], x + 290, y + 84);
}
