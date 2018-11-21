// +build !debug

package gml

import "github.com/silbinarywolf/gml-go/gml/internal/room"

func roomEditorUsername() string {
	return ""
}

func EditorInit(exitEditorFunc func(room *room.Room)) {
}

func EditorIsInitialized() bool {
	return false
}

func EditorSetRoom(room *Room) {
}

func editorUpdate() {
}
