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
	isLoaded   bool
	frames     []SpriteFrame
	size       geom.Vec
	imageSpeed float64
}

func (spr *Sprite) Name() string { return spr.name }

func (spr *Sprite) rect() geom.Rect {
	return geom.Rect{
		Vec:  geom.Vec{},
		Size: spr.size,
	}
}

type SpriteIndex int32

func (spriteIndex SpriteIndex) Name() string {
	sprite := &gSpriteManager.assetList[spriteIndex]
	if !sprite.isLoaded {
		panic("sprite is not loaded, cannot retrieve name")
	}
	return sprite.name
}

func (spriteIndex SpriteIndex) Size() geom.Vec {
	sprite := &gSpriteManager.assetList[spriteIndex]
	if !sprite.isLoaded {
		panic("sprite is not loaded, cannot retrieve size")
	}
	return sprite.size
}

func (spriteIndex SpriteIndex) ImageSpeed() float64 {
	sprite := &gSpriteManager.assetList[spriteIndex]
	if !sprite.isLoaded {
		panic("sprite is not loaded, cannot retrieve imageSpeed")
	}
	return sprite.imageSpeed
}

func (spriteIndex SpriteIndex) isLoaded() bool {
	sprite := &gSpriteManager.assetList[spriteIndex]
	return sprite.isLoaded
}

func (spriteIndex SpriteIndex) get() *Sprite {
	sprite := &gSpriteManager.assetList[spriteIndex]
	if !sprite.isLoaded {
		panic("sprite is not loaded")
	}
	return sprite
}

/*func (spr *Sprite) GetFrame(index int) *SpriteFrame {
	return &spr.frames[index]
}*/

func newSprite(name string, frames []SpriteFrame, config spriteConfig) *Sprite {
	spr := new(Sprite)
	spr.isLoaded = true
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
		spr.size = geom.Vec{
			X: float64(width),
			Y: float64(height),
		}
	}
	return spr
}
