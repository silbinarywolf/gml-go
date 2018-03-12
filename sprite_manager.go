package gml

import (
	"fmt"
	"image"
	"os"

	"github.com/hajimehoshi/ebiten"
)

var g_spriteManager = newSpriteManager()

func newSpriteManager() SpriteManager {
	manager := SpriteManager{}
	manager.assetMap = make(map[string]*Sprite)
	return manager
}

type SpriteManager struct {
	assetMap map[string]*Sprite
}

func LoadSprite(name string) *Sprite {
	manager := g_spriteManager

	// Use already loaded asset
	if res, ok := manager.assetMap[name]; ok {
		return res
	}

	// Load frames
	//
	// NOTE(Jake): 2018-03-12
	//
	// This is slow but makes managing assets simpler. Will most likely add something to pack
	// everything into a texture page for "production" builds.
	//
	folderPath := fmt.Sprintf("%s/assets/sprites/%s/", currentDirectory(), name)
	frames := make([]*ebiten.Image, 0, 10)
	for i := 0; true; i++ {
		path := fmt.Sprintf("%s%d.png", folderPath, i)
		imageFileData, err := os.Open(path)
		if err != nil {
			if i == 0 {
				panic(fmt.Errorf("Unable to find image: %s", path))
			}
			break
		}
		image, _, err := image.Decode(imageFileData)
		imageFileData.Close()
		if err != nil {
			panic(fmt.Errorf("Unable to decode image: %s", path))
		}
		sheet, err := ebiten.NewImageFromImage(image, ebiten.FilterDefault)
		if err != nil {
			panic(fmt.Errorf("Unable to use image with ebiten.NewImageFromImage: %s", path))
		}
		frames = append(frames, sheet)
	}

	result := new(Sprite)
	result.frames = frames
	manager.assetMap[name] = result
	return result
}
