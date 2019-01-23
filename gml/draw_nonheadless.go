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

func DrawSprite(spriteIndex sprite.SpriteIndex, subimage float64, x, y float64) {
	DrawSpriteExt(spriteIndex, subimage, x, y, geom.Vec{1, 1}, 1.0)
}

func DrawSpriteScaled(spriteIndex sprite.SpriteIndex, subimage float64, x, y float64, scale geom.Vec) {
	DrawSpriteExt(spriteIndex, subimage, x, y, scale, 1.0)
}

func DrawSpriteExt(spriteIndex sprite.SpriteIndex, subimage float64, x, y float64, scale geom.Vec, alpha float64) {
	if spriteIndex == sprite.SprUndefined {
		// If no sprite in use, draw nothing
		return
	}
	// draw_sprite_ext( sprite, subimg, x, y, xscale, yscale, rot, colour, alpha );
	position := geom.Vec{
		X: x,
		Y: y,
	}
	position = maybeApplyOffsetByCamera(position)

	frame := sprite.GetRawFrame(spriteIndex, int(math.Floor(subimage)))
	op := ebiten.DrawImageOptions{}
	op.GeoM.Scale(scale.X, scale.Y)
	op.GeoM.Translate(position.X, position.Y)

	op.ColorM.Scale(1.0, 1.0, 1.0, alpha)
	//op.Colorgeom.RotateHue(float64(360))

	drawGetTarget().DrawImage(frame, &op)
}

func DrawRectangle(x, y float64, size geom.Vec, col color.Color) {
	position := geom.Vec{
		X: x,
		Y: y,
	}
	position = maybeApplyOffsetByCamera(position)

	ebitenutil.DrawRect(drawGetTarget(), position.X, position.Y, size.X, size.Y, col)
}

func DrawRectangleBorder(x, y float64, size geom.Vec, color color.Color, borderSize float64, borderColor color.Color) {
	position := geom.Vec{
		X: x,
		Y: y,
	}
	position = maybeApplyOffsetByCamera(position)
	ebitenutil.DrawRect(drawGetTarget(), position.X, position.Y, size.X, size.Y, borderColor)
	position.X += borderSize
	position.Y += borderSize
	size.X -= borderSize * 2
	size.Y -= borderSize * 2
	ebitenutil.DrawRect(drawGetTarget(), position.X, position.Y, size.X, size.Y, color)
}

func DrawText(x, y float64, message string) {
	DrawTextColor(x, y, message, color.White)
}

func DrawTextColor(x, y float64, message string, col color.Color) {
	if !hasFontSet() {
		panic("Must call DrawSetFont() before calling DrawText.")
	}
	text.Draw(drawGetTarget(), message, fontFont(gFontManager.currentFont), int(x), int(y), col)
}

/*func drawText(font FontIndex, message string) {
	text.Draw(drawGetTarget(), message, fontFont(gFontManager.currentFont), int(position.X), int(position.Y), color.White)
}*/

func DrawTextF(x, y float64, format string, args ...interface{}) {
	DrawText(x, y, fmt.Sprintf(format, args...))
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
