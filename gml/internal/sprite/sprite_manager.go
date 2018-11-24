package sprite

import (
	"bytes"
	"encoding/gob"
	"errors"
	"io/ioutil"

	"github.com/silbinarywolf/gml-go/gml/internal/file"
)

const (
	SpriteDirectoryBase = "sprite"
)

var (
	g_spriteManager = newSpriteManager()
)

type spriteManager struct {
	assetList         []Sprite
	spriteNameToIndex map[string]SpriteIndex
	spriteIndexToName []string
}

func newSpriteManager() *spriteManager {
	manager := &spriteManager{}
	return manager
}

func SpriteInitializeIndexToName(indexToName []string, nameToIndex map[string]SpriteIndex) {
	g_spriteManager.spriteIndexToName = indexToName
	g_spriteManager.spriteNameToIndex = nameToIndex
	g_spriteManager.assetList = make([]Sprite, len(g_spriteManager.spriteIndexToName))
}

// todo(Jake): 2018-24-11 - Github #14
// Remove SpriteList() as it's brittle and only used by sprite_selector.go
func SpriteList() []*Sprite {
	panic("Broke sprite_selector(), need to fix")
	return nil
	//return g_spriteManager.assetList[1:]
}

func SpriteLoadByName(name string) *Sprite {
	index := g_spriteManager.spriteNameToIndex[name]
	return SpriteLoad(index)
}

func SpriteLoad(index SpriteIndex) *Sprite {
	manager := g_spriteManager

	sprite := &manager.assetList[index]
	// todo(Jake): have a "isUsed" var or function instead of checking
	// for frames
	if sprite.isUsed() {
		return sprite
	}
	name := g_spriteManager.spriteIndexToName[index]
	// todo(Jake): change loadSprite() to return Sprite, not *Sprite
	result := loadSprite(name)
	*sprite = *result
	return result
}

func loadSpriteFromData(name string) *spriteAsset {
	path := file.AssetDirectory + "/" + SpriteDirectoryBase + "/" + name + ".data"
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
