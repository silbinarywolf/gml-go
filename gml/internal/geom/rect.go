package geom

type Rect struct {
	Vec      // Position (contains X,Y)
	Size Vec // Size (X,Y)
}

func (rect *Rect) Pos() Vec {
	return rect.Vec
}

func (rect *Rect) Left() float64   { return rect.Vec.X }
func (rect *Rect) Right() float64  { return rect.Vec.X + rect.Size.X }
func (rect *Rect) Top() float64    { return rect.Vec.Y }
func (rect *Rect) Bottom() float64 { return rect.Vec.Y + rect.Size.Y }

func (rect *Rect) DistancePoint(point Vec) float64 {
	yDist := 0.0
	if point.Y < rect.Top() {
		yDist = rect.Top() - point.Y
	} else if point.Y > rect.Bottom() {
		yDist = point.Y - rect.Bottom()
	}
	xDist := 0.0
	if point.X < rect.Left() {
		xDist = rect.Left() - point.X
	} else if point.X > rect.Right() {
		xDist = point.X - rect.Right()
	}
	return Vec{xDist, yDist}.Norm()
}

// DistanceRect will give you the distance in pixels between two rectangles
// this is useful for seeing how far an object is from another object.
func (rect *Rect) DistanceRect(otherRect Rect) float64 {
	// source: https://stackoverflow.com/questions/4978323/how-to-calculate-distance-between-two-rectangles-context-a-game-in-lua
	left := otherRect.Right() < rect.Left()
	right := rect.Right() < otherRect.Left()
	bottom := otherRect.Bottom() < rect.Top()
	top := rect.Bottom() < otherRect.Top()
	if top && left {
		// dist((x1, y1b), (x2b, y2))
		return Vec{rect.Left(), rect.Bottom()}.DistancePoint(Vec{otherRect.Right(), otherRect.Top()})
	} else if left && bottom {
		// dist((x1, y1), (x2b, y2b))
		return Vec{rect.Left(), rect.Top()}.DistancePoint(Vec{otherRect.Right(), otherRect.Bottom()})
	} else if bottom && right {
		// dist((x1b, y1), (x2, y2b))
		return Vec{rect.Right(), rect.Top()}.DistancePoint(Vec{otherRect.Left(), otherRect.Bottom()})
	} else if right && top {
		// dist((x1b, y1b), (x2, y2))
		return Vec{rect.Right(), rect.Bottom()}.DistancePoint(Vec{otherRect.Left(), otherRect.Bottom()})
	} else if left {
		// x1 - x2b
		return rect.Left() - otherRect.Right()
	} else if right {
		// x2 - x1b
		return otherRect.Left() - rect.Right()
	} else if bottom {
		// y1 - y2b
		return rect.Top() - otherRect.Bottom()
	} else if top {
		// y2 - y1b
		return otherRect.Top() - rect.Bottom()
	}
	return 0
}

func (rect *Rect) CollisionPoint(pos Vec) bool {
	return pos.X > rect.Left() && pos.X < rect.Right() &&
		pos.Y > rect.Top() && pos.Y < rect.Bottom()
}

func (r1 Rect) CollisionRectangle(r2 Rect) bool {
	return r1.Right() > r2.Left() && r1.Bottom() > r2.Top() &&
		r1.Left() < r2.Right() && r1.Top() < r2.Bottom()
}
