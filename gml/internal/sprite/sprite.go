package sprite

import (
	"github.com/silbinarywolf/gml-go/gml/internal/geom"
)

const (
	maxCollisionMasks = 3
)

const SprUndefined SpriteIndex = 0

type Sprite struct {
	name       string
	frames     []SpriteFrame
	size       geom.Size
	imageSpeed float64
}

func (spr *Sprite) Name() string   { return spr.name }
func (spr *Sprite) isLoaded() bool { return len(spr.frames) > 0 }
func (spr *Sprite) rect() geom.Rect {
	return geom.Rect{
		Vec:  geom.Vec{},
		Size: spr.size,
	}
}

type SpriteIndex int32

func (spriteIndex SpriteIndex) Name() string    { return gSpriteManager.assetList[spriteIndex].name }
func (spriteIndex SpriteIndex) Size() geom.Size { return gSpriteManager.assetList[spriteIndex].size }
func (spriteIndex SpriteIndex) ImageSpeed() float64 {
	return gSpriteManager.assetList[spriteIndex].imageSpeed
}
func (spriteIndex SpriteIndex) IsValid() bool {
	return spriteIndex > 0
}
func (spriteIndex SpriteIndex) IsLoaded() bool {
	return len(gSpriteManager.assetList[spriteIndex].frames) > 0
}

func Frames(spriteIndex SpriteIndex) []SpriteFrame {
	return gSpriteManager.assetList[spriteIndex].frames
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
