// +build !headless

package gml

import (
	"fmt"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"

	"github.com/silbinarywolf/gml-go/gml/internal/geom"
	"github.com/silbinarywolf/gml-go/gml/internal/sprite"
)

var (
	isDrawGuiMode = false
	emptyImage    *ebiten.Image
	op            = &ebiten.DrawImageOptions{}
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

	// NOTE(Jake): 2019-01-26 - #91
	// DrawSprite is like DrawSpriteExt except we apply no scaling or colorM
	// changes. To avoid unnecessary allocations, we pulled out the code from
	// DrawSpriteExt and placed it here with a few modifications.
	frame := sprite.GetRawFrame(spriteIndex, int(math.Floor(subimage)))
	op.GeoM.Reset()
	op.GeoM.Translate(position.X, position.Y)
	op.ColorM.Reset()

	drawGetTarget().DrawImage(frame, op)
}

func DrawSpriteAlpha(spriteIndex sprite.SpriteIndex, subimage float64, x, y float64, alpha float64) {
	DrawSpriteExt(spriteIndex, subimage, x, y, geom.Vec{1, 1}, alpha)
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
	op.GeoM.Reset()
	op.GeoM.Scale(scale.X, scale.Y)
	op.GeoM.Translate(position.X, position.Y)
	op.ColorM.Reset()
	op.ColorM.Scale(1.0, 1.0, 1.0, alpha)
	//op.Colorgeom.RotateHue(float64(360))

	drawGetTarget().DrawImage(frame, op)
}

func DrawRectangle(x, y, w, h float64, col color.Color) {
	position := geom.Vec{
		X: x,
		Y: y,
	}
	position = maybeApplyOffsetByCamera(position)

	drawRect(drawGetTarget(), position.X, position.Y, w, h, col)
}

func DrawRectangleBorder(x, y, w, h float64, color color.Color, borderSize float64, borderColor color.Color) {
	position := geom.Vec{
		X: x,
		Y: y,
	}
	size := geom.Vec{
		X: w,
		Y: h,
	}
	position = maybeApplyOffsetByCamera(position)
	drawRect(drawGetTarget(), position.X, position.Y, size.X, size.Y, borderColor)
	position.X += borderSize
	position.Y += borderSize
	size.X -= borderSize * 2
	size.Y -= borderSize * 2
	drawRect(drawGetTarget(), position.X, position.Y, size.X, size.Y, color)
}

func DrawText(x, y float64, message string) {
	DrawTextColor(x, y, message, color.White)
}

func DrawTextColorAlpha(x, y float64, message string, col color.Color, alpha float64) {
	if !hasFontSet() {
		panic("Must call DrawSetFont() before calling DrawText.")
	}
	r, g, b, a := col.RGBA()
	c := color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
	c.A = uint8(float64(c.A) * alpha)
	text.Draw(drawGetTarget(), message, fontFont(gFontManager.currentFont), int(x), int(y), c)
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
	// NOTE(Jake): 2019-01-26
	// "gCameraManager.camerasEnabledCount > 1" is here so that we render directly to
	// gScreen if we are only using 1 camera.
	if camera := cameraGetActive(); camera != nil && gCameraManager.camerasEnabledCount > 1 {
		return camera.screen
	}
	return gScreen
}

func init() {
	emptyImage, _ = ebiten.NewImage(16, 16, ebiten.FilterDefault)
	emptyImage.Fill(color.White)
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

func colorScale(clr color.Color) (rf, gf, bf, af float64) {
	r, g, b, a := clr.RGBA()
	if a == 0 {
		return 0, 0, 0, 0
	}

	rf = float64(r) / float64(a)
	gf = float64(g) / float64(a)
	bf = float64(b) / float64(a)
	af = float64(a) / 0xffff
	return
}

// drawRect draws a rectangle on the given destination dst.
//
// DrawRect is intended to be used mainly for debugging or prototyping purpose.
func drawRect(dst *ebiten.Image, x, y, width, height float64, clr color.Color) {
	ew, eh := emptyImage.Size()

	op.GeoM.Reset()
	op.GeoM.Scale(width/float64(ew), height/float64(eh))
	op.GeoM.Translate(x, y)
	op.ColorM.Scale(colorScale(clr))
	// Filter must be 'nearest' filter (default).
	// Linear filtering would make edges blurred.
	dst.DrawImage(emptyImage, op)
}
