package gml

type Vec struct {
	X, Y float64
}

func V(x float64, y float64) Vec {
	return Vec{X: x, Y: y}
}
