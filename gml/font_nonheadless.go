// +build !headless

package gml

import "golang.org/x/image/font"

type fontData struct {
	font font.Face
}

// DrawSetFont() will set the font to be used for DrawText functions
func DrawSetFont(font FontIndex) {
	fontData := &gFontManager.assetList[font]
	if fontData.font == nil {
		// Load font if not yet loaded
		fontLoad(font)
	}
	gFontManager.currentFont = font
}

// StringWidth() will return the width of the input string in pixels.
func StringWidth(text string) float64 {
	fontFace := fontFont(gFontManager.currentFont)
	if fontFace == nil {
		return 0
	}
	x := font.MeasureString(fontFace, text)
	return float64(x.Round())
}
