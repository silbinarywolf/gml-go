package gml

var (
	currentCamera *camera
	cameraList    [8]camera
)

type camera struct {
	enabled bool
	follow  ObjectType
	Vec
	size Vec
}

func CameraSetEnabled(index int) {
	view := &cameraList[index]
	view.enabled = true
}

func CameraSetViewPos(index int, pos Vec) {
	view := &cameraList[index]
	view.Vec = pos
}

func CameraSetViewSize(index int, size Vec) {
	view := &cameraList[index]
	view.size = size
}

func CameraSetViewTarget(index int, inst ObjectType) {
	view := &cameraList[index]
	view.follow = inst
}

func (view *camera) update() {
	if view.follow != nil {
		inst := currentCamera.follow.BaseObject()
		roomInst := inst.room.room
		left := float64(roomInst.Left)
		right := float64(roomInst.Right)
		top := float64(roomInst.Top)
		bottom := float64(roomInst.Bottom)

		view.X = inst.X - (view.size.X / 2)
		view.Y = inst.Y - (view.size.Y / 2)
		if view.X < left {
			view.X = left
		}
		if view.X+view.size.X > right {
			view.X = right - view.size.X
		}
		if view.Y < top {
			view.Y = top
		}
		if view.Y+view.size.Y > bottom {
			view.Y = bottom - view.size.Y
		}
	}
}
