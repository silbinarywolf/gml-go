package sprite

import (
	"bytes"
	"encoding/gob"
	"errors"
	"io/ioutil"

	"github.com/silbinarywolf/gml-go/gml/internal/file"
)

var (
	g_spriteManager = newSpriteManager()
)

func newSpriteManager() SpriteManager {
	manager := SpriteManager{}
	manager.assetMap = make(map[string]*Sprite)
	return manager
}

type SpriteManager struct {
	assetMap map[string]*Sprite
}

func loadSpriteFromData(name string) *spriteAsset {
	path := file.AssetsDirectory + "/sprites/" + name + ".data"
	fileData, err := file.OpenFile(path)
	if err != nil {
		//panic(errors.New("Unable to find image: " + path))
		return nil
	}
	bytesData, err := ioutil.ReadAll(fileData)
	fileData.Close()
	if err != nil {
		panic(errors.New("Unable to read bytes from image: " + path))
	}
	buf := bytes.NewReader(bytesData)
	asset := new(spriteAsset)
	gob.NewDecoder(buf).Decode(asset)
	return asset
}

func LoadSprite(name string) *Sprite {
	manager := g_spriteManager

	// Use already loaded asset
	if res, ok := manager.assetMap[name]; ok {
		return res
	}

	// If debug mode, write out the sprite
	debugWriteSprite(name)

	// Load from *.data
	spriteAsset := loadSpriteFromData(name)
	if spriteAsset == nil {
		panic("Unable to load sprite from data file: " + name)
	}
	frames := make([]SpriteFrame, len(spriteAsset.Frames))
	for i := 0; i < len(spriteAsset.Frames); i++ {
		frameAsset := spriteAsset.Frames[i]
		frame, err := createFrame(frameAsset)
		if err != nil {
			panic("Sprite frame load error for \"" + name + "\": " + err.Error())
		}
		frames[i] = frame
	}

	// Create sprite
	result := newSprite(spriteAsset.Name, frames, spriteConfig{
		ImageSpeed: spriteAsset.ImageSpeed,
	})
	manager.assetMap[name] = result

	return result
}
