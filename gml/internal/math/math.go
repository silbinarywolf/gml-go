package math

type Vec struct {
	X, Y float64
}

func V(x float64, y float64) Vec {
	return Vec{X: x, Y: y}
}

type Rect struct {
	LeftTop     Vec
	RightBottom Vec
}

func (rect *Rect) Left() float64   { return rect.LeftTop.X }
func (rect *Rect) Right() float64  { return rect.RightBottom.X }
func (rect *Rect) Top() float64    { return rect.LeftTop.Y }
func (rect *Rect) Bottom() float64 { return rect.RightBottom.Y }

func R(a Vec, b Vec) Rect {
	rect := Rect{}
	if a.X < b.X {
		rect.LeftTop.X = a.X
		rect.RightBottom.X = b.X
	} else {
		rect.LeftTop.X = b.X
		rect.RightBottom.X = a.X
	}
	if a.Y < b.Y {
		rect.LeftTop.Y = a.Y
		rect.RightBottom.Y = b.Y
	} else {
		rect.LeftTop.Y = b.Y
		rect.RightBottom.Y = a.Y
	}
	return rect
}
