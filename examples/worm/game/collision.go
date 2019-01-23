package game

import (
	"github.com/silbinarywolf/gml-go/gml"
)

func HandleCollisionForWormOrWormPart(self *gml.Object, master *Worm) {
	for _, id := range gml.CollisionRectList(self, self.X, self.Y) {
		inst, ok := gml.InstanceGet(id).(*Wall)
		if !ok {
			continue
		}
		if !master.InAir &&
			inst.DontKillPlayerIfInDirt {
			// Special case where wall is jutting into the ground
			// but not enough that the player should die.
			continue
		}
		master.TriggerDeath()
		break
	}
}
