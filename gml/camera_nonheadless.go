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
	follow  ObjectType
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

	if inst := view.follow; inst != nil {
		roomInst := roomGetInstance(inst.BaseObject().RoomInstanceIndex())
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

func CameraSetViewTarget(index int, inst ObjectType) {
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

func cameraInstanceDestroy(inst ObjectType) {
	manager := gCameraManager
	for i := 0; i < len(manager.cameras); i++ {
		view := &manager.cameras[i]
		if view.follow == inst {
			view.follow = nil
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

	// Update player follow
	if view.follow != nil {
		inst := view.follow.BaseObject()
		if inst != nil {
			roomInst := roomGetInstance(inst.BaseObject().RoomInstanceIndex())
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
