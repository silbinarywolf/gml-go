package gml

import "github.com/hajimehoshi/ebiten"

type Sprite struct {
	frames []*ebiten.Image
	// todo(Jake): Get image speed from config.json
}

//draw_sprite(sprite, subimg, x, y);
func (sprite *Sprite) DrawSprite(subimage float64, position Vec) {
	frame := sprite.frames[int(subimage)]

	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(position.X, position.Y)
	g_screen.DrawImage(frame, &op)
}

type SpriteState struct {
	sprite      *Sprite
	imageIndex  float64
	imageNumber float64
	imageSpeed  float64
}

func (state *SpriteState) Sprite() *Sprite      { return state.sprite }
func (state *SpriteState) ImageIndex() float64  { return state.imageIndex }
func (state *SpriteState) ImageNumber() float64 { return state.imageNumber }
func (state *SpriteState) ImageSpeed() float64  { return state.imageSpeed }

func (state *SpriteState) SetSprite(sprite *Sprite) {
	state.sprite = sprite
	state.imageIndex = 0
	state.imageNumber = float64(len(sprite.frames))
	// todo(Jake): 2018-03-12 - Reset to config.json sprite_speed
}

func (state *SpriteState) SetImageIndex(newImageIndex float64) {
	state.imageIndex = newImageIndex
	if state.imageIndex > state.imageNumber {
		state.imageIndex = 0
	}
	if state.imageIndex < 0 {
		state.imageIndex = 0
	}
}

func (state *SpriteState) imageUpdate() {
	imageSpeed := state.imageSpeed // * dt
	state.imageIndex += imageSpeed
	if state.imageIndex > state.imageNumber {
		state.imageIndex = 0
	}
}
