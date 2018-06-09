package gml

import (
	"fmt"
	"math"
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

func snapToGrid(val float64, grid float64) float64 {
	base := math.Floor(val / grid)
	return base * grid
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
	roomEditor := gRoomEditor
	count := room.UserEntityCount
	room.UserEntityCount++

	// Get unique username
	username := roomEditorUsername()

	//
	inst := roomEditor.objectIndexToData[objectIndex]
	//baseObj := inst.BaseObject()
	roomObj := &RoomObject{
		Filename:    "entity_" + username + "_" + strconv.FormatInt(count, 10),
		ObjectIndex: int32(objectIndex),
		X:           int32(pos.X),
		Y:           int32(pos.Y),
	}
	room.Instances = append(room.Instances, roomObj)
	return inst
}

func EditorUpdate() {
	room := roomEditorEditingRoom()
	if room == nil {
		return
	}
	roomEditor := gRoomEditor

	// NOTE(Jake): 2018-06-04
	//
	// A hack to set the camera context
	//
	cameraSetActive(0)
	defer cameraClearActive()

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

	{
		// Left click
		if MouseCheckPressed(MbLeft) {
			mousePos := MousePosition()
			mousePos.X = snapToGrid(mousePos.X, 32)
			mousePos.Y = snapToGrid(mousePos.Y, 32)
			fmt.Printf("Mouse X: %v, Y: %v \n", mousePos.X, mousePos.Y)
			inst := EditorAddInstance(mousePos, 2)
			fmt.Printf("%v\n", inst)
		}

		// Right click
		if MouseCheckPressed(MbRight) {
			mousePos := MousePosition()
			//mousePos.X = snapToGrid(mousePos.X, 32)
			//mousePos.Y = snapToGrid(mousePos.Y, 32)

			//fmt.Printf("Right Mouse X: %v, Y: %v \n", mousePos.X, mousePos.Y)

			objectIndexToData := roomEditor.objectIndexToData
			instances := room.Instances
			for i, obj := range instances {
				objectIndex := obj.ObjectIndex
				inst := objectIndexToData[objectIndex]
				if inst == nil {
					continue
				}
				baseObj := inst.BaseObject()
				pos := V(float64(obj.X), float64(obj.Y))
				size := baseObj.Size
				left := pos.X
				right := left + size.X
				top := pos.Y
				bottom := top + size.Y
				if mousePos.X >= left && mousePos.X < right &&
					mousePos.Y >= top && mousePos.Y < bottom {
					// Unordered delete instance
					lastEntry := instances[len(instances)-1]
					instances[i] = lastEntry
					instances = instances[:len(instances)-1]
				}
			}
		}
	}
}

func EditorDraw() {
	room := roomEditorEditingRoom()
	if room == nil {
		return
	}
	roomEditor := gRoomEditor

	// NOTE(Jake): 2018-06-04
	//
	// A hack to set the camera context
	//
	cameraSetActive(0)
	defer cameraClearActive()

	objectIndexToData := roomEditor.objectIndexToData
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
