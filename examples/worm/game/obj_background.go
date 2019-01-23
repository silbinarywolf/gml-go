package game

import (
	"github.com/silbinarywolf/gml-go/gml"
)

type Background struct {
	gml.Object
}

func (self *Background) Create() {
	self.SetDepth(DepthBackground)
}

func (self *Background) Draw() {
	roomSize := gml.RoomInstanceSize(self.RoomInstanceIndex())

	// Draw background
	gml.DrawSprite(SprSky, 0, 0, 0)

	// Draw back city
	{
		size := gml.SpriteSize(SprBackCity)
		gml.DrawSprite(SprBackCity, 0, 15, 350)
		gml.DrawSprite(SprBackCity, 0, 15+size.X, 350)
	}

	// Draw front city
	{
		size := gml.SpriteSize(SprFrontCity)
		gml.DrawSprite(SprFrontCity, 0, 15, 350)
		gml.DrawSprite(SprFrontCity, 0, 15+size.X, 350)
	}

	// Draw grass
	{
		size := gml.SpriteSize(SprFrontGrass)
		gml.DrawSprite(SprFrontGrass, 0, 0, roomSize.Y-size.Y)
		gml.DrawSprite(SprFrontGrass, 0, size.X, roomSize.Y-size.Y)
		gml.DrawSprite(SprFrontGrass, 0, size.X*2, roomSize.Y-size.Y)
	}

	gml.DrawTextF(32, 32, "%s", gml.FrameUsage())
}
