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
	self.SetDepth(DepthWormBody)

	self.SeperationWidth = self.Size.X/2 + self.Size.X/6
	self.YDrag = self.Y
}

func (self *WormBody) Update() {
	// Update immediately as body parts lag behind
	self.YDrag = self.Y

	//
	{
		seperationWidth := self.SeperationWidth
		switch parent := self.Parent.Get(); parent := parent.(type) {
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

	master := self.Master.Get().(*Worm)
	if !master.Dead {
		HandleCollisionForWormOrWormPart(&self.Object, master)
	}
}
