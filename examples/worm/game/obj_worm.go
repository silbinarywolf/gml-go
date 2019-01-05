package game

import "github.com/silbinarywolf/gml-go/gml"

type Worm struct {
	gml.Object
	//enemyCreateAlarm gml.Alarm
}

func (self *Worm) Create() {
	self.SetSprite(SprWormHead)

	//roomSize := RoomInstanceSize(self.RoomInstanceIndex())
	self.X = 304
	self.Y = 528
}
