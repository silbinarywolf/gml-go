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

// DrawGetGUI returns whether Draw functions will draw relative to the screen or not
func DrawGetGUI() bool {
	return isDrawGuiMode
}

// DrawSetGUI allows you to set whether you want to draw relative to the screen (true) or to the world (false)
func DrawSetGUI(guiMode bool) {
	isDrawGuiMode = guiMode
}

func DrawSprite(spriteIndex sprite.SpriteIndex, subimage float64, position geom.Vec) {
	DrawSpriteExt(spriteIndex, subimage, position, geom.Vec{1, 1}, 1.0)
}

func DrawSpriteScaled(spriteIndex sprite.SpriteIndex, subimage float64, position geom.Vec, scale geom.Vec) {
	DrawSpriteExt(spriteIndex, subimage, position, scale, 1.0)
}

func DrawSpriteExt(spriteIndex sprite.SpriteIndex, subimage float64, position geom.Vec, scale geom.Vec, alpha float64) {
	if spriteIndex == sprite.SprUndefined {
		// If no sprite in use, draw nothing
		return
	}
	// draw_sprite_ext( sprite, subimg, x, y, xscale, yscale, rot, colour, alpha );
	position = maybeApplyOffsetByCamera(position)

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
	DrawTextColor(position, message, color.White)
}

func DrawTextColor(position geom.Vec, message string, col color.Color) {
	if !hasFontSet() {
		panic("Must call DrawSetFont() before calling DrawText.")
	}
	text.Draw(drawGetTarget(), message, fontFont(gFontManager.currentFont), int(position.X), int(position.Y), col)
}

/*func drawText(font FontIndex, message string) {
	text.Draw(drawGetTarget(), message, fontFont(gFontManager.currentFont), int(position.X), int(position.Y), color.White)
}*/

func DrawTextF(position Vec, format string, args ...interface{}) {
	DrawText(position, fmt.Sprintf(format, args...))
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
