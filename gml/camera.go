package gml

import (
	"github.com/silbinarywolf/gml-go/gml/internal/object"
)

var (
	__currentCamera *camera
	cameraList      [8]camera
)

type camera struct {
	enabled bool
	follow  object.ObjectType
	Vec
	size Vec
}

func CameraSetEnabled(index int) {
	view := &cameraList[index]
	view.enabled = true
}

func cameraGetActive() *camera {
	return __currentCamera
}

func cameraSetActive(index int) {
	__currentCamera = &cameraList[index]
}

func cameraClearActive() {
	__currentCamera = nil
}

func CameraSetViewPos(index int, pos Vec) {
	view := &cameraList[index]
	view.Vec = pos

	if inst := view.follow; inst != nil {
		roomInst := RoomGetInstance(inst.BaseObject().RoomInstanceIndex())
		if roomInst != nil {
			room := roomInst.room
			left := float64(room.Left)
			right := float64(room.Right)
			top := float64(room.Top)
			bottom := float64(room.Bottom)

			view.X = pos.X - (view.size.X / 2)
			view.Y = pos.Y - (view.size.Y / 2)
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
}

func CameraSetViewSize(index int, size Vec) {
	view := &cameraList[index]
	view.size = size
}

func CameraSetViewTarget(index int, inst object.ObjectType) {
	view := &cameraList[index]
	view.follow = inst
}

func (view *camera) update() {
	if view.follow != nil {
		//cam := cameraGetActive()
		inst := view.follow.BaseObject()

		roomInst := RoomGetInstance(inst.RoomInstanceIndex())
		if roomInst != nil {
			room := roomInst.room
			left := float64(room.Left)
			right := float64(room.Right)
			top := float64(room.Top)
			bottom := float64(room.Bottom)

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
}
