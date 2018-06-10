package gml

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"os/user"
	"path"
	"strconv"
	"strings"

	m "github.com/silbinarywolf/gml-go/gml/internal/math"
	"github.com/silbinarywolf/gml-go/gml/internal/object"
)

type roomEditor struct {
	initialized       bool
	username          string
	editingRoom       *Room
	objectIndexToData []object.ObjectType

	camPos       Vec
	lastMousePos Vec
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
		lastMousePos:      MousePosition(),
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

func roomEditorObjectIndexToData(objectIndex int32) object.ObjectType {
	index := object.ObjectIndex(objectIndex)
	return gRoomEditor.objectIndexToData[index]
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

func EditorAddInstance(pos Vec, objectIndex object.ObjectIndex) *RoomObject {
	room := roomEditorEditingRoom()
	if room == nil {
		return nil
	}
	count := room.UserEntityCount
	room.UserEntityCount++

	// Get unique username
	username := roomEditorUsername()

	//
	//inst := roomEditor.objectIndexToData[objectIndex]
	//baseObj := inst.BaseObject()
	roomObj := &RoomObject{
		Filename:    "entity_" + username + "_" + strconv.FormatInt(count, 10),
		ObjectIndex: int32(objectIndex),
		X:           int32(pos.X),
		Y:           int32(pos.Y),
	}
	room.Instances = append(room.Instances, roomObj)
	return roomObj
}

func EditorRemoveInstance(index int) {
	room := roomEditorEditingRoom()

	// Unordered delete instance
	entryBeingDeleted := room.Instances[index]
	lastEntry := room.Instances[len(room.Instances)-1]
	room.Instances[index] = lastEntry
	room.Instances = room.Instances[:len(room.Instances)-1]

	// Track deleted entities
	room.DeletedInstances = append(room.DeletedInstances, entryBeingDeleted)
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
		camPos := roomEditor.camPos
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
	if KeyboardCheck(VkControl) &&
		KeyboardCheckPressed(VkS) {
		EditorSave()
		println("Saved room:", room.Filepath)
	}

	{
		lastMousePos := roomEditor.lastMousePos
		roomEditor.lastMousePos = MousePosition()

		grid := V(32, 32)

		// Left click
		if MouseCheckButton(MbLeft) {
			// NOTE(Jake): 2018-06-10
			//
			// We need to handle mouse click between the last mouse position
			// and current mouse position so that there are no gaps when you're
			// dragging the mouse across long distances.
			//
			rect := m.R(MousePosition(), lastMousePos)
			for x := rect.Left(); x <= rect.Right(); x += grid.X {
				for y := rect.Top(); y <= rect.Bottom(); y += grid.Y {
					mousePos := V(x, y)

					// Check to make sure we aren't placing over the top
					// of an existing entity
					hasCollision := false
					for _, obj := range room.Instances {
						inst := roomEditorObjectIndexToData(obj.ObjectIndex)
						if inst == nil {
							continue
						}
						pos := V(float64(obj.X), float64(obj.Y))
						size := inst.BaseObject().Size
						left := pos.X
						right := left + size.X
						top := pos.Y
						bottom := top + size.Y
						hasCollision = hasCollision ||
							(mousePos.X >= left && mousePos.X < right &&
								mousePos.Y >= top && mousePos.Y < bottom)
					}

					//
					if !hasCollision {
						// Snap to grid
						mousePos.X = snapToGrid(mousePos.X, grid.X)
						mousePos.Y = snapToGrid(mousePos.Y, grid.Y)

						roomObj := EditorAddInstance(mousePos, 2)
						fmt.Printf("Create entity: %v\n", roomObj)
					}
				}
			}
		}

		// Holding Right click
		if MouseCheckButton(MbRight) {
			// NOTE(Jake): 2018-06-10
			//
			// We need to handle mouse click between the last mouse position
			// and current mouse position so that there are no gaps when you're
			// dragging the mouse across long distances.
			//
			rect := m.R(MousePosition(), lastMousePos)
			for x := rect.Left(); x <= rect.Right(); x += grid.X {
				for y := rect.Top(); y <= rect.Bottom(); y += grid.Y {
					mousePos := V(x, y)

					for i, obj := range room.Instances {
						inst := roomEditorObjectIndexToData(obj.ObjectIndex)
						if inst == nil {
							continue
						}
						pos := V(float64(obj.X), float64(obj.Y))
						size := inst.BaseObject().Size
						left := pos.X
						right := left + size.X
						top := pos.Y
						bottom := top + size.Y
						if mousePos.X >= left && mousePos.X < right &&
							mousePos.Y >= top && mousePos.Y < bottom {
							EditorRemoveInstance(i)
							fmt.Printf("Deleted entity: %s\n", obj.GetFilename())
						}
					}
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

	roomDirectory := room.Filepath

	// Write objects
	{
		for _, obj := range room.Instances {
			data := roomEditorObjectIndexToData(obj.ObjectIndex)
			fname := obj.Filename
			filepath := roomDirectory + "/" + fname + ".txt"

			file, err := os.Create(filepath)
			if err != nil {
				println("Error writing file:", err.Error())

				// todo(jake): 2018-06-10
				//
				// Error recovery here?
				//

				return
			}

			name := data.ObjectName()
			x := obj.X
			y := obj.Y

			//
			w := bufio.NewWriter(file)
			w.WriteString(name)
			w.WriteByte('\n')
			w.WriteString(strconv.Itoa(int(x)))
			w.WriteByte('\n')
			w.WriteString(strconv.Itoa(int(y)))
			w.WriteByte('\n')
			w.Flush()

			file.Close()
		}
	}

	// Delete objects
	deletedInstances := room.DeletedInstances
	// NOTE(Jake): 2018-06-10
	//
	// Clear out 'DeletedInstances' before we save
	// the map file as data.
	//
	// We don't want this information
	// serialized.
	//
	room.DeletedInstances = nil
	for _, obj := range deletedInstances {
		fname := obj.Filename
		filepath := roomDirectory + "/" + fname + ".txt"
		err := os.Remove(filepath)
		if err != nil {
			println("Error deleting instance:", err.Error())
			continue
		}
	}

	// Save data copy of file
	room.writeDataFile(room.Filepath)
}
