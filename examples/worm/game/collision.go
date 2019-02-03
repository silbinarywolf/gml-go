package game

import (
	"github.com/silbinarywolf/gml-go/gml"
)

func HandleCollisionForWormOrWormPart(self *gml.Object, master *Worm) {
	for _, id := range gml.CollisionRectList(self, self.X, self.Y) {
		inst, ok := id.Get().(*Wall)
		if !ok {
			continue
		}
		if !master.InAir &&
			inst.DontKillPlayerIfInDirt {
			// Special case where wall is jutting into the ground
			// but not enough that the player should die.
			continue
		}
		if inst.DontKillPlayer {
			// Special case for when you reset the game, walls that
			// existed from the previous game will still render on-screen
			// but they won't kill you
			continue
		}
		master.TriggerDeath()
		break
	}
}
