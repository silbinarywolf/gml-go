package game

import (
	"github.com/silbinarywolf/gml-go/gml"
)

const (
	FrontCityHspeed = -2
	BackCityHspeed  = -1
)

type Background struct {
	gml.Object
	IsPaused         bool
	BackCityOffset   float64
	FrontCityOffset  float64
	FrontGrassOffset float64
}

func (self *Background) Create() {
	self.SetDepth(DepthBackground)
}

func (self *Background) Update() {
	if Global.HasWormStopped() {
		return
	}

	// Update back city
	{
		size := gml.SpriteSize(SprBackCity)
		self.BackCityOffset += BackCityHspeed * gml.DeltaTime()
		if self.BackCityOffset < -size.X {
			self.BackCityOffset += size.X
		}
	}

	// Update front city
	{
		size := gml.SpriteSize(SprFrontCity)
		self.FrontCityOffset += FrontCityHspeed * gml.DeltaTime()
		if self.FrontCityOffset < -size.X {
			self.FrontCityOffset += size.X
		}
	}

	// Update grass
	{
		size := gml.SpriteSize(SprFrontGrass)
		self.FrontGrassOffset += -WallSpeed()
		if self.FrontGrassOffset < -size.X {
			self.FrontGrassOffset += size.X
		}
	}
}

func (self *Background) Draw() {
	roomSize := gml.RoomInstanceSize(self.RoomInstanceIndex())

	// Draw background
	gml.DrawSprite(SprSky, 0, 0, 0)

	// Draw back city
	{
		size := gml.SpriteSize(SprBackCity)
		gml.DrawSprite(SprBackCity, 0, self.BackCityOffset+15, 350)
		gml.DrawSprite(SprBackCity, 0, self.BackCityOffset+15+size.X, 350)
		gml.DrawSprite(SprBackCity, 0, self.BackCityOffset+15+(size.X*2), 350)
	}

	// Draw front city
	{
		size := gml.SpriteSize(SprFrontCity)
		gml.DrawSprite(SprFrontCity, 0, self.FrontCityOffset+15, 350)
		gml.DrawSprite(SprFrontCity, 0, self.FrontCityOffset+15+size.X, 350)
		gml.DrawSprite(SprFrontCity, 0, self.FrontCityOffset+15+(size.X*2), 350)
	}

	// Draw grass
	{
		size := gml.SpriteSize(SprFrontGrass)
		gml.DrawSprite(SprFrontGrass, 0, self.FrontGrassOffset, roomSize.Y-size.Y)
		gml.DrawSprite(SprFrontGrass, 0, self.FrontGrassOffset+size.X, roomSize.Y-size.Y)
		gml.DrawSprite(SprFrontGrass, 0, self.FrontGrassOffset+(size.X*2), roomSize.Y-size.Y)
		gml.DrawSprite(SprFrontGrass, 0, self.FrontGrassOffset+(size.X*3), roomSize.Y-size.Y)
	}
}
