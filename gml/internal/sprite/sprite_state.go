package sprite

import (
	"github.com/silbinarywolf/gml-go/gml/internal/dt"
	"github.com/silbinarywolf/gml-go/gml/internal/geom"
)

type SpriteState struct {
	spriteIndex SpriteIndex
	ImageScale  geom.Vec
	imageIndex  float64
}

func GetCollisionMask(spriteIndex SpriteIndex, imageIndex int, kind int) *CollisionMask {
	spr := spriteIndex.get()
	if spr == nil {
		return nil
	}
	return &spr.frames[imageIndex].collisionMasks[kind]
}

func (state *SpriteState) SpriteIndex() SpriteIndex { return state.spriteIndex }
func (state *SpriteState) sprite() SpriteIndex      { return state.spriteIndex }
func (state *SpriteState) ImageIndex() float64      { return state.imageIndex }
func (state *SpriteState) ImageSpeed() float64 {
	if state.spriteIndex == SprUndefined {
		return 0
	}
	return state.spriteIndex.ImageSpeed()
}
func (state *SpriteState) ImageNumber() float64 {
	if state.spriteIndex == SprUndefined {
		return 0
	}
	spr := state.spriteIndex.get()
	return float64(len(spr.frames))
}

func (state *SpriteState) SetSprite(spriteIndex SpriteIndex) {
	if state.spriteIndex != spriteIndex {
		if !spriteIndex.isLoaded() {
			SpriteLoad(spriteIndex)
		}
		state.spriteIndex = spriteIndex
		state.imageIndex = 0
	}
}

func (state *SpriteState) SetImageIndex(imageIndex float64) {
	state.imageIndex = imageIndex
	imageNumber := state.ImageNumber()
	if imageNumber > 0 {
		for state.imageIndex >= imageNumber {
			state.imageIndex -= imageNumber
		}
		if state.imageIndex < 0 {
			state.imageIndex = 0
		}
	}
}

func (state *SpriteState) ImageUpdate() {
	imageNumber := state.ImageNumber()
	if imageNumber > 0 {
		imageSpeed := state.ImageSpeed() * dt.DeltaTime()
		state.imageIndex += imageSpeed
		for state.imageIndex >= state.ImageNumber() {
			state.imageIndex -= state.ImageNumber()
		}
	}
}
