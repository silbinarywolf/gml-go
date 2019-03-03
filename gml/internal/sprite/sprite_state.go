package sprite

import (
	"bytes"
	"encoding/gob"

	"github.com/silbinarywolf/gml-go/gml/internal/dt"
	"github.com/silbinarywolf/gml-go/gml/internal/geom"
)

type spriteStateSerialize struct {
	SpriteIndex SpriteIndex
	ImageScale  geom.Vec
	ImageIndex  float64
}

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

func (state SpriteState) MarshalBinary() ([]byte, error) {
	w := spriteStateSerialize{
		SpriteIndex: state.spriteIndex,
		ImageScale:  state.ImageScale,
		ImageIndex:  state.imageIndex,
	}
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(w); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (state *SpriteState) UnmarshalBinary(data []byte) error {
	w := spriteStateSerialize{}
	reader := bytes.NewReader(data)
	dec := gob.NewDecoder(reader)
	if err := dec.Decode(&w); err != nil {
		return err
	}
	state.SetSprite(w.SpriteIndex)
	state.SetImageIndex(w.ImageIndex)
	state.ImageScale = w.ImageScale
	return nil
}
