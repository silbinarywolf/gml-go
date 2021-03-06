package gml

import (
	"encoding/json"
	"errors"
	"io/ioutil"

	"github.com/golang/freetype/truetype"
	"github.com/silbinarywolf/gml-go/gml/internal/file"
	"golang.org/x/image/font"
)

var gFontManager = newFontManager()

const (
	fontDirectoryBase = "font"
	fontDefaultSize   = 16
	fontDefaultDPI    = 96
)

type fontManager struct {
	currentFont      FontIndex
	assetList        []fontData
	assetNameToIndex map[string]FontIndex
	assetIndexToName []string
}

type fontConfig struct {
	Name     string  `json:"Name"`
	FontSize float64 `json:"FontSize"`
	DPI      float64 `json:"DPI"`
}

func hasFontSet() bool {
	return gFontManager.currentFont != fntUndefined
}

func newFontManager() *fontManager {
	return &fontManager{}
}

// InitFontGeneratedData is used by code generated by gmlgo so you can query a font by index or name
func InitFontGeneratedData(indexToName []string, nameToIndex map[string]FontIndex) {
	gFontManager.assetIndexToName = indexToName
	gFontManager.assetNameToIndex = nameToIndex
	gFontManager.assetList = make([]fontData, len(gFontManager.assetIndexToName))
}

func fontFont(fontIndex FontIndex) font.Face {
	fontData := &gFontManager.assetList[fontIndex]
	if fontData.font == nil {
		return nil
	}
	return fontData.font
}

func fontLoad(fontIndex FontIndex) {
	fontData := &gFontManager.assetList[fontIndex]

	// If already loaded, return early
	if fontData.font != nil {
		return
	}

	name := gFontManager.assetIndexToName[fontIndex]

	// Load font config
	config := fontConfig{}
	{
		configPath := AssetDirectory() + "/" + fontDirectoryBase + "/" + name + "/config.json"
		fileData, err := file.OpenFile(configPath)
		if err != nil {
			panic(errors.New("Unable to find font: " + configPath + ". Error: " + err.Error()))
		}
		bytes, err := ioutil.ReadAll(fileData)
		if err != nil {
			panic("Error loading load config.json for font: " + configPath + "\n" + "Error: " + err.Error())
		}
		if err := json.Unmarshal(bytes, &config); err != nil {
			panic("Error unmarshalling load config.json for font: " + configPath + "\n" + "Error: " + err.Error())
		}
		if config.Name == "" {
			panic("Missing \"Name\" from " + configPath)
		}
		// Setup defaults
		if config.FontSize == 0 {
			config.FontSize = fontDefaultSize
		}
		if config.DPI == 0 {
			config.DPI = fontDefaultDPI
		}
	}

	// Load ttf
	path := AssetDirectory() + "/" + fontDirectoryBase + "/data/" + config.Name
	fileData, err := file.OpenFile(path)
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
	font := truetype.NewFace(tt, &truetype.Options{
		Size:    config.FontSize, // 16px if DPI is 96
		DPI:     config.DPI,
		Hinting: font.HintingFull,
	})

	fontData.font = font
}
