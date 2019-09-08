package game

import (
	"image/color"

	"github.com/silbinarywolf/gml-go/example/worm/asset"
	"github.com/silbinarywolf/gml-go/gml"
)

type LerpDir int

const (
	LerpIn  LerpDir = 0
	LerpOut LerpDir = 1
)

const (
	LerpSpeed = 0.1
	LerpTimer = 4 * (LerpSpeed * 60)
)

type Notification struct {
	Active  bool
	Text    string
	Lerp    float64
	LerpDir LerpDir
}

func (self *Notification) SetNotification(text string) {
	*self = Notification{}
	self.Text = text
	self.Active = true
}

func (self *Notification) Update() {
	if !self.Active {
		return
	}
	switch self.LerpDir {
	case LerpIn:
		self.Lerp += LerpSpeed
		if self.Lerp > 1.0+LerpTimer {
			self.Lerp = 1.0
			self.LerpDir = LerpOut
		}
	case LerpOut:
		self.Lerp -= LerpSpeed
		if self.Lerp < 0-LerpTimer {
			self.Active = false
		}
	}
}

func (self *Notification) Draw() {
	//NotificationBackX := (display_get_gui_width() >> 1) - (notification_width >> 1)
	var x, y float64
	x = gml.CameraGetViewSize(0).X/2 - asset.SprNotificationBg.Size().X/2
	y = -110 + (200 * LerpGain(self.Lerp, 0.25))
	gml.DrawSprite(asset.SprNotificationBg, 0, x, y)
	x += 15
	y += 15
	// NOTE(Jake): 2019-02-04
	// Original source code used: draw_set_color(make_color_rgb(3,3,3))
	gml.DrawTextColor(x, y, self.Text, color.RGBA{3, 3, 3, 255})
}

// Lerp Bias
// http://blog.demofox.org/2012/09/24/bias-and-gain-are-your-friend/
// http://demofox.org/biasgain.html
//
// LerpBias was a script in the original Game Maker source code and has
// been copied across as is, including the comments referencing sources above.
func LerpBias(time, bias float64) float64 {
	if time > 1.0 {
		time = 1.0
	} else if time < 0 {
		time = 0.0
	}
	return (time / ((((1.0 / bias) - 2.0) * (1.0 - time)) + 1.0))
}

// Lerp Gain
// http://blog.demofox.org/2012/09/24/bias-and-gain-are-your-friend/
// http://demofox.org/biasgain.html
//
// LerpGain was a script in the original Game Maker source code and has
// been copied across as is, including the comments referencing sources above.
func LerpGain(time, gain float64) float64 {
	if time < 0.5 {
		return LerpBias(time*2.0, gain) * 0.5
	} else {
		return LerpBias(time*2.0-1.0, 1.0-gain)*0.5 + 0.5
	}
}
