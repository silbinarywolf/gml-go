package game

import (
	"math/rand"

	"github.com/silbinarywolf/gml-go/examples/worm/game/wall"
	"github.com/silbinarywolf/gml-go/gml"
)

type WallSpawner struct {
	SpawnWallTimer      gml.Alarm
	PreviousWallSpawned int
}

func (self *WallSpawner) Reset() {
	self.PreviousWallSpawned = -1
	self.SpawnWallTimer.Set(1)
}

func (self *WallSpawner) Update(roomInstanceIndex gml.RoomInstanceIndex) {
	if self.SpawnWallTimer.Tick() {
		// Get wall info
		var wallInfo wall.WallInfo
		{
			const WallX = 976
			wallSets := wall.WallSets()
			wallSet := wallSets[rand.Intn(len(wallSets))]

			// Select Wall randomly, make sure it isn't the same as before
			wallInfoIndex := rand.Intn(len(wallSet))
			if self.PreviousWallSpawned == wallInfoIndex {
				if wallInfoIndex > 0 {
					wallInfoIndex -= 1
				} else {
					wallInfoIndex += 1
				}
			}
			self.PreviousWallSpawned = wallInfoIndex
			wallInfo = wallSet[wallInfoIndex]
		}

		// Set timer
		ticksTillNextWall := 100 + rand.Intn(10)
		if wallInfo.TimeTillNext > 0 {
			ticksTillNextWall += wallInfo.TimeTillNext
			if wallInfo.TimeTillNextRandom < 0 {
				// Allow negative random numbers, reduce gap between spawning
				ticksTillNextWall -= rand.Intn(-wallInfo.TimeTillNextRandom)
			} else if wallInfo.TimeTillNextRandom > 0 {
				ticksTillNextWall += rand.Intn(wallInfo.TimeTillNextRandom)
			}
		}

		//if additional_spacing != WALL_TIME_EMPTY {
		//	timer += additional_spacing
		//}

		// Spawn wall
		{
			const WallX = 976
			wallSets := wall.WallSets()
			wallSet := wallSets[rand.Intn(len(wallSets))]

			// Select Wall randomly, make sure it isn't the same as before
			wallInfoIndex := rand.Intn(len(wallSet))
			if self.PreviousWallSpawned == wallInfoIndex {
				if wallInfoIndex > 0 {
					wallInfoIndex -= 1
				} else {
					wallInfoIndex += 1
				}
			}
			self.PreviousWallSpawned = wallInfoIndex

			wallInfo := wallSet[wallInfoIndex]
			for _, wall := range wallInfo.WallList {
				inst := gml.InstanceCreate(WallX, wall.Y, roomInstanceIndex, ObjWall).(*Wall)
				inst.DontKillPlayerIfInDirt = wall.IsInDirt
			}
		}

		// Add extra time once you've obtained the wings for the first time
		//if (!flight_extended_timer_used)
		//{
		//    additional_time += room_speed * 4;
		//    flight_extended_timer_used = true;
		//}
		self.SpawnWallTimer.Set(ticksTillNextWall)
	}
}
