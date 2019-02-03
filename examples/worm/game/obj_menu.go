package game

import (
	"image/color"

	"github.com/silbinarywolf/gml-go/gml"
)

const (
	MenuFadeOutSpeed = 0.05
	CreditText       = "Created by Silbinary Wolf | Art by Dansknapp | Music by Magicdweedoo"
)

type Menu struct {
	gml.Object
	ImageAlpha       float64
	IsHoveringOnMenu bool
	IsFadingAway     bool
}

func (self *Menu) Create() {
	self.SetDepth(DepthMenu)
	self.ImageAlpha = 1.0

	size := SprPlayButton.Size()
	screenSize := gml.CameraGetViewSize(0)
	self.X = (screenSize.X / 2) - (size.X / 2)
	self.Y = (screenSize.Y / 2) - (size.Y / 2)
	self.Size = size
}

func (self *Menu) Update() {
	if self.IsFadingAway {
		self.ImageAlpha -= MenuFadeOutSpeed
		if self.ImageAlpha < 0 {
			self.ImageAlpha = 0
		}
		return
	}
	self.IsHoveringOnMenu = self.CollisionPoint(gml.MousePosition())
	if gml.MouseCheckPressed(gml.MbLeft) &&
		self.IsHoveringOnMenu {
		inst, ok := Global.Player.Get().(*Worm)
		if !ok {
			panic("Cannot find Worm")
		}
		inst.WallSpawner.Reset()
		//inst.inputDisabledTimer.Set(DesignedMaxTPS * 0.5)
		//inst.DisableInput = true
		self.IsFadingAway = true
	}
}

func (self *Menu) Draw() {
	screenSize := gml.CameraGetViewSize(0)
	x := 16.0
	y := 16.0

	// Draw sound icon
	{
		spr := SprSoundIcon
		gml.DrawSpriteAlpha(spr, 0, x, y, self.ImageAlpha)
		x += spr.Size().X
	}

	// Draw music icon
	gml.DrawSpriteAlpha(SprMusicIcon, 0, x+4, y, self.ImageAlpha)

	// Draw title
	gml.DrawSpriteAlpha(SprTitle, 0, (screenSize.X/2)-(SprTitle.Size().X/2), 20, self.ImageAlpha)

	// Draw button
	{
		frame := 0.0
		if self.IsHoveringOnMenu {
			frame = 1
		}
		gml.DrawSpriteAlpha(SprPlayButton, frame, self.X, self.Y, self.ImageAlpha)
	}

	// Draw credits
	{
		x := (screenSize.X / 2) - (gml.StringWidth(CreditText) / 2) + 4 // 48.0
		y := screenSize.Y - 35

		gml.DrawTextColorAlpha(x-1, y, CreditText, color.Black, self.ImageAlpha)
		gml.DrawTextColorAlpha(x, y+1, CreditText, color.White, self.ImageAlpha)
	}
}
