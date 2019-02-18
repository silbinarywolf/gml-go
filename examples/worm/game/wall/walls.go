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
		{X: 0, Y: -300},
		{X: 48, Y: -285},
		{X: 96, Y: -350},
		// repeat(3)
		{X: 0, Y: 320},
		{X: 48, Y: 320},
		{X: 96, Y: 320},
	},
	TimeTillNext:       10,
	TimeTillNextRandom: 10,
}

// Fly 2

var wallTightFly1 = WallInfo{
	WallList: []Wall{
		{X: 0, Y: -300},
		{X: 48, Y: -275},
		{X: 96, Y: -350},
		{X: 144, Y: -350},
		// Repeat(4)
		{X: 0, Y: 320},
		{X: 48, Y: 320},
		{X: 96, Y: 320},
		{X: 144, Y: 320},
	},
	TimeTillNext:       10,
	TimeTillNextRandom: 10,
}

// Fly 3

var wallLooseFly3 = WallInfo{
	WallList: []Wall{
		{X: 0, Y: -400},
		{X: 0, Y: 242},
	},
	TimeTillNext:       10,
	TimeTillNextRandom: 10,
}

var wallLooseRow3 = WallInfo{
	WallList: []Wall{
		// repeat(12)
		{X: 0, Y: -300}, {X: 0, Y: 380},
		{X: 48, Y: -300}, {X: 48, Y: 380},
		{X: 48 * 2, Y: -300}, {X: 48 * 2, Y: 380},
		{X: 48 * 3, Y: -300}, {X: 48 * 3, Y: 380},
		{X: 48 * 4, Y: -300}, {X: 48 * 4, Y: 380},
		{X: 48 * 5, Y: -300}, {X: 48 * 5, Y: 380},
		{X: 48 * 6, Y: -300}, {X: 48 * 6, Y: 380},
		{X: 48 * 7, Y: -300}, {X: 48 * 7, Y: 380},
		{X: 48 * 8, Y: -300}, {X: 48 * 8, Y: 380},
		{X: 48 * 9, Y: -300}, {X: 48 * 9, Y: 380},
		{X: 48 * 10, Y: -300}, {X: 48 * 10, Y: 380},
		{X: 48 * 11, Y: -300}, {X: 48 * 11, Y: 380},
	},
	TimeTillNext:       90,
	TimeTillNextRandom: 10,
}

var wallTightFall2 = WallInfo{
	WallList: []Wall{
		// repeat(4)
		{X: 0, Y: 210},
		{X: 48, Y: 210 + 48},
		{X: 48 * 2, Y: 210 + (48 * 2)},
		{X: 48 * 3, Y: 210 + (48 * 3)},
	},
	TimeTillNext:       10,
	TimeTillNextRandom: 10,
}

// Fly 4

var wallLoose4 = WallInfo{
	WallList: []Wall{
		// repeat(2)
		{X: 0, Y: 200},
		{X: 48, Y: 200},
	},
	TimeTillNext:       30,
	TimeTillNextRandom: 10,
}
