package game

import (
	"math/rand"

	"github.com/silbinarywolf/gml-go/example/worm/game/wall"
	"github.com/silbinarywolf/gml-go/gml"
	"github.com/silbinarywolf/gml-go/gml/alarm"
)

type WallSpawner struct {
	WallList            []wall.WallInfo
	SpawnWallTimer      alarm.Alarm
	PreviousWallSpawned int
}

func (self *WallSpawner) Reset() {
	self.PreviousWallSpawned = -1
	self.SpawnWallTimer.Set(DesignedMaxTPS * 1.25)

	// Reset
	if self.WallList == nil {
		self.WallList = make([]wall.WallInfo, 0, 50)
	} else {
		self.WallList = self.WallList[:0]
	}
	self.WallList = append(self.WallList, wall.WallSetFlat...)
}

func (self *WallSpawner) Update(roomInstanceIndex gml.RoomInstanceIndex) {
	if self.SpawnWallTimer.Tick() {
		//if additional_spacing != WALL_TIME_EMPTY {
		//	timer += additional_spacing
		//}

		// Spawn wall
		const WallX = 976

		// Select Wall randomly, make sure it isn't the same as before
		wallInfoIndex := rand.Intn(len(self.WallList))
		if self.PreviousWallSpawned == wallInfoIndex {
			wallInfoIndex -= rand.Intn(1) - 1
			if wallInfoIndex < 0 {
				wallInfoIndex = len(self.WallList) - 1
			}
			if wallInfoIndex >= len(self.WallList) {
				wallInfoIndex = 0
			}
		}
		self.PreviousWallSpawned = wallInfoIndex
		//wallInfoIndex = 0

		wallInfo := &self.WallList[wallInfoIndex]
		roomInstanceIndex.InstanceCreate(WallX, 0, ObjCheckpoint)
		for _, wall := range wallInfo.WallList {
			inst := roomInstanceIndex.InstanceCreate(WallX+wall.X, wall.Y, ObjWall).(*Wall)
			inst.DontKillPlayerIfInDirt = wall.IsInDirt
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

		// Add extra time once you've obtained the wings for the first time
		//if (!flight_extended_timer_used)
		//{
		//    additional_time += room_speed * 4;
		//    flight_extended_timer_used = true;
		//}
		self.SpawnWallTimer.Set(float64(ticksTillNextWall))
	}
}
