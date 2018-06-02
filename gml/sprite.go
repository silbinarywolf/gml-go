package gml

import "github.com/hajimehoshi/ebiten"

type Sprite struct {
	name   string
	frames []*ebiten.Image
	size   Vec
	// todo(Jake): Get image speed from config.json
}

func newSprite(name string, frames []*ebiten.Image) *Sprite {
	sprite := new(Sprite)
	sprite.name = name
	sprite.frames = frames

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
	sprite.size = Vec{
		X: float64(width),
		Y: float64(height),
	}
	return sprite
}

//draw_sprite(sprite, subimg, x, y);
func (sprite *Sprite) DrawSprite(subimage float64, position Vec) {
	frame := sprite.frames[int(subimage)]

	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(position.X-currentCamera.X, position.Y-currentCamera.Y)
	gScreen.DrawImage(frame, &op)
}
