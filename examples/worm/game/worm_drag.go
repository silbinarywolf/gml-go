package game

import (
	"github.com/silbinarywolf/gml-go/gml"
)

type WormDrag struct {
	DragTimer gml.Alarm
	YDrag     float64
}

func (drag *WormDrag) Update(self *gml.Object) {
	if drag.DragTimer.Update(2) {
		drag.YDrag = self.Y
	}
}

func HandleCollision(inst *Worm) {
	for _, id := range gml.CollisionRectList(inst, inst.Pos()) {
		_, ok := gml.InstanceGet(id).(*Wall)
		if !ok {
			continue
		}
		inst.Dead = true
		break
	}
}
