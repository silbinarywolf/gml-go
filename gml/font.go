package gml

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
)

type Font struct {
	font font.Face
}

func DrawText(position Vec, message string) {
	if !g_fontManager.hasFontSet() {
		panic("Must call DrawSetFont() before calling DrawText.")
	}
	text.Draw(gScreen, message, g_fontManager.currentFont.font, int(position.X), int(position.Y), color.White)
}

func DrawSetFont(font *Font) {
	g_fontManager.currentFont = font
}
