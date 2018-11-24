// +build !headless

package sprite

import (
	"bytes"
	"image"

	"github.com/hajimehoshi/ebiten"
)

type SpriteFrame struct {
	spriteFrameShared
	image *ebiten.Image
}

func (frame *SpriteFrame) Size() (width int, height int) { return frame.image.Size() }

func createFrame(frameData spriteAssetFrame) (SpriteFrame, error) {
	buf := bytes.NewReader(frameData.Data)
	image, _, err := image.Decode(buf)
	if err != nil {
		return SpriteFrame{}, err
	}
	sheet, err := ebiten.NewImageFromImage(image, ebiten.FilterDefault)
	if err != nil {
		return SpriteFrame{}, err
	}
	r := SpriteFrame{
		image: sheet,
	}
	r.init(frameData)
	return r, nil
}

// NOTE(Jake): 2018-06-17
//
// This is called by draw_nonheadless.go in the parent package
// so that it can draw the image.
//
func GetRawFrame(spriteIndex SpriteIndex, index int) *ebiten.Image {
	// NOTE(Jake): 2018-06-17
	//
	// Golang does not "cast", it uses type conversion, which means
	// a float64 -> int will *round* not simply *floor* as you might
	// expect in C/C++.
	//
	// https://stackoverflow.com/questions/35115868/how-to-round-to-nearest-int-when-casting-float-to-int-in-go
	//
	return Frames(spriteIndex)[index].image
}
