package sprite

import (
	"github.com/silbinarywolf/gml-go/gml/internal/geom"
)

type spriteAssetFrame struct {
	Size geom.Vec
	Data []byte
}

type spriteAsset struct {
	Name       string
	Size       geom.Vec
	ImageSpeed float64
	Frames     []spriteAssetFrame
}

func newSpriteAsset(name string, frames []spriteAssetFrame, config spriteConfig) *spriteAsset {
	spr := new(spriteAsset)
	spr.Name = name
	spr.Frames = frames
	spr.ImageSpeed = config.ImageSpeed

	var width, height float64
	for _, frame := range frames {
		frameWidth, frameHeight := frame.Size.X, frame.Size.Y
		if width < frameWidth {
			width = frameWidth
		}
		if height < frameHeight {
			height = frameHeight
		}
	}
	spr.Size = geom.Vec{
		X: width,
		Y: height,
	}
	return spr
}
