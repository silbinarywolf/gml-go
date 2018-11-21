// +build headless

package gml

import (
	"github.com/silbinarywolf/gml-go/gml/internal/geom"
	"github.com/silbinarywolf/gml-go/gml/internal/object"
)

type cameraManager struct {
}

type camera struct {
}

func newCameraState() *cameraManager {
	return nil
}

func (view *camera) Reset() {
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

func CameraSetViewPos(index int, pos geom.Vec) {
}

func CameraSetViewSize(index int, size geom.Vec) {
}

func CameraSetViewTarget(index int, inst object.ObjectType) {
}

func cameraClear(index int) {
}

func cameraDraw(index int) {
}

func cameraInstanceDestroy(inst object.ObjectType) {
}

func (view *camera) update() {
}
