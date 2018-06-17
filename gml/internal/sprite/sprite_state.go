package sprite

import (
	"github.com/silbinarywolf/gml-go/gml/internal/math"
)

type SpriteState struct {
	sprite      *Sprite
	ImageScale  math.Vec
	imageIndex  float64
	imageNumber int
	imageSpeed  float64
}

func (state *SpriteState) Sprite() *Sprite      { return state.sprite }
func (state *SpriteState) ImageIndex() float64  { return state.imageIndex }
func (state *SpriteState) ImageNumber() float64 { return float64(state.imageNumber) }
func (state *SpriteState) ImageSpeed() float64  { return state.imageSpeed }

func (state *SpriteState) SetSprite(sprite *Sprite) {
	state.sprite = sprite
	state.imageIndex = 0
	state.imageNumber = len(sprite.frames)
	state.imageSpeed = sprite.imageSpeed
}

func (state *SpriteState) SetImageIndex(imageIndex float64) {
	state.imageIndex = imageIndex
	if state.imageIndex > state.ImageNumber() {
		state.imageIndex = 0
	}
	if state.imageIndex < 0 {
		state.imageIndex = 0
	}
}

func (state *SpriteState) ImageUpdate() {
	imageSpeed := state.ImageSpeed() // * dt
	state.imageIndex += imageSpeed
	if state.imageIndex > state.ImageNumber() {
		state.imageIndex = 0
	}
}
