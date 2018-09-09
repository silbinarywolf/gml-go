// +build debug

package sprite

import (
	"bytes"
	"encoding/gob"
	"errors"
	"image"
	"image/png"
	"io/ioutil"
	"log"
	"path/filepath"
	"strconv"

	"github.com/fsnotify/fsnotify"
	"github.com/silbinarywolf/gml-go/gml/internal/file"
	"github.com/silbinarywolf/gml-go/gml/internal/geom"
)

var (
	watcher *fsnotify.Watcher
)

func init() {
	var err error
	watcher, err = fsnotify.NewWatcher()
	if err != nil {
		panic(err)
	}
	//watcher.Close()
}

func DebugWatch() {
	// Get list of sprites updated this frame
	var watcherSpritesToUpdate []string
FileWatchLoop:
	for {
		select {
		case event := <-watcher.Events:
			log.Println("event:", event)
			//if event.Op&fsnotify.Write == fsnotify.Write {
			//log.Println("modified sprite:", spriteName)
			//}

			spriteName := filepath.Base(filepath.Dir(event.Name))
			// Only reload sprite once
			for _, otherSpriteName := range watcherSpritesToUpdate {
				if otherSpriteName == spriteName {
					continue FileWatchLoop
				}
			}
			watcherSpritesToUpdate = append(watcherSpritesToUpdate, spriteName)
		case err := <-watcher.Errors:
			println("error:", err.Error())
		default:
			break FileWatchLoop
		}
	}

	// If those sprites are loaded, reload them
	manager := g_spriteManager
	for _, spriteName := range watcherSpritesToUpdate {
		spr := manager.assetMap[spriteName]
		if spr != nil {
			newSprData := loadSprite(spriteName)
			*spr = *newSprData
		}
	}
}

func debugWriteSprite(name string) {
	folderPath := file.AssetsDirectory + "/sprites/" + name + "/"

	// NOTE(Jake): 2018-06-18
	//
	// Watch for changes to the sprite (so we can reload it live!)
	//
	watcher.Remove(folderPath)
	watcher.Add(folderPath)

	// Load frames
	//
	// NOTE(Jake): 2018-03-12
	//
	// This is slow but makes managing assets simpler. Will most likely add something to pack
	// everything into a texture page for "production" builds.
	//
	frames := make([]spriteAssetFrame, 0, 10)
	for i := 0; true; i++ {
		path := folderPath + strconv.Itoa(i) + ".png"
		imageFileData, err := file.OpenFile(path)
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
		var buf bytes.Buffer
		err = png.Encode(&buf, image)
		if err != nil {
			panic(errors.New("Unable to encode image to bytes: " + path))
		}
		imageSize := image.Bounds().Size()
		frame := spriteAssetFrame{
			Size: geom.Vec{float64(imageSize.X), float64(imageSize.Y)},
			Data: buf.Bytes(),
		}
		frames = append(frames, frame)
	}

	// Read config information (if it exists)
	var config spriteConfig
	configPath := folderPath + "config.json"
	config = loadConfig(configPath)

	// Create sprite
	asset := newSpriteAsset(name, frames, config)

	// Write to file
	{
		spritePath := file.AssetsDirectory + "/sprites/" + name
		var data bytes.Buffer
		gob.NewEncoder(&data).Encode(asset)
		err := ioutil.WriteFile(spritePath+".data", data.Bytes(), 0644)
		if err != nil {
			panic(errors.New("Unable to write sprite out to file: " + spritePath))
		}
	}

}
