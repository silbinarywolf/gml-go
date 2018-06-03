package gml

import (
	"errors"
	"image"
	"strconv"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
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
	folderPath := AssetsDirectory() + "/sprites/" + name + "/"
	frames := make([]*ebiten.Image, 0, 10)
	for i := 0; true; i++ {
		path := folderPath + strconv.Itoa(i) + ".png" //fmt.Sprintf("%s%d.png", folderPath, i)
		imageFileData, err := ebitenutil.OpenFile(path)
		if err != nil {
			if i == 0 {
				panic(errors.New("Unable to find image: " + path))
			}
			break
		}
		image, _, err := image.Decode(imageFileData)
		imageFileData.Close()
		if err != nil {
			panic(errors.New("Unable to decode image: " + path))
		}
		sheet, err := ebiten.NewImageFromImage(image, ebiten.FilterDefault)
		if err != nil {
			panic(errors.New("Unable to use image with ebiten.NewImageFromImage: " + path))
		}
		frames = append(frames, sheet)
	}

	result := newSprite(name, frames)
	manager.assetMap[name] = result
	return result
}
