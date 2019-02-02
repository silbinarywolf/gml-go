// +build debug

package audio

import (
	"errors"
	"io/ioutil"

	"github.com/silbinarywolf/gml-go/gml/internal/file"
)

func debugLoadAndWriteSoundAsset(name string) *soundAsset {
	var kind soundKind

	path := file.AssetDirectory + "/" + SoundDirectoryBase + "/" + name + "/sound.wav"
	fileData, err := file.OpenFile(path)
	if err != nil {
		// Fallback to MP3
		path := file.AssetDirectory + "/" + SoundDirectoryBase + "/" + name + "/sound.mp3"
		fileData, err = file.OpenFile(path)
		if err != nil {
			panic(err)
		}
		kind = soundKindMP3
	} else {
		kind = soundKindWAV
	}
	data, err := ioutil.ReadAll(fileData)
	if err != nil {
		panic(errors.New("Unable to read font file into bytes: " + path))
	}
	fileData.Close()

	return &soundAsset{
		kind: kind,
		data: data,
	}
}
