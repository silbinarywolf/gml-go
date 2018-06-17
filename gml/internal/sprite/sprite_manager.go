package sprite

import (
	"errors"
	"strconv"

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
	folderPath := file.AssetsDirectory + "/sprites/" + name + "/"
	frames := make([]SpriteFrame, 0, 10)
	for i := 0; true; i++ {
		path := folderPath + strconv.Itoa(i) + ".png"
		frame, err := createFrame(path, i)
		if err != nil {
			if i == 0 {
				panic(errors.New("Unable to find image: " + path))
			}
			break
		}
		frames = append(frames, frame)
	}

	// Read config information (if it exists)
	var config spriteConfig
	configPath := folderPath + "config.json"
	config = loadConfig(configPath)

	// Create sprite
	result := newSprite(name, frames, config)
	manager.assetMap[name] = result

	return result
}
