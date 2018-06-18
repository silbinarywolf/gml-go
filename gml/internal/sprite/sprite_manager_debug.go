// +build debug

package sprite

import (
	"bytes"
	"encoding/gob"
	"errors"
	"image"
	"image/png"
	"io/ioutil"
	"strconv"

	"github.com/silbinarywolf/gml-go/gml/internal/file"
	"github.com/silbinarywolf/gml-go/gml/internal/math"
)

func debugWriteSprite(name string) {
	// Load frames
	//
	// NOTE(Jake): 2018-03-12
	//
	// This is slow but makes managing assets simpler. Will most likely add something to pack
	// everything into a texture page for "production" builds.
	//
	folderPath := file.AssetsDirectory + "/sprites/" + name + "/"
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
			Size: math.V(float64(imageSize.X), float64(imageSize.Y)),
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
