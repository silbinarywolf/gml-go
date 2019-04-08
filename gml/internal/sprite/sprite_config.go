package sprite

import (
	"encoding/json"
	"io/ioutil"

	"github.com/silbinarywolf/gml-go/gml/internal/file"
)

type spriteConfig struct {
	// NOTE(Jake): 2018-06-17
	//
	// Remember! JSON unmarshal won't work on
	// unexported fields!
	//
	ImageSpeed     float64                       `json:"ImageSpeed"`
	CollisionMasks map[int]map[int]CollisionMask `json:"CollisionMasks"`
}

func loadConfig(path string) spriteConfig {
	configPath := file.AssetDirectory + "/" + SpriteDirectoryBase + "/" + path + "/config.json"
	fileData, err := file.OpenFile(configPath)
	if err != nil {
		return spriteConfig{}
	}
	bytesData, err := ioutil.ReadAll(fileData)
	fileData.Close()

	var result spriteConfig
	err = json.Unmarshal(bytesData, &result)
	if err != nil {
		panic("Unmarshal error:" + err.Error())
	}
	return result
}
