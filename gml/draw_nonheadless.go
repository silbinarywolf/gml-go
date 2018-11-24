// +build !headless

package gml

import (
	"fmt"
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

func DrawGetGUI() bool {
	return isDrawGuiMode
}

func DrawSetGUI(guiMode bool) {
	isDrawGuiMode = guiMode
}

func DrawSprite(spriteIndex sprite.SpriteIndex, subimage float64, position geom.Vec) {
	DrawSpriteExt(spriteIndex, subimage, position, geom.Vec{1, 1}, 1.0)
}

func DrawSpriteScaled(spriteIndex sprite.SpriteIndex, subimage float64, position geom.Vec, scale geom.Vec) {
	DrawSpriteExt(spriteIndex, subimage, position, scale, 1.0)
}

// draw_sprite_ext( sprite, subimg, x, y, xscale, yscale, rot, colour, alpha );
func DrawSpriteExt(spriteIndex sprite.SpriteIndex, subimage float64, position geom.Vec, scale geom.Vec, alpha float64) {
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

	frame := sprite.GetRawFrame(spriteIndex, int(math.Floor(subimage)))
	op := ebiten.DrawImageOptions{}
	op.GeoM.Scale(scale.X, scale.Y)
	op.GeoM.Translate(position.X, position.Y)

	op.ColorM.Scale(1.0, 1.0, 1.0, alpha)
	//op.Colorgeom.RotateHue(float64(360))

	drawGetTarget().DrawImage(frame, &op)
}

func DrawRectangle(position geom.Vec, size geom.Vec, col color.Color) {
	position = maybeApplyOffsetByCamera(position)

	ebitenutil.DrawRect(drawGetTarget(), position.X, position.Y, size.X, size.Y, col)
}

func DrawRectangleBorder(position geom.Vec, size geom.Vec, color color.Color, borderSize float64, borderColor color.Color) {
	position = maybeApplyOffsetByCamera(position)
	ebitenutil.DrawRect(drawGetTarget(), position.X, position.Y, size.X, size.Y, borderColor)
	position.X += borderSize
	position.Y += borderSize
	size.X -= borderSize * 2
	size.Y -= borderSize * 2
	ebitenutil.DrawRect(drawGetTarget(), position.X, position.Y, size.X, size.Y, color)
}

func DrawText(position geom.Vec, message string) {
	if !g_fontManager.hasFontSet() {
		panic("Must call DrawSetFont() before calling DrawText.")
	}
	position = maybeApplyOffsetByCamera(position)
	text.Draw(drawGetTarget(), message, g_fontManager.currentFont.font, int(position.X), int(position.Y), color.White)
}

func DrawTextColor(position geom.Vec, message string, col color.Color) {
	if !g_fontManager.hasFontSet() {
		panic("Must call DrawSetFont() before calling DrawText.")
	}
	position = maybeApplyOffsetByCamera(position)
	text.Draw(drawGetTarget(), message, g_fontManager.currentFont.font, int(position.X), int(position.Y), col)
}

func DrawTextF(position Vec, message string, args ...interface{}) {
	message = fmt.Sprintf(message, args...)
	DrawText(position, message)
}

func drawGetTarget() *ebiten.Image {
	if camera := cameraGetActive(); camera != nil {
		return camera.screen
	}
	return gScreen
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
