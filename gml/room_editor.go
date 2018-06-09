package gml

import (
	"os/user"
	"path"
	"strconv"
	"strings"

	"github.com/silbinarywolf/gml-go/gml/internal/object"
)

type roomEditor struct {
	initialized       bool
	username          string
	editingRoom       *Room
	objectIndexToData []object.ObjectType

	camPos Vec
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
	idToEntityData := object.IDToEntityData()
	objectIndexToData := make([]object.ObjectType, len(idToEntityData))
	for i, obj := range idToEntityData {
		if obj == nil {
			continue
		}
		objectIndex := obj.ObjectIndex()
		inst := object.NewRawInstance(objectIndex, i, 0)
		inst.Create()
		objectIndexToData[i] = inst
	}

	return &roomEditor{
		initialized: true,
		username:    username,
		//editingRoom: nil,
		objectIndexToData: objectIndexToData,
	}
}

var (
	gRoomEditor *roomEditor
)

func roomEditorUsername() string {
	return gRoomEditor.username
}

func roomEditorEditingRoom() *Room {
	return gRoomEditor.editingRoom
}

func EditorInit() {
	if gRoomEditor != nil {
		panic("EditorInit: Room Editor is already initialized.")
	}
	gRoomEditor = newRoomEditor()
}

func EditorIsInitialized() bool {
	return gRoomEditor != nil
}

func EditorIsActive() bool {
	return gRoomEditor != nil && roomEditorEditingRoom() != nil
}

func EditorSetRoom(room *Room) {
	gRoomEditor.editingRoom = room

	//
	CameraSetEnabled(0)
	CameraSetViewSize(0, V(float64(windowWidth()), float64(windowHeight())))
}

func EditorAddInstance(pos Vec, objectIndex object.ObjectIndex) object.ObjectType {
	room := roomEditorEditingRoom()
	if room == nil {
		return nil
	}
	count := room.UserEntityCount
	room.UserEntityCount++

	// Get unique username
	username := roomEditorUsername()

	//
	inst := gRoomEditor.objectIndexToData[objectIndex]
	baseObj := inst.BaseObject()
	roomObj := &RoomObject{
		Filename:    "entity_" + username + "_" + strconv.FormatInt(count, 10),
		ObjectIndex: int32(objectIndex),
		X:           int32(baseObj.X),
		Y:           int32(baseObj.Y),
	}
	room.Instances = append(room.Instances, roomObj)
	return inst
}

func EditorUpdate() {
	room := roomEditorEditingRoom()
	if room == nil {
		return
	}

	{
		// Move camera
		camPos := gRoomEditor.camPos
		var speed float64 = 4
		if KeyboardCheck(VkShift) {
			speed = 8
		}
		if KeyboardCheck(VkRight) || KeyboardCheck(VkD) {
			camPos.X += speed
		} else if KeyboardCheck(VkLeft) || KeyboardCheck(VkA) {
			camPos.X -= speed
		}
		if KeyboardCheck(VkUp) || KeyboardCheck(VkW) {
			camPos.Y -= speed
		} else if KeyboardCheck(VkDown) || KeyboardCheck(VkS) {
			camPos.Y += speed
		}
		gRoomEditor.camPos = camPos
		CameraSetViewPos(0, camPos)
	}

	//
}

func EditorDraw() {
	room := roomEditorEditingRoom()
	if room == nil {
		return
	}

	// NOTE(Jake): 2018-06-04
	//
	// A hack to set the camera context, should probably
	// add an internal function for this
	//
	currentCamera = &cameraList[0]

	objectIndexToData := gRoomEditor.objectIndexToData
	instances := room.Instances
	for _, obj := range instances {
		objectIndex := obj.ObjectIndex
		inst := objectIndexToData[objectIndex]
		if inst == nil {
			continue
		}
		baseObj := inst.BaseObject()
		baseObj.X = float64(obj.X)
		baseObj.Y = float64(obj.Y)
		inst.Draw()
	}
}

func EditorSave() {
	room := roomEditorEditingRoom()
	if room == nil {
		return
	}
}
