// This file contains scripts transcribed to raw data structures.
// I've retained the naming used in the Game Maker version and also retained some math
// used to calculate the positions so that I can compare versions easily into the future,
// just in-case.

package wall

// WallInfo is the data structure used to describe a spawned in piece of wall data
type WallInfo struct {
	WallList           []Wall
	TimeTillNext       int
	TimeTillNextRandom int
}

// Wall describes how a wall is spawned
type Wall struct {
	X        float64
	Y        float64
	IsInDirt bool // special wall that is shorter at the bottom to player doesn't collide with it
}

var wallLoose = WallInfo{
	WallList: []Wall{
		{Y: -302},
		{Y: 426},
	},
	TimeTillNext: 15,
}

var wallSemitight = WallInfo{
	WallList: []Wall{
		{Y: -272},
		{Y: 345},
	},
	TimeTillNext:       10,
	TimeTillNextRandom: 10,
}

var wallSemiloose = WallInfo{
	WallList: []Wall{
		{Y: 346},
	},
	TimeTillNext: 10,
}

var wallAbovesurface = WallInfo{
	WallList: []Wall{
		{
			Y:        68,
			IsInDirt: true,
		},
	},
	// so the timer doesn't wait long to spawn another thing
	TimeTillNext:       -5,
	TimeTillNextRandom: -10,
}

var wallTight = WallInfo{
	WallList: []Wall{
		{Y: -272},
		{Y: 320},
	},
	TimeTillNext:       10,
	TimeTillNextRandom: 10,
}

// Fly 1

var wallSemiloose1 = WallInfo{
	WallList: []Wall{
		{X: 48, Y: -310},
		{Y: 346},
		{X: 96, Y: 346},
	},
	TimeTillNext: 10,
}

var wallLooseFly1 = WallInfo{
	WallList: []Wall{
		{Y: 272},
		{Y: -325},
	},
	TimeTillNext:       10,
	TimeTillNextRandom: 10,
}

var wallMediumFly1 = WallInfo{
	WallList: []Wall{
		{Y: -300},
		{X: 48, Y: -285},
		{X: 96, Y: -350},
		// Repeat(3)
		{X: 48, Y: 320},
		{X: 96, Y: 320},
		{X: 144, Y: 320},
	},
	TimeTillNext:       10,
	TimeTillNextRandom: 10,
}

// Fly 2
