package geom

import "math"

type Vec struct {
	X float64 `json:"X"`
	Y float64 `json:"Y"`
}

func (v Vec) Sub(ov Vec) Vec { return Vec{v.X - ov.X, v.Y - ov.Y} }

// Norm returns the vector's norm.
func (v Vec) Norm() float64 { return math.Sqrt(v.Dot(v)) }

func (v Vec) Dot(ov Vec) float64 { return v.X*ov.X + v.Y*ov.Y }

func (v Vec) DistancePoint(ov Vec) float64 { return v.Sub(ov).Norm() }
