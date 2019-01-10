package wall

// WallInfo is the data structure used to describe a spawned in piece of wall data
type WallInfo struct {
	WallYList          []float64
	TimeTillNext       int
	TimeTillNextRandom int
}

var wallLoose = WallInfo{
	WallYList: []float64{
		-302,
		426,
	},
	TimeTillNext: 15,
}

var wallSemitight = WallInfo{
	WallYList: []float64{
		-272,
		345,
	},
	TimeTillNext:       10,
	TimeTillNextRandom: 10,
}

var wallSemiloose = WallInfo{
	WallYList: []float64{
		346,
	},
	TimeTillNext: 10,
}

var wallAbovesurface = WallInfo{
	WallYList: []float64{
		68,
	},
	// so the timer doesn't wait long to spawn another thing
	TimeTillNext:       -5,
	TimeTillNextRandom: -10,
}
