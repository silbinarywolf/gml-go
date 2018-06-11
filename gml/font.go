package gml

import (
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

func DrawSetFont(font *Font) {
	g_fontManager.currentFont = font
}
