// +build !headless

package gml

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/text"
	"github.com/silbinarywolf/gml-go/gml/internal/sprite"
)

func DrawSprite(spr *sprite.Sprite, subimage float64, position Vec) {
	screen := gScreen
	{
		camPos := cameraGetActive().Vec
		position.X -= camPos.X
		position.Y -= camPos.Y
	}

	frame := sprite.GetFrame(spr, int(subimage))
	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(position.X, position.Y)
	screen.DrawImage(frame, &op)
}

// draw_sprite_ext( sprite, subimg, x, y, xscale, yscale, rot, colour, alpha );
func DrawSpriteExt(spr *sprite.Sprite, subimage float64, position Vec, scale Vec) {
	screen := gScreen
	{
		camPos := cameraGetActive().Vec
		position.X -= camPos.X
		position.Y -= camPos.Y
	}

	frame := sprite.GetFrame(spr, int(subimage))
	op := ebiten.DrawImageOptions{}
	// op.GeoM.Scale(width/float64(ew), height/float64(eh))
	op.GeoM.Scale(scale.X, scale.Y)
	op.GeoM.Translate(position.X, position.Y)
	screen.DrawImage(frame, &op)
}

func DrawRectangle(pos Vec, size Vec, col color.RGBA) {
	screen := gScreen
	ebitenutil.DrawRect(screen, pos.X, pos.Y, size.X, size.Y, col)
}

func DrawText(position Vec, message string) {
	screen := gScreen
	if !g_fontManager.hasFontSet() {
		panic("Must call DrawSetFont() before calling DrawText.")
	}
	text.Draw(screen, message, g_fontManager.currentFont.font, int(position.X), int(position.Y), color.White)
}
