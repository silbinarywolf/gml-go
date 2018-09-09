// +build !headless

package gml

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/text"

	"github.com/silbinarywolf/gml-go/gml/internal/geom"
	"github.com/silbinarywolf/gml-go/gml/internal/sprite"
)

var (
	isDrawGuiMode = false
)

func DrawSetGUI(guiMode bool) {
	isDrawGuiMode = guiMode
}

func DrawSprite(spr *sprite.Sprite, subimage float64, position geom.Vec) {
	screen := gScreen
	position = maybeApplyOffsetByCamera(position)
	frame := sprite.GetRawFrame(spr, int(math.Floor(subimage)))
	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(position.X, position.Y)
	screen.DrawImage(frame, &op)
}

func DrawSpriteScaled(spr *sprite.Sprite, subimage float64, position geom.Vec, scale geom.Vec) {
	DrawSpriteExt(spr, subimage, position, scale, 1.0)
}

// draw_sprite_ext( sprite, subimg, x, y, xscale, yscale, rot, colour, alpha );
func DrawSpriteExt(spr *sprite.Sprite, subimage float64, position geom.Vec, scale geom.Vec, alpha float64) {
	screen := gScreen
	position = maybeApplyOffsetByCamera(position)
	// NOTE(Jake): 2018-07-09
	//
	// This doesn't work. A cleaner solution might be to
	// render everything to a seperate image if possible then
	// scale that and render.
	//
	// Since this is only really needed for the map editor, I dont
	// have a problem with it.
	//
	//scale.X *= view.Scale().X
	//scale.Y *= view.Scale().Y

	frame := sprite.GetRawFrame(spr, int(math.Floor(subimage)))
	op := ebiten.DrawImageOptions{}
	op.GeoM.Scale(scale.X, scale.Y)
	op.GeoM.Translate(position.X, position.Y)

	op.ColorM.Scale(1.0, 1.0, 1.0, alpha)
	//op.Colorgeom.RotateHue(float64(360))

	screen.DrawImage(frame, &op)
}

func DrawRectangle(position geom.Vec, size geom.Vec, col color.Color) {
	screen := gScreen
	position = maybeApplyOffsetByCamera(position)

	ebitenutil.DrawRect(screen, position.X, position.Y, size.X, size.Y, col)
}

func DrawRectangleBorder(position geom.Vec, size geom.Vec, color color.Color, borderSize float64, borderColor color.Color) {
	screen := gScreen
	position = maybeApplyOffsetByCamera(position)
	ebitenutil.DrawRect(screen, position.X, position.Y, size.X, size.Y, borderColor)
	position.X += borderSize
	position.Y += borderSize
	size.X -= borderSize * 2
	size.Y -= borderSize * 2
	ebitenutil.DrawRect(screen, position.X, position.Y, size.X, size.Y, color)
}

func DrawText(position geom.Vec, message string) {
	screen := gScreen
	if !g_fontManager.hasFontSet() {
		panic("Must call DrawSetFont() before calling DrawText.")
	}
	position = maybeApplyOffsetByCamera(position)
	text.Draw(screen, message, g_fontManager.currentFont.font, int(position.X), int(position.Y), color.White)
}

func DrawTextColor(position geom.Vec, message string, col color.Color) {
	screen := gScreen
	if !g_fontManager.hasFontSet() {
		panic("Must call DrawSetFont() before calling DrawText.")
	}
	position = maybeApplyOffsetByCamera(position)
	text.Draw(screen, message, g_fontManager.currentFont.font, int(position.X), int(position.Y), col)
}

func maybeApplyOffsetByCamera(position geom.Vec) geom.Vec {
	if !isDrawGuiMode {
		if view := cameraGetActive(); view != nil {
			position.X -= view.X
			position.Y -= view.Y
		}
	}
	return position
}
