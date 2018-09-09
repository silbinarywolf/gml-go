package sprite

import (
	"github.com/silbinarywolf/gml-go/gml/internal/geom"
)

type Sprite struct {
	name       string
	frames     []SpriteFrame
	size       geom.Size
	imageSpeed float64
}

func (spr *Sprite) Name() string    { return spr.name }
func (spr *Sprite) Size() geom.Size { return spr.size }

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
	spr.size = geom.Size{
		X: int32(width),
		Y: int32(height),
	}
	return spr
}
