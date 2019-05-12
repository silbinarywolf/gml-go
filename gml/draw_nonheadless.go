// +build !headless

package gml

import (
	"image/color"
	"math"
	"strings"

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
	position.X = math.Floor(position.X)
	position.Y = math.Floor(position.Y)

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
	DrawSpriteExt(spriteIndex, subimage, x, y, 0, geom.Vec{1, 1}, alpha)
}

func DrawSpriteScaled(spriteIndex sprite.SpriteIndex, subimage float64, x, y float64, scale geom.Vec) {
	DrawSpriteExt(spriteIndex, subimage, x, y, 0, scale, 1.0)
}

func DrawSpriteRotated(spriteIndex sprite.SpriteIndex, subimage float64, x, y float64, rotation float64) {
	DrawSpriteExt(spriteIndex, subimage, x, y, rotation, geom.Vec{1, 1}, 1.0)
}

func DrawSpriteExt(spriteIndex sprite.SpriteIndex, subimage float64, x, y float64, rotation float64, scale geom.Vec, alpha float64) {
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
	op.GeoM.Rotate(rotation / 57.2958)
	op.GeoM.Translate(position.X, position.Y)
	op.ColorM.Reset()
	op.ColorM.Scale(1.0, 1.0, 1.0, alpha)
	//op.Colorgeom.RotateHue(float64(360))

	drawGetTarget().DrawImage(frame, op)
}

// DrawSpriteColor will draw a sprite with a color blend
func DrawSpriteColor(spriteIndex sprite.SpriteIndex, subimage float64, x, y float64, col color.Color) {
	if spriteIndex == sprite.SprUndefined {
		// If no sprite in use, draw nothing
		return
	}
	position := geom.Vec{
		X: x,
		Y: y,
	}
	position = maybeApplyOffsetByCamera(position)

	r, g, b, a := col.RGBA()
	frame := sprite.GetRawFrame(spriteIndex, int(math.Floor(subimage)))
	op.GeoM.Reset()
	op.GeoM.Translate(position.X, position.Y)
	op.ColorM.Reset()
	op.ColorM.Scale(float64(r)/255, float64(g)/255, float64(b)/255, float64(a)/255)

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

func DrawRectangleBorder(x, y, w, h float64, col color.Color, borderSize float64, borderColor color.Color) {
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
	if col != color.Transparent {
		drawRect(drawGetTarget(), position.X, position.Y, size.X, size.Y, col)
	}
}

func DrawText(x, y float64, message string, col color.Color) {
	DrawTextColor(x, y, message, col)
}

func DrawTextColorAlpha(x, y float64, message string, col color.Color, alpha float64) {
	//if !hasFontSet() {
	//	panic("Must call DrawSetFont() before calling DrawText.")
	//}
	r, g, b, a := col.RGBA()
	c := color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
	c.A = uint8(float64(c.A) * alpha)
	DrawTextColor(x, y, message, c)
}

func DrawTextColor(x, y float64, message string, col color.Color) {
	if !hasFontSet() {
		panic("Must call DrawSetFont() before calling DrawText.")
	}
	fontFace := fontFont(gFontManager.currentFont)
	if fontFace == nil {
		return
	}

	// NOTE(Jake): 2019-05-12
	// Add initial space so text draws from the left-top corner.
	// Changed from previous string height calculation method
	{
		ascent := float64(fontFace.Metrics().Ascent.Ceil())
		y += ascent
	}

	// Draw lines
	pos := maybeApplyOffsetByCamera(Vec{
		X: x,
		Y: y,
	})
	leadingHeight := float64(fontFace.Metrics().Height.Ceil())
	for _, line := range strings.Split(message, "\n") {
		text.Draw(drawGetTarget(), line, fontFace, int(pos.X), int(pos.Y), col)
		pos.Y += leadingHeight
	}
}

func drawGetTarget() *ebiten.Image {
	// NOTE(Jake): 2019-01-26
	// "gCameraManager.camerasEnabledCount > 1" is here so that we render directly to
	// gScreen if we are only using 1 camera.
	if camera := cameraGetActive(); camera != nil && gCameraManager.camerasEnabledCount > 1 {
		return camera.surface
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
	op.ColorM.Reset()
	op.ColorM.Scale(colorScale(clr))
	// Filter must be 'nearest' filter (default).
	// Linear filtering would make edges blurred.
	dst.DrawImage(emptyImage, op)
}
