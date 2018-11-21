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

type spriteManager struct {
	assetMap  map[string]*Sprite
	assetList []*Sprite
}

func newSpriteManager() *spriteManager {
	manager := &spriteManager{}
	manager.assetMap = make(map[string]*Sprite)
	manager.assetList = make([]*Sprite, 1, 10)
	return manager
}

func SpriteList() []*Sprite {
	return g_spriteManager.assetList[1:]
}

func LoadSprite(name string) *Sprite {
	manager := g_spriteManager

	// Use already loaded asset
	if res, ok := manager.assetMap[name]; ok {
		return res
	}
	result := loadSprite(name)
	manager.assetMap[name] = result
	manager.assetList = append(manager.assetList, result)

	return result
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

func loadSprite(name string) *Sprite {
	// If debug mode, write out the sprite to *.data file
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
	return result
}
