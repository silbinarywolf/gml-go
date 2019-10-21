package wall

/*var wallSets = [][]WallInfo{
	wallSetFlat,
}*/

var WallSetFlat = []WallInfo{
	wallLoose,
	wallSemitight,
	wallSemiloose,
	wallAbovesurface,
}

var WallSetFlatHard = []WallInfo{
	wallTight,
	wallAbovesurface,
}

var WallSetFly1 = []WallInfo{
	wallLooseFly1,
	wallSemiloose1,
	wallMediumFly1,
}

var WallSetFly2 = []WallInfo{
	wallTightFly1,
}

var WallSetFly3 = []WallInfo{
	wallLooseFly3,
	wallLooseRow3,
	wallTightFall2,
}

var WallSetFly4 = []WallInfo{
	wallLoose4,
}

var WallSetFly5 = []WallInfo{
	// no data for this in original Game Maker game
}
