package gml

import (
	"math"
	"strconv"

	"github.com/silbinarywolf/gml-go/gml/internal/geom"
)

var (
	gCameraManager *cameraManager = new(cameraManager)
)

type cameraManager struct {
	cameras             [8]camera
	current             *camera
	camerasEnabledCount int
}

type camera struct {
	cameraSurface
	enabled bool
	follow  InstanceIndex
	geom.Rect
	scale geom.Vec
	//updateFunc func()
}

func (manager *cameraManager) reset() {
	for i := 0; i < len(manager.cameras); i++ {
		view := &manager.cameras[i]
		*view = camera{}
		view.scale.X = 1
		view.scale.Y = 1
	}

	// Setup 1st camera
	CameraCreate(0, 0, 0, WindowWidth(), WindowHeight())
}

func CameraCreate(index int, windowX, windowY, windowWidth, windowHeight float64) {
	view := &gCameraManager.cameras[index]
	if view.enabled {
		panic("Camera " + strconv.Itoa(index) + " is already enabled.")
	}
	if windowWidth == 0 ||
		windowHeight == 0 {
		panic("Cannot have camera window width or height of 0")
	}
	view.X = windowX
	view.Y = windowY
	view.Size.X = windowWidth
	view.Size.Y = windowHeight
	view.enabled = true
	gCameraManager.camerasEnabledCount++
}

func CameraDestroy(index int) {
	view := &gCameraManager.cameras[index]
	if !view.enabled {
		panic("Camera " + strconv.Itoa(index) + " is not enabled.")
	}
	view.enabled = false
	gCameraManager.camerasEnabledCount--
}

func CameraSetSize(index int, windowWidth, windowHeight float64) {
	view := &gCameraManager.cameras[index]
	if !view.enabled {
		panic("Camera " + strconv.Itoa(index) + " is not enabled.")
	}
	view.Size.X = windowWidth
	view.Size.Y = windowHeight
}

// cameraUpdate wlll move the camera to center on the follow object and
// ensure it fits the room dimensions
func cameraUpdate() {
	for i, _ := range gCameraManager.cameras {
		view := &gCameraManager.cameras[i]
		if inst := view.follow.Get(); inst != nil {
			inst := inst.BaseObject()
			view.X = inst.X - (view.Size.X / 2)
			view.Y = inst.Y - (view.Size.Y / 2)
		}
		view.cameraFitToRoomDimensions()

		// todo: Jake: 2019-02-24 - https://github.com/silbinarywolf/gml-go/issues/103
		// Add this API when needed
		//if view.updateFunc != nil {
		//	view.updateFunc()
		//}
	}
}

// cameraGetActive gets the current camera we're drawing objects onto
func cameraGetActive() *camera {
	return gCameraManager.current
}

// cameraSetActive gets the current camera we want to draw objects onto
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

func CameraGetViewSize(index int) geom.Vec {
	view := &gCameraManager.cameras[index]
	return view.Size
}

func CameraSetViewPos(index int, x, y float64) {
	view := &gCameraManager.cameras[index]
	view.Vec = geom.Vec{
		X: x,
		Y: y,
	}

	view.cameraFitToRoomDimensions()
}

//func CameraSetUpdateFunction(index int, updateFunc func()) {
//	view := &gCameraManager.cameras[index]
//	view.updateFunc = updateFunc
//}

func CameraSetViewSize(index int, width, height float64) {
	view := &gCameraManager.cameras[index]
	view.Size = geom.Vec{
		X: width,
		Y: height,
	}
}

func CameraSetViewTarget(index int, inst InstanceIndex) {
	view := &gCameraManager.cameras[index]
	view.follow = inst
}

// cameraHasMultipleEnabled is generally used to disable
// rendering to an offscreen surface if using 1 camera.
func cameraHasMultipleEnabled() bool {
	return gCameraManager.camerasEnabledCount > 1
}

func cameraInstanceDestroy(instanceIndex InstanceIndex) {
	manager := gCameraManager
	for i := 0; i < len(manager.cameras); i++ {
		view := &manager.cameras[i]
		if view.follow == instanceIndex {
			view.follow = Noone
		}
	}
}

func (view *camera) cameraFitToRoomDimensions() {
	// If we're following an object, snap the camera to fit to the room
	if inst := view.follow.Get(); inst != nil {
		roomInst := roomGetInstance(inst.BaseObject().RoomInstanceIndex())
		if roomInst != nil {
			var left, right, top, bottom float64
			left = roomInst.Left()
			right = roomInst.Right()
			top = roomInst.Top()
			bottom = roomInst.Bottom()

			if view.X < left {
				view.X = left
			}
			if view.X+view.Size.X > right {
				view.X = right - view.Size.X
			}
			if view.Y < top {
				view.Y = top
			}
			if view.Y+view.Size.Y > bottom {
				view.Y = bottom - view.Size.Y
			}
			// NOTE(Jake): 2018-12-23
			// IIRC, Need to round these values otherwise draw calls show
			// gaps/artifacts.
			view.X = math.Floor(view.X)
			view.Y = math.Floor(view.Y)
		}
	}
}
