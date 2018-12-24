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
	cameras             [8]camera
	current             *camera
	camerasEnabledCount int
}

type camera struct {
	enabled bool
	follow  InstanceIndex
	geom.Rect
	windowPos geom.Vec
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
	view.Size = WindowSize()
	view.scale.X = 1
	view.scale.Y = 1
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
	view.Size.X = windowWidth
	view.Size.Y = windowHeight
	view.enabled = true
	gCameraManager.camerasEnabledCount++
}

func CameraDestroy(index int) {
	view := &gCameraManager.cameras[index]
	if !view.enabled {
		panic("Camera " + strconv.Itoa(index) + " is not enabled.")
		return
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

func CameraSetViewPos(index int, x, y float64) {
	view := &gCameraManager.cameras[index]
	view.Vec = geom.Vec{
		X: x,
		Y: y,
	}

	view.cameraFitToRoomDimensions()
}

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
			if int(view.Size.X) != view.screen.Bounds().Max.X ||
				int(view.Size.Y) != view.screen.Bounds().Max.Y {
				mustCreateNewRenderTarget = true
			}
		}
		if mustCreateNewRenderTarget {
			image, err := ebiten.NewImage(int(view.Size.X), int(view.Size.Y), ebiten.FilterDefault)
			if err != nil {
				panic(err)
			}
			view.screen = image
		}
	}

	//
	if inst := InstanceGet(view.follow); inst != nil {
		inst := inst.BaseObject()
		view.X = inst.X - (view.Size.X / 2)
		view.Y = inst.Y - (view.Size.Y / 2)
	}
	view.cameraFitToRoomDimensions()
}

func (view *camera) cameraFitToRoomDimensions() {
	// If we're following an object, snap the camera to fit to the room
	if inst := InstanceGet(view.follow); inst != nil {
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
