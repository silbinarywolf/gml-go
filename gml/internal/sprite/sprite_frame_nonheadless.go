// +build !headless

package sprite

import (
	"bytes"
	"image"
	"image/draw"

	"github.com/hajimehoshi/ebiten"
)

type SpriteFrame struct {
	spriteFrameShared
	image *ebiten.Image
}

func (frame *SpriteFrame) Size() (width int, height int) { return frame.image.Size() }

func createFrame(frameData spriteAssetFrame) (SpriteFrame, error) {
	buf := bytes.NewReader(frameData.Data)
	imageData, _, err := image.Decode(buf)
	if err != nil {
		return SpriteFrame{}, err
	}

	// Crop transparency from image
	{
		imageSize := imageData.Bounds()
		top := 0
	Top:
		for y := 0; y < imageSize.Size().Y; y++ {
			for x := 0; x < imageSize.Size().X; x++ {
				r, g, b, a := imageData.At(x, y).RGBA()
				//fmt.Printf("r: %v, g: %v, b: %v, a: %v\n", r, g, b, a)
				if r == 0 && g == 0 && b == 0 && a == 0 {
					continue
				}
				top = y
				break Top
			}
		}
		left := 0
	Left:
		for x := 0; x < imageSize.Size().X; x++ {
			for y := 0; y < imageSize.Size().Y; y++ {
				r, g, b, a := imageData.At(x, y).RGBA()
				//fmt.Printf("r: %v, g: %v, b: %v, a: %v\n", r, g, b, a)
				if r == 0 && g == 0 && b == 0 && a == 0 {
					continue
				}
				left = x
				break Left
			}
		}
		right := 0
	Right:
		for x := imageSize.Size().Y - 1; x >= 0; x-- {
			for y := 0; y < imageSize.Size().Y; y++ {
				r, g, b, a := imageData.At(x, y).RGBA()
				//fmt.Printf("r: %v, g: %v, b: %v, a: %v\n", r, g, b, a)
				if r == 0 && g == 0 && b == 0 && a == 0 {
					continue
				}
				right = x
				break Right
			}
		}
		bottom := 0
	Bottom:
		for x := 0; x < imageSize.Size().X; x++ {
			for y := imageSize.Size().Y - 1; y >= 0; y-- {
				r, g, b, a := imageData.At(x, y).RGBA()
				//fmt.Printf("r: %v, g: %v, b: %v, a: %v\n", r, g, b, a)
				if r == 0 && g == 0 && b == 0 && a == 0 {
					continue
				}
				bottom = y
				break Bottom
			}
		}
		size := image.Rect(0, 0, right-left, bottom-top)
		croppedImageData := image.NewRGBA(size)
		draw.Draw(croppedImageData, size, imageData, image.Point{X: left, Y: top}, draw.Src)
		//fmt.Printf("left: %v, top: %v, right: %v, bottom: %v\n", left, top, right, bottom)
	}

	sheet, err := ebiten.NewImageFromImage(imageData, ebiten.FilterDefault)
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
