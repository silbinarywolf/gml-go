// +build !headless

package sprite

import (
	"errors"
	"image"

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

func GetFrame(spr *Sprite, index int) *ebiten.Image {
	return spr.frames[index].image
}
