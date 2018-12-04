package sprite

import (
	"github.com/silbinarywolf/gml-go/gml/internal/geom"
)

type SpriteState struct {
	spriteIndex SpriteIndex
	ImageScale  geom.Vec
	imageIndex  float64
}

func GetCollisionMask(spriteIndex SpriteIndex, imageIndex int, kind int) *CollisionMask {
	spr := sprite(spriteIndex)
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
	spr := sprite(state.spriteIndex)
	return spr.imageSpeed
}
func (state *SpriteState) ImageNumber() float64 {
	if state.spriteIndex == SprUndefined {
		return 0
	}
	spr := sprite(state.spriteIndex)
	return float64(len(spr.frames))
}

func (state *SpriteState) SetSprite(spriteIndex SpriteIndex) {
	if state.spriteIndex != spriteIndex {
		if !spriteIndex.IsLoaded() {
			SpriteLoad(spriteIndex)
		}
		state.spriteIndex = spriteIndex
		state.imageIndex = 0
	}
}

func (state *SpriteState) SetImageIndex(imageIndex float64) {
	state.imageIndex = imageIndex
	if state.imageIndex >= state.ImageNumber() {
		state.imageIndex = 0
	}
	if state.imageIndex < 0 {
		state.imageIndex = 0
	}
}

func (state *SpriteState) ImageUpdate() {
	imageSpeed := state.ImageSpeed() // * dt
	state.imageIndex += imageSpeed
	if state.imageIndex >= state.ImageNumber() {
		state.imageIndex = 0
	}
}
