// +build headless

package gml

import (
	"github.com/silbinarywolf/gml-go/gml/internal/geom"
)

type cameraManager struct {
}

type camera struct {
}

func newCameraState() *cameraManager {
	return nil
}

func (manager *cameraManager) reset() {
}

func CameraCreate(index int, windowX, windowY, windowWidth, windowHeight float64) {
}

func CameraSetSize(index int, windowWidth, windowHeight float64) {
}

// cameraGetActive gets the current camera we're drawing objects onto
func cameraGetActive() *camera {
	return nil
}

// cameraSetActive gets the current camera we want to draw objects onto
func cameraSetActive(index int) {
}

func cameraClearActive() {
}

func CameraGetViewPos(index int) geom.Vec {
	return geom.Vec{0, 0}
}

func CameraGetViewSize(index int) geom.Vec {
	return geom.Vec{0, 0}
}

func CameraSetViewPos(index int, x, y float64) {
}

func CameraSetViewSize(index int, width, height float64) {
}

func CameraSetViewTarget(index int, instanceIndex InstanceIndex) {
}

func cameraClear(index int) {
}

func cameraDraw(index int) {
}

func cameraInstanceDestroy(instanceIndex InstanceIndex) {
}

func (view *camera) update() {
}
