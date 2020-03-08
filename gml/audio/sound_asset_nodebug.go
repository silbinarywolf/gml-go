// +build js !debug

package audio

import (
	"bytes"
	"encoding/gob"
	"errors"
	"io/ioutil"

	"github.com/silbinarywolf/gml-go/gml/internal/file"
)

func debugLoadAndWriteSoundAsset(name string) *soundAsset {
	// Load from file
	path := file.AssetDirectory + "/" + SoundDirectoryBase + "/" + name + ".data"
	fileData, err := file.OpenFile(path)
	if err != nil {
		panic(errors.New("Unable to find sound asset: " + path))
	}
	bytesData, err := ioutil.ReadAll(fileData)
	fileData.Close()
	if err != nil {
		panic(errors.New("Unable to read bytes from sound datafile: " + path))
	}
	buf := bytes.NewReader(bytesData)
	asset := new(soundAsset)
	gob.NewDecoder(buf).Decode(asset)
	return asset
}
