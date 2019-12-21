// +build debug

package gml

import (
	"os"

	"github.com/silbinarywolf/gml-go/gml/internal/geom"
	"github.com/silbinarywolf/gml-go/gml/internal/sprite"
	"github.com/silbinarywolf/gml-go/gml/internal/user"
)

func debugConfigPath(name string) string {
	configPath := user.HomeDir() + "/.gmlgo"
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		os.Mkdir(configPath, 0700)
	}
	configPath = configPath + "/" + name + ".json"
	return configPath
}

func debugUpdate() {
	sprite.DebugWatch()

	if KeyboardCheck(VkControl) {
		if KeyboardCheckPressed(VkA) {
			if contextUpdate() != animationEditor {
				ContextUpdatePush(animationEditor)
			} else {
				ContextUpdatePop(animationEditor)

				// Reset camera
				CameraSetViewSize(0, WindowSize().X, WindowSize().Y)
				CameraSetViewTarget(0, Noone)
			}
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
	mousePos := MouseScreenPosition()
	left := pos.X
	right := left + float64(size.X)
	top := pos.Y
	bottom := top + float64(size.Y)
	return mousePos.X >= left && mousePos.X < right &&
		mousePos.Y >= top && mousePos.Y < bottom
}
