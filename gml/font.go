package gml

import (
	"golang.org/x/image/font"
)

const fntUndefined FontIndex = 0

type FontIndex int32

type fontData struct {
	font font.Face
}

// DrawSetFont will set the font index to be used for DrawText functions
func DrawSetFont(font FontIndex) {
	fontData := &gFontManager.assetList[font]
	if fontData.font == nil {
		// Load font if not yet loaded
		fontLoad(font)
	}
	gFontManager.currentFont = font
}

// DrawGetFont will get the font index used for DrawText functions
func DrawGetFont() FontIndex {
	return gFontManager.currentFont
}

// StringWidth will return the width of the input string in pixels.
func StringWidth(text string) float64 {
	fontFace := fontFont(gFontManager.currentFont)
	if fontFace == nil {
		return 0
	}
	width := 0
	start := 0
	prevI := 0
	for i, c := range text {
		if c == '\n' {
			newWidth := font.MeasureString(fontFace, text[start:prevI]).Ceil()
			start = i
			if newWidth > width {
				width = newWidth
			}
		}
		prevI = i
	}
	newWidth := font.MeasureString(fontFace, text[start:]).Ceil()
	if newWidth > width {
		width = newWidth
	}
	return float64(width)
}

// StringHeight will return height of the input string in pixels.
func StringHeight(text string) float64 {
	fontFace := fontFont(gFontManager.currentFont)
	if fontFace == nil {
		return 0
	}
	bounds, _ := font.BoundString(fontFace, text)
	//log.Printf("StringHeight: %v\n", bounds)
	return float64(-bounds.Min.Y.Ceil())
	//return float64(x.Round())
}
