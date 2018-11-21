package sprite

import (
	"github.com/silbinarywolf/gml-go/gml/internal/geom"
)

const (
	maxCollisionMasks = 3
)

type Sprite struct {
	name       string
	frames     []SpriteFrame
	size       geom.Size
	imageSpeed float64
}

func (spr *Sprite) Name() string        { return spr.name }
func (spr *Sprite) Size() geom.Size     { return spr.size }
func (spr *Sprite) ImageSpeed() float64 { return spr.imageSpeed }
func (spr *Sprite) rect() geom.Rect {
	return geom.Rect{
		Vec:  geom.Vec{},
		Size: spr.Size(),
	}
}

func GetCollisionMask(spr *Sprite, frame int, kind int) *CollisionMask {
	// masks := &spr.frames[frame].collisionMasks[kind].masks[kind]
	//if len(masks) == 0 {
	//	panic("Should have at least 1 collision mask defined")
	//}
	return &spr.frames[frame].collisionMasks[kind]
}

/*func (spr *Sprite) GetFrame(index int) *SpriteFrame {
	return &spr.frames[index]
}*/

func newSprite(name string, frames []SpriteFrame, config spriteConfig) *Sprite {
	spr := new(Sprite)
	spr.name = name
	spr.frames = frames
	spr.imageSpeed = config.ImageSpeed

	if len(frames) > 0 {
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
			//frame.collisionMasks = make([]SpriteCollisionMask, maxCollisionMasks)
		}
		spr.size = geom.Size{
			X: int32(width),
			Y: int32(height),
		}
	}
	return spr
}
