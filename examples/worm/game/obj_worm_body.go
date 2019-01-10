package game

import (
	"github.com/silbinarywolf/gml-go/gml"
)

const (
	SproutLerp      = 0.1
	SproutLerpSpeed = 0.1
)

type WormBody struct {
	gml.Object
	WormDrag

	Index           int
	Parent          gml.InstanceIndex
	Master          gml.InstanceIndex
	SeperationWidth float64
}

func (self *WormBody) Create() {
	self.SetSprite(SprWormBody)

	self.SeperationWidth = self.Size.X/2 + self.Size.X/6
	self.YDrag = self.Y
}

func (self *WormBody) Update() {
	self.WormDrag.Update(&self.Object)

	//
	{
		seperationWidth := self.SeperationWidth
		switch parent := gml.InstanceGet(self.Parent); parent := parent.(type) {
		case *Worm:
			self.X = parent.X - seperationWidth
			self.Y = parent.YDrag
		case *WormBody:
			self.X = parent.X - seperationWidth
			self.Y = parent.YDrag
		}
		//if parent := gml.InstanceGet(self.Parent); parent != nil {
		/*parent := parent.BaseObject()
		seperationWidth := self.SeperationWidth
		self.X = parent.X - seperationWidth
		self.Y = parent.YDrag*/
		/*if (instance_exists(parent)) {
		    var sep_width = seperation_width;
		    if (sprout)
		    {
		        sep_width *= sprout_lerp;
		        sprout_lerp += sprout_lerp_speed;
		        sprout_lerp_speed += 0.05;
		        if (sprout_lerp > 1.0)
		        {
		            sprout_lerp = 0.1;
		            sprout_lerp_speed = 0.1;
		            sprout = false;
		        }
		    }
		    x = parent.x - sep_width;
		    y = parent.ylag;
		}*/
		//}
	}

	master := gml.InstanceGet(self.Master).(*Worm)
	if !master.Dead {
		// Wall
		for _, id := range gml.CollisionRectList(self, self.Pos()) {
			_, ok := gml.InstanceGet(id).(*Wall)
			if !ok {
				continue
			}
			master.TriggerDeath()
			break
		}
	}
}
