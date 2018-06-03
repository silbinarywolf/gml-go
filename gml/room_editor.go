package gml

import (
	"os/user"
	"path"
	"strconv"
	"strings"
)

type roomEditor struct {
	initialized       bool
	username          string
	objectIndexToData []ObjectType
}

func newRoomEditor() *roomEditor {
	// Set username
	user, _ := user.Current()
	username := user.Username
	username = path.Clean(username)
	username = strings.Replace(username, "/", "-", -1)
	username = strings.Replace(username, "\\", "-", -1)
	username = strings.Replace(username, "_", "-", -1)

	// Create stub instances to use for rendering map view
	objectIndexToData := make([]ObjectType, len(gObjectManager.idToEntityData))
	for i, obj := range gObjectManager.idToEntityData {
		objectIndex := obj.ObjectIndex()
		inst := newInstance(objectIndex)
		inst.Create()
		objectIndexToData[i] = inst
	}

	return &roomEditor{
		initialized:       true,
		username:          username,
		objectIndexToData: objectIndexToData,
	}
}

var (
	gRoomEditor *roomEditor
)

func roomEditorUsername() string {
	return gRoomEditor.username
}

func EditorInit() {
	if gRoomEditor != nil {
		panic("Room Editor is already initialized.")
	}
	gRoomEditor = newRoomEditor()
}

func EditorIsInitialized() bool {
	return gRoomEditor != nil
}

func (roomInst *RoomInstance) EditorAddInstance(pos Vec, objectIndex ObjectIndex) ObjectType {
	room := roomInst.room
	count := room.UserEntityCount
	room.UserEntityCount++

	// Get unique username
	username := roomEditorUsername()

	//
	inst := roomInst.InstanceCreate(pos, ObjectIndex(objectIndex))
	be := inst.BaseObject()
	roomObj := &RoomObject{
		Filename:    "entity_" + username + "_" + strconv.FormatInt(count, 10),
		ObjectIndex: int32(objectIndex),
		X:           int32(be.X),
		Y:           int32(be.Y),
	}
	room.Instances = append(room.Instances, roomObj)
	return inst
}

func (roomInst *RoomInstance) EditorUpdate() {

}

func (roomInst *RoomInstance) EditorDraw() {

}

func (roomInst *RoomInstance) EditorSave() {

}
