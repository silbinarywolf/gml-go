package wall

// WallInfo is the data structure used to describe a spawned in piece of wall data
type wallInfo struct {
	WallYList          []float64
	TimeTillNext       float64
	TimeTillNextRandom float64
}

var wallLoose = wallInfo{
	WallYList: []float64{
		-302,
		426,
	},
	TimeTillNext: 15,
}

var wallSemitight = wallInfo{
	WallYList: []float64{
		-272,
		345,
	},
	TimeTillNext:       10,
	TimeTillNextRandom: 10,
}

var wallSemiloose = wallInfo{
	WallYList: []float64{
		346,
	},
	TimeTillNext: 10,
}

var wallAbovesurface = wallInfo{
	WallYList: []float64{
		68,
	},
	// so the timer doesn't wait long to spawn another thing
	TimeTillNext:       -5,
	TimeTillNextRandom: -10,
}
