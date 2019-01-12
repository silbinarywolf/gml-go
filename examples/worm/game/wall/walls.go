package wall

// WallInfo is the data structure used to describe a spawned in piece of wall data
type WallInfo struct {
	WallList           []Wall
	TimeTillNext       int
	TimeTillNextRandom int
}

// Wall describes how a wall is spawned
type Wall struct {
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
