// +build !headless

package sprite

import (
	"errors"
	"image"
	"math"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type SpriteFrame struct {
	image *ebiten.Image
}

func (frame *SpriteFrame) Size() (width int, height int) { return frame.image.Size() }

func createFrame(path string, i int) (SpriteFrame, error) {
	imageFileData, err := ebitenutil.OpenFile(path)
	if err != nil {
		return SpriteFrame{}, err
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
	return SpriteFrame{
		image: sheet,
	}, nil
}

// NOTE(Jake): 2018-06-17
//
// This is called by draw_nonheadless.go in the parent package
// so that it can draw the image.
//
func GetRawFrame(spr *Sprite, index float64) *ebiten.Image {
	// NOTE(Jake): 2018-06-17
	//
	// Golang does not "cast", it uses type conversion, which means
	// a float64 -> int will *round* not simply *floor* as you might
	// expect in C/C++.
	//
	// https://stackoverflow.com/questions/35115868/how-to-round-to-nearest-int-when-casting-float-to-int-in-go
	//
	i := int(math.Floor(index))
	return spr.frames[i].image
}
