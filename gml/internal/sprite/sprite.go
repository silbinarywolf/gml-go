package sprite

import (
	"github.com/silbinarywolf/gml-go/gml/internal/math"
)

type Sprite struct {
	name       string
	frames     []SpriteFrame
	size       math.Vec
	imageSpeed float64
}

func (spr *Sprite) Name() string   { return spr.name }
func (spr *Sprite) Size() math.Vec { return spr.size }

/*func (spr *Sprite) GetFrame(index int) *SpriteFrame {
	return &spr.frames[index]
}*/

func newSprite(name string, frames []SpriteFrame, config spriteConfig) *Sprite {
	spr := new(Sprite)
	spr.name = name
	spr.frames = frames
	spr.imageSpeed = config.ImageSpeed

	width := 0
	height := 0
	for _, frame := range frames {
		frameWidth, frameHeight := frame.Size()
		if width < frameWidth {
			width = frameWidth
		}
		if height < frameHeight {
			height = frameHeight
		}
	}
	spr.size = math.Vec{
		X: float64(width),
		Y: float64(height),
	}
	return spr
}
