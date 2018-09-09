package geom

//
// NOTE(Jake): 2018-08-11
//
// This code is only used by the room_editor_debug.go.
// Need to look at removing it in favour of making rect.go nicer.
//

type RoomEditorDebugRect struct {
	LeftTop     Vec
	RightBottom Vec
}

func (rect *RoomEditorDebugRect) Left() float64   { return rect.LeftTop.X }
func (rect *RoomEditorDebugRect) Right() float64  { return rect.RightBottom.X }
func (rect *RoomEditorDebugRect) Top() float64    { return rect.LeftTop.Y }
func (rect *RoomEditorDebugRect) Bottom() float64 { return rect.RightBottom.Y }

func R(a Vec, b Vec) RoomEditorDebugRect {
	rect := RoomEditorDebugRect{}
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
