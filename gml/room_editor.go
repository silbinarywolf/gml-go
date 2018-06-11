// +build !debug

package gml

import (
	"github.com/silbinarywolf/gml-go/gml/internal/object"
)

func roomEditorUsername() string {
	return ""
}

func EditorInit() {
}

func EditorIsInitialized() bool {
	return false
}

func EditorIsActive() bool {
	return false
}

func EditorSetRoom(room *Room) {
}

func EditorAddInstance(pos Vec, objectIndex object.ObjectIndex) *RoomObject {
	return nil
}

func EditorRemoveInstance(index int) {
}

func EditorUpdate() {
}

func EditorDraw() {
}

func EditorSave() {
}
