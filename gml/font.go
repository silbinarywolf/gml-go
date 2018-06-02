package gml

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
)

type Font struct {
	font font.Face
}

// NOTE(Jake): 2018-06-02
//
// Technically related to font code, so I'll put it here for now
//
func StringWidth(text string) float64 {
	currentFont := g_fontManager.currentFont
	if currentFont == nil {
		return 0
	}
	face := currentFont.font
	x := font.MeasureString(face, text)
	return float64(x.Round())
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
