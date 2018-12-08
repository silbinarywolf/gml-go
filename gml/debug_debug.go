// +build debug

package gml

import (
	"os"

	"github.com/silbinarywolf/gml-go/gml/internal/geom"
	"github.com/silbinarywolf/gml-go/gml/internal/sprite"
	"github.com/silbinarywolf/gml-go/gml/internal/user"
)

const (
	debugMode = true
)

var (
	debugMenuID = debugMenuNone
)

func debugConfigPath(name string) string {
	configPath := user.HomeDir() + "/.gmlgo"
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		os.Mkdir(configPath, 0700)
	}
	configPath = configPath + "/" + name + ".json"
	return configPath
}

func debugMenuOpenOrToggleClosed(id debugMenu) {
	if debugMenuID != id {
		debugMenuID = id
	} else {
		debugMenuID = debugMenuNone

		// Reset camera
		CameraSetViewSize(0, geom.Vec{float64(WindowWidth()), float64(WindowHeight())})
		CameraSetViewTarget(0, nil)
	}
}

func debugUpdate() {
	sprite.DebugWatch()

	if KeyboardCheck(VkControl) {
		if KeyboardCheckPressed(VkA) {
			debugMenuOpenOrToggleClosed(debugMenuAnimationEditor)
		}
		if KeyboardCheckPressed(VkR) {
			debugMenuOpenOrToggleClosed(debugMenuRoomEditor)
		}
	}
}

func debugDrawIsMouseOver(pos geom.Vec, size geom.Vec) bool {
	if DrawGetGUI() {
		return isMouseScreenOver(pos, size)
	} else {
		return isMouseOver(pos, size)
	}
}

func isMouseOver(pos geom.Vec, size geom.Vec) bool {
	mousePos := MousePosition()
	left := pos.X
	right := left + float64(size.X)
	top := pos.Y
	bottom := top + float64(size.Y)
	return mousePos.X >= left && mousePos.X < right &&
		mousePos.Y >= top && mousePos.Y < bottom
}

func isMouseScreenOver(pos geom.Vec, size geom.Vec) bool {
	mousePos := mouseScreenPosition()
	left := pos.X
	right := left + float64(size.X)
	top := pos.Y
	bottom := top + float64(size.Y)
	return mousePos.X >= left && mousePos.X < right &&
		mousePos.Y >= top && mousePos.Y < bottom
}
