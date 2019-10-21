// +build debug

package gml

import (
	"image/color"

	"github.com/silbinarywolf/gml-go/gml/internal/geom"
)

func drawInputText(pos *geom.Vec, label string, text string, isFocused bool) bool {
	size := geom.Vec{100, 20}
	DrawText(pos.X, pos.Y, label, color.White)
	pos.Y += 12
	borderCol := color.RGBA{255, 255, 255, 255}
	isMouseOver := debugDrawIsMouseOver(*pos, size)
	if isMouseOver {
		borderCol = color.RGBA{255, 255, 0, 255}
	}
	if isFocused {
		text = KeyboardString() + "|"
		borderCol = color.RGBA{255, 0, 0, 255}
	}
	DrawRectangleBorder(pos.X, pos.Y, size.X, size.Y, color.Black, 2, borderCol)
	DrawText(pos.X+8, pos.Y+16, text, color.White)
	pos.Y += size.Y
	if MouseCheckPressed(MbLeft) && isMouseOver {
		if !isFocused {
			SetKeyboardString(text)
		}
		return true
	}
	if isFocused &&
		(KeyboardCheckPressed(VkEnter) || KeyboardCheckPressed(VkNumpadEnter)) {
		return true
	}
	return false
}

func drawButton(pos geom.Vec, text string) bool {
	// Config
	paddingH := 32.0
	borderWidth := 2.0
	size := geom.Vec{StringWidth(text) + paddingH, 24}

	// Handle mouse over
	isMouseOver := debugDrawIsMouseOver(pos, size)
	var innerRectColor color.RGBA
	if isMouseOver {
		innerRectColor = color.RGBA{180, 180, 180, 255}
	} else {
		innerRectColor = color.RGBA{255, 255, 255, 255}
	}

	// Draw Border (outer rect)
	DrawRectangleBorder(pos.X, pos.Y, size.X, size.Y, innerRectColor, borderWidth, color.RGBA{0, 162, 232, 255})
	/*	pos.X += borderWidth
		pos.Y += borderWidth
		size.X -= borderWidth * 2
		size.Y -= borderWidth * 2

		// Draw Rect (inner rect)
		DrawRectangle(pos, size, innerRectColor)*/

	// Draw Text
	pos.X += paddingH * 0.5
	pos.Y += 16
	DrawText(pos.X, pos.Y, text, color.Black)
	return MouseCheckPressed(MbLeft) && isMouseOver
}
