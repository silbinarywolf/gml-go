package sprite

import (
	"github.com/silbinarywolf/gml-go/gml/internal/math"

	"github.com/hajimehoshi/ebiten"
)

type Sprite struct {
	name   string
	frames []*ebiten.Image
	size   math.Vec
	// todo(Jake): Get image speed from config.json
}

func (spr *Sprite) Name() string   { return spr.name }
func (spr *Sprite) Size() math.Vec { return spr.size }

func (spr *Sprite) GetFrame(index int) *ebiten.Image {
	return spr.frames[index]
}

func newSprite(name string, frames []*ebiten.Image) *Sprite {
	spr := new(Sprite)
	spr.name = name
	spr.frames = frames

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
