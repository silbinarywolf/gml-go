// +build headless

package sprite

import (
	"errors"
	"image"
	_ "image/png"
	"os"
	"path/filepath"
)

type SpriteFrame struct {
	width, height int
}

func (frame *SpriteFrame) Size() (width int, height int) { return frame.width, frame.height }

func createFrame(path string, i int) (SpriteFrame, error) {
	imageFileData, err := os.Open(filepath.FromSlash(path))
	if err != nil {
		return SpriteFrame{}, err
	}
	image, _, err := image.Decode(imageFileData)
	if err != nil {
		panic(errors.New("Unable to decode image: " + path))
	}
	width := image.Bounds().Size().X
	height := image.Bounds().Size().Y
	imageFileData.Close()
	return SpriteFrame{
		width:  width,
		height: height,
	}, nil
}

//func GetFrame(spr *Sprite, index int) *SpriteFrame {
//	return nil
//}
