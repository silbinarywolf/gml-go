// +build debug

package gml

import (
	"bufio"
	"image/color"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/silbinarywolf/gml-go/gml/internal/file"
	m "github.com/silbinarywolf/gml-go/gml/internal/math"
	"github.com/silbinarywolf/gml-go/gml/internal/object"
)

type roomEditor struct {
	initialized       bool
	editingRoom       *Room
	objectIndexToData []object.ObjectType

	camPos       Vec
	lastMousePos Vec

	isEntityMenuOpen   bool
	entityMenuFiltered []object.ObjectType
	objectSelected     object.ObjectType
}

func newRoomEditor() *roomEditor {
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
		//editingRoom: nil,
		objectIndexToData:  objectIndexToData,
		lastMousePos:       MousePosition(),
		entityMenuFiltered: make([]object.ObjectType, 0, len(objectIndexToData)),
	}
}

var (
	gRoomEditor *roomEditor
)

func roomEditorUsername() string {
	return file.DebugUsernameFileSafe()
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
	isMenuOpen := roomEditor.isEntityMenuOpen

	// NOTE(Jake): 2018-06-04
	//
	// A hack to set the camera context
	//
	cameraSetActive(0)
	defer cameraClearActive()

	isHoldingControl := KeyboardCheck(VkControl)
	if isHoldingControl {
		// Open entity select menu
		if KeyboardCheckPressed(VkP) {
			roomEditor.isEntityMenuOpen = !roomEditor.isEntityMenuOpen
			if roomEditor.isEntityMenuOpen {
				ClearKeyboardString()
			}
		}

		// Save
		if KeyboardCheckPressed(VkS) {
			EditorSave()
			println("Saved room:", room.Filepath)
		}
	}

	// Close open menus
	if KeyboardCheckPressed(VkEscape) {
		roomEditor.isEntityMenuOpen = false
	}

	// Handle filtering / selection
	if roomEditor.isEntityMenuOpen {
		typingText := KeyboardString()
		roomEditor.entityMenuFiltered = roomEditor.entityMenuFiltered[:0]
		for _, obj := range roomEditor.objectIndexToData {
			if obj == nil {
				continue
			}
			name := obj.ObjectName()
			hasMatch := typingText == "" ||
				strings.Contains(name, typingText)
			if !hasMatch {
				continue
			}
			obj.BaseObject().ImageUpdate()
			roomEditor.entityMenuFiltered = append(roomEditor.entityMenuFiltered, obj)
		}

		//
		if KeyboardCheckPressed(VkEnter) &&
			len(roomEditor.entityMenuFiltered) > 0 {
			selectedObj := roomEditor.entityMenuFiltered[0]
			roomEditor.objectSelected = selectedObj
			roomEditor.isEntityMenuOpen = false
		}
	}

	if !isMenuOpen &&
		!isHoldingControl {
		{
			// Move camera
			camPos := &roomEditor.camPos
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
			CameraSetViewPos(0, *camPos)
		}

		lastMousePos := roomEditor.lastMousePos
		roomEditor.lastMousePos = MousePosition()

		grid := V(32, 32)

		// Left click
		if MouseCheckButton(MbLeft) &&
			roomEditor.objectSelected != nil {
			objectIndexSelected := roomEditor.objectSelected.ObjectIndex()
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

						roomObj := EditorAddInstance(mousePos, objectIndexSelected)
						println("Create entity:", roomObj.String())
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
							println("Deleted entity:", obj.Filename)
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
	isMenuOpen := roomEditor.isEntityMenuOpen

	// NOTE(Jake): 2018-06-04
	//
	// A hack to set the camera context
	//
	cameraSetActive(0)
	defer cameraClearActive()
	currentCamera := cameraGetActive()

	instances := room.Instances
	for _, obj := range instances {
		inst := roomEditorObjectIndexToData(obj.ObjectIndex)
		if inst == nil {
			continue
		}
		baseObj := inst.BaseObject()
		baseObj.X = float64(obj.X)
		baseObj.Y = float64(obj.Y)
		inst.Draw()
	}
	if !isMenuOpen {
		grid := V(32, 32)

		// Draw selected
		if selectedObj := roomEditor.objectSelected; selectedObj != nil {
			mousePos := MousePosition()
			mousePos.X = snapToGrid(mousePos.X, grid.X)
			mousePos.Y = snapToGrid(mousePos.Y, grid.Y)

			// Draw
			baseObj := selectedObj.BaseObject()
			baseObj.Vec = mousePos
			selectedObj.Draw()
		}
	}

	// Entity select menu
	if roomEditor.isEntityMenuOpen {
		{
			//screen := gScreen
			//screen.Fill(color.NRGBA{0x00, 0x00, 0x00, 128})
			DrawRectangle(m.V(0, 0), m.V(2048, 2048), color.RGBA{0, 0, 0, 190})
		}
		var x float64 = 128
		var y float64 = 32
		{
			searchText := "Search for object (type + press enter)"
			DrawText(m.V(128-(StringWidth(searchText)/4), y), searchText)
			y += 24
		}
		{
			typingText := KeyboardString()
			DrawText(m.V(x, y), typingText)
			DrawText(m.V(x+StringWidth(typingText), y), "|")
			y += 24
		}
		previewSize := m.V(32, 32)
		for _, obj := range roomEditor.entityMenuFiltered {
			//
			baseObj := obj.BaseObject()
			pos := &baseObj.Vec
			size := baseObj.Size
			oldImageScale := baseObj.ImageScale
			{
				pos.X = x - 40 + currentCamera.X
				pos.Y = y - (previewSize.Y / 2) + currentCamera.Y
				baseObj.ImageScale.X = previewSize.X / size.X
				baseObj.ImageScale.Y = previewSize.Y / size.Y
				obj.Draw()
				DrawText(m.V(x, y), obj.ObjectName())
			}
			baseObj.ImageScale = oldImageScale
			y += 48
		}
	}
}

func EditorSave() {
	room := roomEditorEditingRoom()
	if room == nil {
		return
	}

	roomDirectory := room.Filepath

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

				continue
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
			// NOTE(Jake): 2018-06-11
			//
			// Why are we writing the relative filename [x] times here?
			//
			// To stop git from confusing deleted entities with renaming/moving
			// a file. (ie. you'd delete an entity, create a new entity, but git would handle
			// it as a rename/move)
			//
			// I'm actually not sure if a Git rename/move could lead to merge conflicts...
			// but I'm going to err on the side of caution here until I decide to investigate
			// this further.
			//
			w.WriteString(fname)
			w.WriteString(fname)
			w.WriteString(fname)
			w.WriteString(fname)
			w.Flush()

			file.Close()
		}
	}

	// Save data copy of file
	room.DebugWriteDataFile(room.Filepath)
}
