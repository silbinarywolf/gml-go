package gml

import "golang.org/x/image/font"

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

// StringWidth() will return the width of the input string in pixels.
func StringWidth(text string) float64 {
	fontFace := fontFont(gFontManager.currentFont)
	if fontFace == nil {
		return 0
	}
	x := font.MeasureString(fontFace, text)
	return float64(x.Round())
}
