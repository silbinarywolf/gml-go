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
	ImageSpeed float64 `json:"ImageSpeed"`
}

func loadConfig(path string) spriteConfig {
	fileData, err := file.OpenFile(path)
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
