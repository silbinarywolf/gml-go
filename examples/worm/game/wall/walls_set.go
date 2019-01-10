package wall

func WallSets() [][]WallInfo {
	return wallSets
}

var wallSets = [][]WallInfo{
	wallSetFlat,
}

var wallSetFlat = []WallInfo{
	wallLoose,
	wallSemitight,
	wallSemiloose,
	wallAbovesurface,
}
