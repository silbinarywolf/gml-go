package gml

import (
	"math"

	"github.com/silbinarywolf/gml-go/gml/internal/geom"
	"github.com/silbinarywolf/gml-go/gml/internal/object"
)

var (
	gCameraManager *cameraManager = newCameraState()
)

type cameraManager struct {
	cameras [8]camera
	current *camera
}

type camera struct {
	enabled bool
	follow  object.ObjectType
	geom.Vec
	size  geom.Vec
	scale geom.Vec
}

func newCameraState() *cameraManager {
	manager := new(cameraManager)
	for i := 0; i < len(manager.cameras); i++ {
		view := &manager.cameras[i]
		view.scale.X = 1
		view.scale.Y = 1
	}
	return manager
}

func (view *camera) Size() geom.Vec {
	return view.size
}

func (view *camera) Scale() geom.Vec {
	return view.scale
}

func CameraSetEnabled(index int) {
	view := &gCameraManager.cameras[index]
	view.enabled = true
}

func cameraGetActive() *camera {
	return gCameraManager.current
}

func cameraSetActive(index int) {
	gCameraManager.current = &gCameraManager.cameras[index]
}

func cameraClearActive() {
	gCameraManager.current = nil
}

func CameraGetViewPos(index int) geom.Vec {
	view := &gCameraManager.cameras[index]
	return view.Vec
}

func CameraSetViewPos(index int, pos geom.Vec) {
	view := &gCameraManager.cameras[index]
	view.Vec = pos

	if inst := view.follow; inst != nil {
		roomInst := RoomGetInstance(object.RoomInstanceIndex(inst.BaseObject()))
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
			view.X = math.Floor(view.X)
			view.Y = math.Floor(view.Y)
		}
	}
}

func CameraSetViewSize(index int, size geom.Vec) {
	view := &gCameraManager.cameras[index]
	view.size = size
}

func CameraSetViewTarget(index int, inst object.ObjectType) {
	view := &gCameraManager.cameras[index]
	view.follow = inst
}

func cameraInstanceDestroy(inst object.ObjectType) {
	manager := gCameraManager
	for i := 0; i < len(manager.cameras); i++ {
		view := &manager.cameras[i]
		if view.follow == inst {
			view.follow = nil
		}
	}
}

func (view *camera) update() {
	if view.follow != nil {
		//cam := cameraGetActive()
		inst := view.follow.BaseObject()
		if inst != nil {
			roomInst := RoomGetInstance(object.RoomInstanceIndex(inst))
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
				view.X = math.Floor(view.X)
				view.Y = math.Floor(view.Y)
			}
		}
	}
}
