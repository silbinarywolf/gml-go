// +build !headless

package gml

import (
	"math"
	"strconv"

	"github.com/hajimehoshi/ebiten"
	"github.com/silbinarywolf/gml-go/gml/internal/geom"
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
	follow  InstanceIndex
	geom.Vec
	windowPos geom.Vec
	size      geom.Vec
	scale     geom.Vec
	screen    *ebiten.Image
}

func newCameraState() *cameraManager {
	manager := new(cameraManager)
	for i := 0; i < len(manager.cameras); i++ {
		view := &manager.cameras[i]
		view.Reset()
	}
	return manager
}

func (view *camera) Reset() {
	view.size.X = float64(WindowWidth())
	view.size.Y = float64(WindowHeight())
	view.scale.X = 1
	view.scale.Y = 1
}

//func (view *camera) Size() geom.Vec {
//	return view.size
//}

func (view *camera) Scale() geom.Vec {
	return view.scale
}

func CameraCreate(index int, windowX, windowY, windowWidth, windowHeight float64) {
	view := &gCameraManager.cameras[index]
	if view.enabled {
		panic("Camera " + strconv.Itoa(index) + " is already enabled.")
		return
	}
	if windowWidth == 0 ||
		windowHeight == 0 {
		panic("Cannot have camera window width or height of 0")
	}
	view.windowPos.X = windowX
	view.windowPos.Y = windowY
	view.size.X = windowWidth
	view.size.Y = windowHeight
	view.enabled = true
}

func CameraSetSize(index int, windowWidth, windowHeight float64) {
	view := &gCameraManager.cameras[index]
	if !view.enabled {
		panic("Camera " + strconv.Itoa(index) + " is not enabled.")
	}
	view.size.X = windowWidth
	view.size.Y = windowHeight
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

func CameraSetViewPos(index int, pos geom.Vec) {
	view := &gCameraManager.cameras[index]
	view.Vec = pos

	view.cameraFitToRoomDimensions()
}

func CameraSetViewSize(index int, size geom.Vec) {
	view := &gCameraManager.cameras[index]
	view.size = size
}

func CameraSetViewTarget(index int, inst InstanceIndex) {
	view := &gCameraManager.cameras[index]
	view.follow = inst
}

func cameraClear(index int) {
	view := &gCameraManager.cameras[index]
	view.screen.Clear()
}

func cameraDraw(index int) {
	view := &gCameraManager.cameras[index]
	op := ebiten.DrawImageOptions{}
	op.GeoM.Scale(view.scale.X, view.scale.Y)
	op.GeoM.Translate(view.windowPos.X, view.windowPos.Y)
	gScreen.DrawImage(view.screen, &op)
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

func (view *camera) update() {
	// Update screen render target
	{
		mustCreateNewRenderTarget := false
		if view.screen == nil {
			// Create new camera
			mustCreateNewRenderTarget = true
		} else {
			// Resize camera
			if int(view.size.X) != view.screen.Bounds().Max.X ||
				int(view.size.Y) != view.screen.Bounds().Max.Y {
				mustCreateNewRenderTarget = true
			}
		}
		if mustCreateNewRenderTarget {
			image, err := ebiten.NewImage(int(view.size.X), int(view.size.Y), ebiten.FilterDefault)
			if err != nil {
				panic(err)
			}
			view.screen = image
		}
	}

	//
	if inst := InstanceGet(view.follow); inst != nil {
		inst := inst.BaseObject()
		view.X = inst.X - (view.size.X / 2)
		view.Y = inst.Y - (view.size.Y / 2)
	}
	view.cameraFitToRoomDimensions()
}

func (view *camera) cameraFitToRoomDimensions() {
	// If we're following an object, snap the camera to fit to the room
	if inst := InstanceGet(view.follow); inst != nil {
		roomInst := roomGetInstance(inst.BaseObject().RoomInstanceIndex())
		if roomInst != nil {
			var left, right, top, bottom float64
			left = 0                         // float64(room.Left)
			right = float64(roomInst.size.X) // float64(room.Right)
			top = 0                          // float64(room.Top)
			bottom = float64(roomInst.size.Y)

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
			// NOTE(Jake): 2018-12-23
			// IIRC, Need to round these values otherwise draw calls show
			// gaps/artifacts.
			view.X = math.Floor(view.X)
			view.Y = math.Floor(view.Y)
		}
	}
}
