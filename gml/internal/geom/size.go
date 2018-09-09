package geom

type Size struct {
	// NOTE(Jake): 2018-07-08
	//
	// When profiling the collision.go tests, I tried adjusting these types to:
	// - uint32, uint16, int16, int, float32
	//
	// When benchmarking with "go test --bench=." and "gjbt --bench=.", it seemed
	// that int32 gave the best performance
	//
	X, Y int32
}
