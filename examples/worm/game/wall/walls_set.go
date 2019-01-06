package wall

func WallSets() [][]wallInfo {
	return wallSets
}

var wallSets = [][]wallInfo{
	wallSetFlat,
}

var wallSetFlat = []wallInfo{
	wallLoose,
	wallSemitight,
	wallSemiloose,
	wallAbovesurface,
}
