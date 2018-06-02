package gml

import (
	"errors"
	"io/ioutil"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"golang.org/x/image/font"
)

var g_fontManager = newFontManager()

type FontManager struct {
	currentFont *Font
	assetMap    map[string]*Font
}

func (manager *FontManager) hasFontSet() bool {
	return manager.currentFont != nil && manager.currentFont.font != nil
}

func newFontManager() *FontManager {
	return &FontManager{
		assetMap: make(map[string]*Font),
	}
}

type FontSettings struct {
	DPI  float64
	Size float64
}

func LoadFont(name string, settings FontSettings) *Font {
	manager := g_fontManager

	// Use already loaded asset
	if result, ok := manager.assetMap[name]; ok {
		return result
	}

	path := WorkingDirectory() + "/assets/fonts/" + name + ".ttf"
	fileData, err := ebitenutil.OpenFile(path)
	if err != nil {
		panic(errors.New("Unable to find font: " + path + ". Error: " + err.Error()))
	}
	defer fileData.Close()
	b, err := ioutil.ReadAll(fileData)
	if err != nil {
		panic(errors.New("Unable to read font file into bytes: " + path))
	}
	tt, err := truetype.Parse(b)
	if err != nil {
		panic(errors.New("Unable to parse true type font file: " + path + ", err: " + err.Error()))
	}

	// Setup defaults
	if settings.DPI == 0 {
		settings.DPI = 72
	}
	if settings.Size == 0 {
		settings.Size = 12 // 12pt == 16px
	}
	font := truetype.NewFace(tt, &truetype.Options{
		Size:    settings.Size,
		DPI:     settings.DPI,
		Hinting: font.HintingFull,
	})

	result := new(Font)
	result.font = font
	manager.assetMap[name] = result
	return result
}
