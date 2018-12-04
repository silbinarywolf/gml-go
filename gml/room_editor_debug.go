// +build debug

package gml

import (
	"bufio"
	"encoding/json"
	"fmt"
	"image/color"
	"io/ioutil"
	"math"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/silbinarywolf/gml-go/gml/internal/file"
	"github.com/silbinarywolf/gml-go/gml/internal/geom"
	"github.com/silbinarywolf/gml-go/gml/internal/object"
	"github.com/silbinarywolf/gml-go/gml/internal/reditor"
	"github.com/silbinarywolf/gml-go/gml/internal/room"
	"github.com/silbinarywolf/gml-go/gml/internal/sprite"
)

//
// NOTE(Jake): 2018-07-10
//
// I'd like to split this out into its own package but before I can
// do that I need to figure out:
// - How I'll access Mouse / Keyboard inputs
// - Split out camera.go into an internal package
//
//

//type roomInfo struct {
//	*room.Room
//}

type roomEditor struct {
	spriteViewer debugSpriteViewer

	initialized  bool
	editingRoom  *room.Room
	editingLayer room.RoomLayer

	objectIndexToData []object.ObjectType
	//spriteList        []*sprite.Sprite
	//spriteMap         map[string]*sprite.Sprite

	camPos             Vec
	lastMousePos       Vec
	lastMouseScreenPos Vec

	menuOpened        reditor.Menu
	menuLayerKind     room.RoomLayerKind
	hasUnsavedChanges bool

	entityMenuFiltered []object.ObjectType
	//spriteMenuFiltered []*sprite.Sprite

	objectSelected object.ObjectType
	spriteSelected SpriteIndex

	mouseHold   [MbSize]bool
	gridEnabled bool

	cameraStateBeforeEnteringEditingMode cameraManager

	// Callbacks
	//exitEditorFunc func(room *room.Room)

	//
	statusText  string
	statusTimer time.Time

	// Constants
	roomDirectory string

	// Backing / temporary pools
	tempLayers []room.RoomLayer
}

type roomEditorConfig struct {
	RoomSelected  string `json:"RoomSelected,omitempty"`
	LayerSelected string `json:"LayerSelected,omitempty"`
	BrushSelected string `json:"BrushSelected,omitempty"`
}

var (
	gRoomEditor *roomEditor
)

func newRoomEditor() *roomEditor {
	// NOTE(Jake): 2018-07-11
	//
	// Create stub instances to use for rendering map view.
	//
	// This provides us:
	// - The entity size (as set in Create())
	// - The default sprite of the object
	//
	objectIndexList := object.ObjectIndexList()
	objectIndexToData := make([]object.ObjectType, len(objectIndexList))
	for i, objectIndex := range objectIndexList {
		inst := object.NewRawInstance(objectIndex, i, 0, 0)
		inst.Create()
		objectIndexToData[i] = inst
	}

	return &roomEditor{
		initialized:       true,
		objectIndexToData: objectIndexToData,
		//objectNameToData:   objectNameToData,
		//spriteList:         spriteList,
		//spriteMap:          spriteMap,
		lastMousePos:       MousePosition(),
		entityMenuFiltered: make([]object.ObjectType, 0, len(objectIndexToData)),
		//spriteMenuFiltered: make([]*sprite.Sprite, 0, len(spriteList)),
		roomDirectory: file.AssetDirectory + "/" + room.RoomDirectoryBase + "/",
		tempLayers:    make([]room.RoomLayer, 0, 25),
		gridEnabled:   false,
	}
}

func (editor *roomEditor) IsMenuOpen() bool {
	return editor.menuOpened != reditor.MenuNone
}

func (editor *roomEditor) calculateAndSortLayers() {
	if editor.editingRoom == nil {
		return
	}
	editor.tempLayers = editor.tempLayers[:0]
	editor.layers()
}

func (editor *roomEditor) MouseClearButton(mb int) {
	editor.mouseHold[mb] = false
}

func (editor *roomEditor) MouseCheckButton(mb int) bool {
	return editor.mouseHold[mb]
}

func (editor *roomEditor) layers() []room.RoomLayer {
	// NOTE(Jake): 2018-08-05
	//
	// At the start of each EditorUpdate() / frame, we reset
	// the tempLayers length to 0 so this will only be calculated once
	// per frame.
	//
	if len(editor.tempLayers) > 0 {
		return editor.tempLayers
	}
	editingRoom := editor.editingRoom

	// Put all layer types into one array and sort
	layers := editor.tempLayers[:0]
	for _, layer := range editingRoom.InstanceLayers {
		layers = append(layers, layer)
	}
	for _, layer := range editingRoom.BackgroundLayers {
		layers = append(layers, layer)
	}
	for _, layer := range editingRoom.SpriteLayers {
		layers = append(layers, layer)
	}
	sort.Slice(layers, func(i, j int) bool {
		return layers[i].GetConfig().Order < layers[j].GetConfig().Order
	})
	editor.tempLayers = layers
	return layers
}

func roomEditorEditingRoom() *room.Room {
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

func editorLazyInit() {
	if gRoomEditor == nil {
		gRoomEditor = newRoomEditor()
	}
	// TODO(Jake): 2018-07-10
	//
	// Load editor font (possibly by embedding data into `reditor`?)
	//
}

/*func EditorIsInitialized() bool {
	return gRoomEditor != nil
}

func EditorSetRoom(room *Room) {
	roomEditor := gRoomEditor
	if roomEditor.editorChangeRoom(room) {
		roomEditor.editorConfigLoad()
		debugMenuOpenOrToggleClosed(debugMenuRoomEditor)
	}
}
*/
func (roomEditor *roomEditor) editorChangeRoom(room *Room) bool {
	if roomEditor.editingRoom == room {
		// If no changes
		return false
	}
	if room == nil {
		roomEditor.editingRoom = nil
		// Reset camera settings back
		*gCameraManager = roomEditor.cameraStateBeforeEnteringEditingMode
		debugMenuOpenOrToggleClosed(debugMenuNone)
		// Execute custom user-code logic
		//if gRoomEditor.exitEditorFunc != nil {
		//	gRoomEditor.exitEditorFunc(editingRoom)
		//}
		return false
	}
	roomEditor.editingRoom = room
	roomEditor.editingLayer = nil
	roomEditor.cameraStateBeforeEnteringEditingMode = *gCameraManager
	roomEditor.calculateAndSortLayers()
	roomEditor.calculateRoomBounds()

	// NOTE(Jake): 2018-07-09
	//
	// If you move around as the player a bit then go into the
	// editor. Retain the same camera position.
	//
	roomEditor.camPos = CameraGetViewPos(0)
	CameraSetViewSize(0, geom.Vec{float64(windowWidth()), float64(windowHeight())})
	CameraSetViewTarget(0, nil)
	return true
}

func editorUpdate() {
	roomEditor := gRoomEditor
	roomEditor.calculateAndSortLayers() // reset layers / recalculate sort order lazily with layers()
	isMenuOpen := roomEditor.IsMenuOpen()
	canUseBrush := true
	grid := geom.Vec{32, 32}

	// Setup mouse left, this is so we can force it to false to disable brush strokes (ie. for UI clicks)
	if MouseCheckPressed(MbLeft) {
		roomEditor.mouseHold[MbLeft] = true
	}
	if !MouseCheckButton(MbLeft) {
		roomEditor.mouseHold[MbLeft] = false
	}
	if MouseCheckPressed(MbMiddle) {
		roomEditor.mouseHold[MbMiddle] = true
	}
	if !MouseCheckButton(MbMiddle) {
		roomEditor.mouseHold[MbMiddle] = false
	}
	if MouseCheckPressed(MbRight) {
		roomEditor.mouseHold[MbRight] = true
	}
	if !MouseCheckButton(MbRight) {
		roomEditor.mouseHold[MbRight] = false
	}

	// Remove status text if time passed
	if roomEditor.statusText != "" &&
		time.Since(roomEditor.statusTimer).Seconds() > 5 {
		roomEditor.statusText = ""
	}

	isHoldingControl := KeyboardCheck(VkControl)
	if isHoldingControl {
		// Enable / disable grid
		if KeyboardCheckPressed(VkG) {
			roomEditor.gridEnabled = !roomEditor.gridEnabled
		}

		switch l := roomEditor.editingLayer.(type) {
		case nil:
			// no-op
			if KeyboardCheckPressed(VkP) {
				roomEditor.setStatusText("No layer selected. Cannot perform CTRL+P action.")
			}
		case *room.RoomLayerInstance:
			// Open entity select menu
			if KeyboardCheckPressed(VkP) {
				if roomEditor.menuOpened == reditor.MenuNone {
					roomEditor.menuOpened = reditor.MenuEntity
					ClearKeyboardString()
				} else {
					roomEditor.menuOpened = reditor.MenuNone
				}
			}

			// Open order set menu
			if KeyboardCheckPressed(VkO) {
				if roomEditor.menuOpened == reditor.MenuNone {
					roomEditor.menuOpened = reditor.MenuSetOrder
					ClearKeyboardString()
				} else {
					roomEditor.menuOpened = reditor.MenuNone
				}
			}
		case *room.RoomLayerBackground:
			// Open background select menu
			if KeyboardCheckPressed(VkP) {
				if roomEditor.menuOpened == reditor.MenuNone {
					roomEditor.menuOpened = reditor.MenuBackground
					ClearKeyboardString()
				} else {
					roomEditor.menuOpened = reditor.MenuNone
				}
			}

			// Open order set menu
			if KeyboardCheckPressed(VkO) {
				if roomEditor.menuOpened == reditor.MenuNone {
					roomEditor.menuOpened = reditor.MenuSetOrder
					ClearKeyboardString()
				} else {
					roomEditor.menuOpened = reditor.MenuNone
				}
			}
		case *room.RoomLayerSprite:
			// Open background select menu
			if KeyboardCheckPressed(VkP) {
				if roomEditor.menuOpened == reditor.MenuNone {
					roomEditor.menuOpened = reditor.MenuSprite
					ClearKeyboardString()
				} else {
					roomEditor.menuOpened = reditor.MenuNone
				}
			}

			// Open order set menu
			if KeyboardCheckPressed(VkO) {
				if roomEditor.menuOpened == reditor.MenuNone {
					roomEditor.menuOpened = reditor.MenuSetOrder
					ClearKeyboardString()
				} else {
					roomEditor.menuOpened = reditor.MenuNone
				}
			}
		default:
			panic(fmt.Sprintf("CTRL+Key: Unimplemented layer type: %T", l))
		}

		// Exit editor
		if KeyboardCheckPressed(VkR) {
			roomEditor.editorChangeRoom(nil)
			return
		}

		// Load map
		if KeyboardCheckPressed(VkL) {
			if roomEditor.menuOpened == reditor.MenuNone {
				roomEditor.menuOpened = reditor.MenuLoadRoom
				ClearKeyboardString()
			} else {
				roomEditor.menuOpened = reditor.MenuNone
			}
		}

		// New map
		if KeyboardCheckPressed(VkN) {
			if roomEditor.menuOpened == reditor.MenuNone {
				roomEditor.menuOpened = reditor.MenuNewRoom
				ClearKeyboardString()
			} else {
				roomEditor.menuOpened = reditor.MenuNone
			}
		}

		// Save
		if KeyboardCheckPressed(VkS) {
			editorSave()
		}
	}

	// Close open menus
	if KeyboardCheckPressed(VkEscape) {
		roomEditor.menuOpened = reditor.MenuNone
	}

	// Handle filtering / selection
	switch roomEditor.menuOpened {
	case reditor.MenuNewRoom,
		reditor.MenuLoadRoom,
		reditor.MenuNewLayer,
		reditor.MenuSetOrder:
		if inputSelectPressed() {
			if typingText := KeyboardString(); len(typingText) > 0 {
				switch roomEditor.menuOpened {
				case reditor.MenuNewRoom:
					// Create new room if typed
					roomEditor.newRoom(typingText)
				case reditor.MenuLoadRoom:
					// Load room if typed
					roomEditor.loadRoom(typingText)
				case reditor.MenuNewLayer:
					// Create new room if typed
					roomEditor.newLayerAndSelected(roomEditor.editingRoom, typingText, roomEditor.menuLayerKind)
				case reditor.MenuSetOrder:
					//
					layer := roomEditor.editingLayer
					if layer == nil {
						roomEditor.setStatusText("Cannot open set order menu when not editing any layer.")
						break
					}
					order, err := strconv.Atoi(typingText)
					if err != nil {
						roomEditor.setStatusText("Cannot set order to invalid value: " + typingText)
						break
					}
					layer.GetConfig().Order = int32(order)
					sort.Slice(roomEditor.editingRoom.InstanceLayers, func(i, j int) bool {
						return roomEditor.editingRoom.InstanceLayers[i].Config.Order < roomEditor.editingRoom.InstanceLayers[j].Config.Order
					})
					sort.Slice(roomEditor.editingRoom.BackgroundLayers, func(i, j int) bool {
						return roomEditor.editingRoom.BackgroundLayers[i].Config.Order < roomEditor.editingRoom.BackgroundLayers[j].Config.Order
					})
					sort.Slice(roomEditor.editingRoom.SpriteLayers, func(i, j int) bool {
						return roomEditor.editingRoom.SpriteLayers[i].Config.Order < roomEditor.editingRoom.SpriteLayers[j].Config.Order
					})
					roomEditor.hasUnsavedChanges = true
				default:
					panic("Unhandled menu type (ie. new room, new layer, set order)")
				}

			}
			// Close menu
			roomEditor.menuOpened = reditor.MenuNone
		}
	}

	/*editingRoom := roomEditor.editingRoom
	if editingRoom == nil {
		return
	}*/

	if !isMenuOpen &&
		!isHoldingControl {
		camPos := &roomEditor.camPos
		{
			// Move camera with WASD
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

		lastMouseScreenPos := roomEditor.lastMouseScreenPos
		roomEditor.lastMouseScreenPos = mouseScreenPosition()

		{
			// Move camera with middle mouse
			if MouseCheckButton(MbMiddle) {
				mouseDistMoved := mouseScreenPosition()
				mouseDistMoved.X -= lastMouseScreenPos.X
				mouseDistMoved.Y -= lastMouseScreenPos.Y
				sensitivity := -1.0
				if mouseDistMoved.X > 0 {
					camPos.X += mouseDistMoved.X * sensitivity
				} else if mouseDistMoved.X < 0 {
					camPos.X += mouseDistMoved.X * sensitivity
				}
				if mouseDistMoved.Y > 0 {
					camPos.Y += mouseDistMoved.Y * sensitivity
				} else if mouseDistMoved.Y < 0 {
					camPos.Y += mouseDistMoved.Y * sensitivity
				}
				CameraSetViewPos(0, *camPos)
			}
		}
	}

	lastMousePos := roomEditor.lastMousePos
	roomEditor.lastMousePos = MousePosition()

	// Draw
	{
		cameraSize := cameraGetActive().size

		{
			// Fill screen with gray
			DrawSetGUI(true)
			DrawRectangle(geom.Vec{0, 0}, cameraSize, color.RGBA{153, 153, 153, 255})
			DrawSetGUI(false)

			// Draw black box using room bounds
			if editingRoom := roomEditor.editingRoom; editingRoom != nil {
				borderWidth := 2.0
				pos := geom.Vec{float64(editingRoom.Left), float64(editingRoom.Top)}
				pos.X -= borderWidth
				pos.Y -= borderWidth
				size := geom.Vec{float64(editingRoom.Right - editingRoom.Left), float64(editingRoom.Bottom - editingRoom.Top)}
				size.X += borderWidth * 2
				size.Y += borderWidth * 2
				DrawRectangleBorder(pos, size, color.Black, 2, color.White)
			}
		}

		// Draw layers
		if editingRoom := roomEditor.editingRoom; editingRoom != nil {
			for _, layer := range roomEditor.layers() {
				switch layer := layer.(type) {
				case *room.RoomLayerInstance:
					// Draw room instances
					for _, obj := range layer.Instances {
						inst := roomEditorObjectIndexToData(obj.ObjectIndex)
						if inst == nil {
							continue
						}
						baseObj := inst.BaseObject()
						baseObj.X = float64(obj.X)
						baseObj.Y = float64(obj.Y)
						DrawSpriteScaled(baseObj.SpriteIndex(), 0, baseObj.Pos(), baseObj.ImageScale)
					}
				case *room.RoomLayerBackground:
					if layer.SpriteName == "" {
						// If no sprite, don't draw anything
						break
					}
					// Draw bg
					sprite := sprite.SpriteLoadByName(layer.SpriteName)
					x := float64(layer.X)
					y := float64(layer.Y)
					width := float64(sprite.Size().X)
					DrawSprite(sprite, 0, geom.Vec{x, y})
					{
						// Tile left
						x := x
						for x > float64(editingRoom.Left) {
							x -= width
							DrawSprite(sprite, 0, geom.Vec{x, y})
						}
					}
					{
						// Tile left
						x := x
						for x < float64(editingRoom.Right) {
							x += width
							DrawSprite(sprite, 0, geom.Vec{x, y})
						}
					}
				case *room.RoomLayerSprite:
					// Draw room sprites
					for _, obj := range layer.Sprites {
						spriteName := obj.SpriteName
						sprite := sprite.SpriteLoadByName(spriteName)
						DrawSprite(sprite, 0, geom.Vec{float64(obj.X), float64(obj.Y)})
					}
				default:
					panic(fmt.Sprintf("Unhandled type: %T", layer))
				}
			}
			// Old render code (before layers)
			/*instances := editingRoom.Instances
			for _, obj := range instances {
				inst := roomEditorObjectIndexToData(obj.ObjectIndex)
				if inst == nil {
					continue
				}
				baseObj := inst.BaseObject()
				baseObj.X = float64(obj.X)
				baseObj.Y = float64(obj.Y)
				inst.Draw()
			}*/
		}

		if canUseBrush &&
			!isMenuOpen {
			// Place tile
			switch layer := roomEditor.editingLayer.(type) {
			case nil:
				// no-op
			case *room.RoomLayerInstance:
				// Draw selected tile
				if selectedObj := roomEditor.objectSelected; selectedObj != nil {
					mousePos := MousePosition()
					mousePos.X = snapToGrid(mousePos.X, grid.X)
					mousePos.Y = snapToGrid(mousePos.Y, grid.Y)

					// Draw
					drawObject(selectedObj, mousePos)
				}
			case *room.RoomLayerSprite:
				// Draw selected sprite
				if selectedBrush := roomEditor.spriteSelected; selectedBrush.IsValid() {
					pos := MousePosition()
					if isHoldingControl {
						maybeNewPos, ok := roomEditor.getSnapPosition(pos, selectedBrush.Size(), layer)
						if ok {
							pos = maybeNewPos
							DrawSpriteExt(selectedBrush, 0, pos, geom.Vec{1, 1}, 0.85)
						}
					} else {
						// Grid mode
						pos.X = snapToGrid(pos.X, grid.X)
						pos.Y = snapToGrid(pos.Y, grid.Y)
						DrawSpriteExt(selectedBrush, 0, pos, geom.Vec{1, 1}, 0.85)
					}
				}
			case *room.RoomLayerBackground:
				// no-op
			default:
				panic(fmt.Sprintf("Unimplemented layer type: %T", layer))
			}
		}

		if !isMenuOpen {
			// Draw grid
			if !isHoldingControl &&
				roomEditor.gridEnabled {
				DrawSetGUI(true)
				cameraPos := CameraGetViewPos(0)
				windowWidth := float64(windowWidth())
				windowHeight := float64(windowHeight())
				xOffset := math.Mod(cameraPos.X, grid.X)
				yOffset := math.Mod(cameraPos.Y, grid.Y)
				for y := 0.0; y < windowHeight+grid.Y; y += grid.Y {
					DrawRectangle(geom.Vec{0, y - yOffset}, geom.Vec{windowWidth + grid.X, 1}, color.White)
				}
				for x := 0.0; x < windowWidth+grid.X; x += grid.X {
					DrawRectangle(geom.Vec{x - xOffset, 0}, geom.Vec{1, windowHeight + grid.Y}, color.White)
				}
				DrawSetGUI(false)
			}
		}

		if roomEditor.menuOpened == reditor.MenuNone {
			// Draw layer widget
			if roomEditor.editingRoom != nil {
				DrawSetGUI(true)
				yStart := 0.0
				var x, y, width, height float64
				x = 0
				y = yStart
				width = 320
				height = float64(len(roomEditor.layers())+1)*32.0 + 16
				DrawRectangle(geom.Vec{x, y}, geom.Vec{width, height}, color.RGBA{0, 0, 0, 220})
				y += 24
				DrawText(geom.Vec{x + 16, y}, "Layers:")
				y += 16
				for _, layer := range roomEditor.layers() {
					config := layer.GetConfig()
					layerName := config.Name
					layerType := "Unknown"
					switch layer.(type) {
					case *room.RoomLayerInstance:
						layerType = "Instance"
					case *room.RoomLayerBackground:
						layerType = "Background"
					case *room.RoomLayerSprite:
						layerType = "Sprite"
					default:
						layerType = "Unhandled case"
					}
					layerText := layerName + " - " + layerType + " Layer (Order: " + strconv.Itoa(int(config.Order)) + ")"
					if layer == roomEditor.editingLayer {
						layerText += " [Selected]"
					}
					if drawTextButton(geom.Vec{x, y}, layerText) {
						if roomEditor.editingLayer == layer {
							roomEditor.editingLayer = nil
						} else {
							roomEditor.editingLayer = layer
						}
						roomEditor.MouseClearButton(MbLeft)
					}
					y += 32
				}
				y = yStart + height
				switch layer := roomEditor.editingLayer.(type) {
				case *room.RoomLayerSprite:
					text := "Collision OFF"
					if layer.Config.HasCollision {
						text = "Collision ON"
					}
					if drawButton(geom.Vec{x, y}, text) {
						layer.Config.HasCollision = !layer.Config.HasCollision
						roomEditor.hasUnsavedChanges = true
						roomEditor.MouseClearButton(MbLeft)
					}
					y += 32
				}
				if drawButton(geom.Vec{x, y}, "Add Instance Layer") {
					ClearKeyboardString()
					roomEditor.menuOpened = reditor.MenuNewLayer
					roomEditor.menuLayerKind = room.RoomLayerKind_Instance
					roomEditor.MouseClearButton(MbLeft)
				}
				y += 32
				if drawButton(geom.Vec{x, y}, "Add Background Layer") {
					ClearKeyboardString()
					roomEditor.menuOpened = reditor.MenuNewLayer
					roomEditor.menuLayerKind = room.RoomLayerKind_Background
					roomEditor.MouseClearButton(MbLeft)
				}
				y += 32
				if drawButton(geom.Vec{x, y}, "Add Sprite Layer") {
					ClearKeyboardString()
					roomEditor.menuOpened = reditor.MenuNewLayer
					roomEditor.menuLayerKind = room.RoomLayerKind_Sprite
					roomEditor.MouseClearButton(MbLeft)
				}
				y += 32
				text := "Grid OFF"
				if roomEditor.gridEnabled {
					text = "Grid ON"
				}
				if drawButton(geom.Vec{x, y}, text) {
					roomEditor.gridEnabled = !roomEditor.gridEnabled
					roomEditor.MouseClearButton(MbLeft)
				}
				y += 32
				if layer := roomEditor.editingLayer; layer != nil {
					if drawButton(geom.Vec{x, y}, "Delete Selected Layer") {
						switch layer := layer.(type) {
						case *room.RoomLayerInstance:
							layerName := layer.Config.Name
							layerUUID := layer.Config.UUID
							didDelete := false
							records := roomEditor.editingRoom.InstanceLayers
							for i, record := range records {
								if layer == record {
									// Unordered Remove
									records[len(records)-1], records[i] = records[i], records[len(records)-1]
									records = records[:len(records)-1]
									didDelete = true
									break
								}
							}
							if !didDelete {
								roomEditor.setStatusText("Failed to delete instance layer.")
								break
							}
							roomEditor.editingLayer = nil
							roomEditor.editingRoom.InstanceLayers = records
							roomEditor.editingRoom.DeletedLayers = append(roomEditor.editingRoom.DeletedLayers, layerUUID)
							roomEditor.setStatusText(fmt.Sprintf("Deleted instance \"%s\" layer. (UUID: \"%s\")", layerName, layerUUID))
						case *room.RoomLayerBackground:
							layerName := layer.Config.Name
							layerUUID := layer.Config.UUID
							didDelete := false
							records := roomEditor.editingRoom.BackgroundLayers
							for i, record := range records {
								if layer == record {
									// Unordered Remove
									records[len(records)-1], records[i] = records[i], records[len(records)-1]
									records = records[:len(records)-1]
									didDelete = true
									break
								}
							}
							if !didDelete {
								roomEditor.setStatusText("Failed to delete instance layer.")
								break
							}
							roomEditor.editingLayer = nil
							roomEditor.editingRoom.BackgroundLayers = records
							roomEditor.editingRoom.DeletedLayers = append(roomEditor.editingRoom.DeletedLayers, layerUUID)
							roomEditor.setStatusText(fmt.Sprintf("Deleted instance \"%s\" layer. (UUID: \"%s\")", layerName, layerUUID))
						case *room.RoomLayerSprite:
							layerName := layer.Config.Name
							layerUUID := layer.Config.UUID
							didDelete := false
							records := roomEditor.editingRoom.SpriteLayers
							for i, record := range records {
								if layer == record {
									// Unordered Remove
									records[len(records)-1], records[i] = records[i], records[len(records)-1]
									records = records[:len(records)-1]
									didDelete = true
									break
								}
							}
							if !didDelete {
								roomEditor.setStatusText("Failed to delete instance layer.")
								break
							}
							roomEditor.editingLayer = nil
							roomEditor.editingRoom.SpriteLayers = records
							roomEditor.editingRoom.DeletedLayers = append(roomEditor.editingRoom.DeletedLayers, layerUUID)
							roomEditor.setStatusText(fmt.Sprintf("Deleted instance \"%s\" layer. (UUID: \"%s\")", layerName, layerUUID))
						default:
							roomEditor.setStatusText(fmt.Sprintf("Unhandled deletion case for: %T", layer))
						}
						roomEditor.MouseClearButton(MbLeft)
					}
					y += 32
				}
				DrawSetGUI(false)
			}
		}

		if canUseBrush &&
			!isMenuOpen {
			// Place tile
			switch layer := roomEditor.editingLayer.(type) {
			case nil:
				// no-op
			case *room.RoomLayerInstance:
				// Left click
				if roomEditor.MouseCheckButton(MbLeft) &&
					roomEditor.objectSelected != nil {
					objectIndexSelected := roomEditor.objectSelected.ObjectIndex()
					// NOTE(Jake): 2018-06-10
					//
					// We need to handle mouse click between the last mouse position
					// and current mouse position so that there are no gaps when you're
					// dragging the mouse across long distances.
					//
					rect := geom.R(MousePosition(), lastMousePos)
					didCreate := false
					for x := rect.Left(); x <= rect.Right(); x += grid.X {
						for y := rect.Top(); y <= rect.Bottom(); y += grid.Y {
							mousePos := geom.Vec{x, y}

							// Check to make sure we aren't placing over the top
							// of an existing entity
							hasCollision := false
							for _, obj := range layer.Instances {
								inst := roomEditorObjectIndexToData(obj.ObjectIndex)
								if inst == nil {
									continue
								}
								pos := geom.Vec{float64(obj.X), float64(obj.Y)}
								size := inst.BaseObject().Size
								left := pos.X
								right := left + float64(size.X)
								top := pos.Y
								bottom := top + float64(size.Y)
								if mousePos.X >= left && mousePos.X < right &&
									mousePos.Y >= top && mousePos.Y < bottom {
									hasCollision = true
								}
							}

							//
							if !hasCollision {
								// Snap to grid
								mousePos.X = snapToGrid(mousePos.X, grid.X)
								mousePos.Y = snapToGrid(mousePos.Y, grid.Y)

								// Add instance
								pos := mousePos
								//count := editingRoom.UserEntityCount
								//editingRoom.UserEntityCount++

								//
								roomObj := &room.RoomObject{
									UUID:        reditor.UUID(),
									ObjectIndex: int32(objectIndexSelected),
									X:           int32(pos.X),
									Y:           int32(pos.Y),
								}
								layer.Instances = append(layer.Instances, roomObj)

								println("Create entity:", roomObj.String())
								didCreate = true
							}
						}
					}
					if didCreate {
						roomEditor.hasUnsavedChanges = true
						roomEditor.calculateRoomBounds()
					}
				}

				// Holding Right click
				if roomEditor.MouseCheckButton(MbRight) {
					// NOTE(Jake): 2018-06-10
					//
					// We need to handle mouse click between the last mouse position
					// and current mouse position so that there are no gaps when you're
					// dragging the mouse across long distances.
					//
					rect := geom.R(MousePosition(), lastMousePos)
					for x := rect.Left(); x <= rect.Right(); x += grid.X {
						for y := rect.Top(); y <= rect.Bottom(); y += grid.Y {
							mousePos := geom.Vec{x, y}
							previousDeletedCount := len(layer.DeletedInstances)

							// Mark deleted entities
							for i, obj := range layer.Instances {
								inst := roomEditorObjectIndexToData(obj.ObjectIndex)
								if inst == nil {
									continue
								}
								pos := geom.Vec{float64(obj.X), float64(obj.Y)}
								size := inst.BaseObject().Size
								left := pos.X
								right := left + float64(size.X)
								top := pos.Y
								bottom := top + float64(size.Y)
								if mousePos.X >= left && mousePos.X < right &&
									mousePos.Y >= top && mousePos.Y < bottom {
									//
									record := layer.Instances[i]

									// Track deleted entities for when you save
									layer.DeletedInstances = append(layer.DeletedInstances, record)
									roomEditor.hasUnsavedChanges = true
								}
							}

							// Handle deletes
							if previousDeletedCount != len(layer.DeletedInstances) {
								// NOTE(Jake): 2018-07-11
								//
								// Iterate over newly deleted entities and delete them from
								// the room
								//
								didDelete := false
								for i := previousDeletedCount; i < len(layer.DeletedInstances); i++ {
									record := layer.DeletedInstances[i]
									for i, obj := range layer.Instances {
										if obj == record {
											{
												// Unordered Remove
												layer.Instances[len(layer.Instances)-1], layer.Instances[i] = layer.Instances[i], layer.Instances[len(layer.Instances)-1]
												layer.Instances = layer.Instances[:len(layer.Instances)-1]
											}
											didDelete = true
											println("Deleted entity: " + record.UUID)
											break
										}
									}
								}
								if didDelete {
									roomEditor.calculateRoomBounds()
									roomEditor.hasUnsavedChanges = true
								}
							}
						}
					}
				}
			case *room.RoomLayerSprite:
				// Left click
				if roomEditor.MouseCheckButton(MbLeft) &&
					roomEditor.spriteSelected.IsValid() {
					spriteName := roomEditor.spriteSelected.Name()
					//spriteWidth := roomEditor.spriteSelected.Size().X
					//spriteHeight := roomEditor.spriteSelected.Size().Y

					// NOTE(Jake): 2018-06-10
					//
					// We need to handle mouse click between the last mouse position
					// and current mouse position so that there are no gaps when you're
					// dragging the mouse across long distances.
					//
					rect := geom.R(MousePosition(), lastMousePos)
					didCreate := false
					if isHoldingControl {
						selectedBrush := roomEditor.spriteSelected
						brushWidth := float64(selectedBrush.Size().X) / 2
						brushHeight := float64(selectedBrush.Size().Y) / 2

						for x := rect.Left(); x <= rect.Right(); x += brushWidth {
							for y := rect.Top(); y <= rect.Bottom(); y += brushHeight {
								pos := geom.Vec{x, y}
								maybeNewPos, ok := roomEditor.getSnapPosition(pos, selectedBrush.Size(), layer)
								if !ok {
									continue
								}
								pos = maybeNewPos

								brushRect := geom.Rect{}
								brushRect.Vec = pos
								brushRect.Size = roomEditor.spriteSelected.Size()

								// Check to make sure we aren't placing over the top
								// of an existing entity
								hasCollision := false
								for _, obj := range layer.Sprites {
									spriteName := obj.SpriteName
									sprite := sprite.SpriteLoadByName(spriteName)

									other := geom.Rect{}
									other.X = float64(obj.X)
									other.Y = float64(obj.Y)
									other.Size = sprite.Size()

									if brushRect.CollisionRectangle(other) {
										hasCollision = true
									}
								}

								//
								if !hasCollision {
									//
									roomSpriteObj := &room.RoomSpriteObject{
										UUID:       reditor.UUID(),
										SpriteName: spriteName,
										X:          int32(pos.X),
										Y:          int32(pos.Y),
									}
									layer.Sprites = append(layer.Sprites, roomSpriteObj)

									println("Create sprite:", roomSpriteObj.String())
									didCreate = true
								}
							}
						}
					} else {
						for x := rect.Left(); x <= rect.Right(); x += grid.X {
							for y := rect.Top(); y <= rect.Bottom(); y += grid.Y {
								pos := geom.Vec{x, y}
								// Grid mode
								pos.X = snapToGrid(pos.X, grid.X)
								pos.Y = snapToGrid(pos.Y, grid.Y)

								brushRect := geom.Rect{}
								brushRect.Vec = pos
								brushRect.Size = roomEditor.spriteSelected.Size()

								// Check to make sure we aren't placing over the top
								// of an existing entity
								hasCollision := false
								for _, obj := range layer.Sprites {
									spriteName := obj.SpriteName
									sprite := sprite.SpriteLoadByName(spriteName)

									other := geom.Rect{}
									other.X = float64(obj.X)
									other.Y = float64(obj.Y)
									other.Size = sprite.Size()

									if brushRect.CollisionRectangle(other) {
										hasCollision = true
									}
								}

								//
								if !hasCollision {
									//
									roomSpriteObj := &room.RoomSpriteObject{
										UUID:       reditor.UUID(),
										SpriteName: spriteName,
										X:          int32(pos.X),
										Y:          int32(pos.Y),
									}
									layer.Sprites = append(layer.Sprites, roomSpriteObj)

									println("Create sprite:", roomSpriteObj.String())
									didCreate = true
								}
							}
						}
					}
					if didCreate {
						roomEditor.hasUnsavedChanges = true
						roomEditor.calculateRoomBounds()
					}
				}

				// Holding Right click
				if roomEditor.MouseCheckButton(MbRight) {
					// NOTE(Jake): 2018-06-10
					//
					// We need to handle mouse click between the last mouse position
					// and current mouse position so that there are no gaps when you're
					// dragging the mouse across long distances.
					//
					rect := geom.R(MousePosition(), lastMousePos)
					for x := rect.Left(); x <= rect.Right(); x += grid.X {
						for y := rect.Top(); y <= rect.Bottom(); y += grid.Y {
							mousePos := geom.Vec{x, y}
							previousDeletedCount := len(layer.DeletedSprites)

							// Mark deleted
							for i, obj := range layer.Sprites {
								spriteName := obj.SpriteName
								sprite := sprite.SpriteLoadByName(spriteName)

								width := float64(sprite.Size().X)
								height := float64(sprite.Size().Y)
								left := float64(obj.X)
								right := left + width
								top := float64(obj.Y)
								bottom := top + height
								if mousePos.X >= left && mousePos.X < right &&
									mousePos.Y >= top && mousePos.Y < bottom {
									//
									record := layer.Sprites[i]

									// Track deleted entities for when you save
									layer.DeletedSprites = append(layer.DeletedSprites, record)
									roomEditor.hasUnsavedChanges = true
								}
							}

							// Handle deletes
							if previousDeletedCount != len(layer.DeletedSprites) {
								// NOTE(Jake): 2018-07-11
								//
								// Iterate over newly deleted entities and delete them from
								// the room
								//
								didDelete := false
								for i := previousDeletedCount; i < len(layer.DeletedSprites); i++ {
									record := layer.DeletedSprites[i]
									for i, obj := range layer.Sprites {
										if obj == record {
											{
												// Unordered Remove
												layer.Sprites[len(layer.Sprites)-1], layer.Sprites[i] = layer.Sprites[i], layer.Sprites[len(layer.Sprites)-1]
												layer.Sprites = layer.Sprites[:len(layer.Sprites)-1]
											}
											didDelete = true
											println("Deleted entity: " + record.UUID)
											break
										}
									}
								}
								if didDelete {
									roomEditor.calculateRoomBounds()
									roomEditor.hasUnsavedChanges = true
								}
							}
						}
					}
				}
			case *room.RoomLayerBackground:
				if roomEditor.MouseCheckButton(MbLeft) {
					mousePos := MousePosition()
					mousePos.X = snapToGrid(mousePos.X, grid.X)
					mousePos.Y = snapToGrid(mousePos.Y, grid.Y)
					layer.X = int32(mousePos.X)
					layer.Y = int32(mousePos.Y)

					roomEditor.calculateRoomBounds()
					roomEditor.hasUnsavedChanges = true
				}
			default:
				panic(fmt.Sprintf("Unimplemented layer type: %T", layer))
			}
		}

		// Select menu
		if roomEditor.menuOpened != reditor.MenuNone {
			switch roomEditor.menuOpened {
			case reditor.MenuEntity:
				DrawSetGUI(true)
				// Add black opacity over screen with menu open
				DrawRectangle(geom.Vec{0, 0}, geom.Vec{2048, 2048}, color.RGBA{0, 0, 0, 190})

				//
				ui := geom.Vec{
					X: float64(windowWidth()) / 2,
					Y: 32,
				}
				typingText := KeyboardString()
				roomEditor.entityMenuFiltered = roomEditor.entityMenuFiltered[:0]
				for _, obj := range roomEditor.objectIndexToData {
					if obj == nil {
						continue
					}
					name := obj.ObjectName()
					hasMatch := hasFilterMatch(name, typingText)
					if !hasMatch {
						continue
					}
					// NOTE(Jake): 2018-07-11
					//
					// Animating in the object list isn't particularly useful.
					//
					//obj.BaseObject().ImageUpdate()
					roomEditor.entityMenuFiltered = append(roomEditor.entityMenuFiltered, obj)
				}

				//
				if inputSelectPressed() &&
					len(roomEditor.entityMenuFiltered) > 0 {
					selectedObj := roomEditor.entityMenuFiltered[0]
					// Set
					roomEditor.objectSelected = selectedObj
					roomEditor.editorConfigSave()

					roomEditor.menuOpened = reditor.MenuNone
				}

				{
					searchText := "Search for object (type + press enter)"
					DrawText(geom.Vec{ui.X - (StringWidth(searchText) / 4), ui.Y}, searchText)
					ui.Y += 24
				}
				{
					typingText := KeyboardString()
					DrawText(ui, typingText)
					DrawText(geom.Vec{ui.X + StringWidth(typingText), ui.Y}, "|")
					ui.Y += 24
				}
				previewSize := geom.Vec{32, 32}
				for _, obj := range roomEditor.entityMenuFiltered {
					//
					baseObj := obj.BaseObject()
					pos := baseObj.Vec
					//size := baseObj.Size
					//oldImageScale := baseObj.ImageScale

					// NOTE(Jake): 2018-07-10
					//
					// I've already wasted time thinking about centering this
					// and the text above. Lets not look into this until we
					// feel the need.
					//
					// Also look at other similar UI experiences, because
					// maybe we wont need to center this to get good UX.
					//
					pos.X = ui.X - 40
					pos.Y = ui.Y - (previewSize.Y / 2)
					//baseObj.ImageScale.X = previewSize.X / float64(size.X)
					//baseObj.ImageScale.Y = previewSize.Y / float64(size.Y)
					//obj.Draw()
					drawObjectPreview(obj, pos, previewSize)
					name := obj.ObjectName()
					DrawText(ui, name)
					//baseObj.ImageScale = oldImageScale
					ui.Y += previewSize.Y + 16
				}
			case reditor.MenuSprite,
				reditor.MenuBackground:
				if spriteSelected, ok := roomEditor.spriteViewer.update(); ok {
					switch roomEditor.menuOpened {
					case reditor.MenuSprite:
						roomEditor.spriteSelected = spriteSelected
						roomEditor.editorConfigSave()
					case reditor.MenuBackground:
						switch layer := roomEditor.editingLayer.(type) {
						case *room.RoomLayerBackground:
							if layer.SpriteName != spriteSelected.Name() {
								layer.SpriteName = spriteSelected.Name()
								roomEditor.hasUnsavedChanges = true
							}
						default:
							panic(fmt.Sprintf("MenuBackground: Unhandled layer type for %T", layer))
						}
					}
					roomEditor.menuOpened = reditor.MenuNone
				}
			case reditor.MenuNewRoom,
				reditor.MenuLoadRoom,
				reditor.MenuNewLayer,
				reditor.MenuSetOrder:
				searchText := ""
				switch roomEditor.menuOpened {
				case reditor.MenuNewRoom:
					searchText = "Enter name for new room:"
				case reditor.MenuLoadRoom:
					searchText = "Enter name to load room:"
				case reditor.MenuNewLayer:
					searchText = "Enter name for new layer:"
				case reditor.MenuSetOrder:
					searchText = "Set value for order (ie. -1000, 100, 10000, 5, -3, 0):"
				default:
					panic("Invalid menu type, No search text defined")
				}

				DrawSetGUI(true)
				// Add black opacity over screen with menu open
				DrawRectangle(geom.Vec{0, 0}, geom.Vec{2048, 2048}, color.RGBA{0, 0, 0, 190})

				//
				ui := geom.Vec{
					X: float64(windowWidth()) / 2,
					Y: 32,
				}
				{
					DrawText(geom.Vec{ui.X - (StringWidth(searchText) / 4), ui.Y}, searchText)
					ui.Y += 24
				}
				{
					typingText := KeyboardString()
					DrawText(ui, typingText)
					DrawText(geom.Vec{ui.X + StringWidth(typingText), ui.Y}, "|")
					ui.Y += 24
				}
			}
			DrawSetGUI(false)
		}

		// Draw status
		{
			DrawSetGUI(true)
			editingString := ""
			if editingRoom := roomEditor.editingRoom; editingRoom != nil {
				editingString += "Editing: " + filepath.Base(roomEditor.editingRoom.Filepath())
			}
			{
				mousePos := MousePosition()
				mousePos.X = snapToGrid(mousePos.X, grid.X)
				mousePos.Y = snapToGrid(mousePos.Y, grid.Y)
				if editingString != "" {
					editingString += " | "
				}
				editingString += strconv.FormatFloat(mousePos.X, 'f', -1, 64) + "," + strconv.FormatFloat(mousePos.Y, 'f', -1, 64) + "px"
			}
			if text := roomEditor.statusText; text != "" {
				editingString += " | " + text
			}
			if selectedObj := roomEditor.objectSelected; selectedObj != nil {
				editingString += " | Selected: " + selectedObj.ObjectName()
				//DrawText(geom.V(0, 16), "Selected: "+selectedObj.ObjectName())
			}
			editingString += " | Frame Usage: " + FrameUsage()
			//editingString += " | CTRL+P = Open Entity Menu"
			//editingString += " | CTRL+N = Create new room"
			DrawRectangle(geom.Vec{0, cameraSize.Y - 32}, geom.Vec{StringWidth(editingString) + 16, 32}, color.RGBA{0, 0, 0, 128})
			DrawText(geom.Vec{0, cameraSize.Y - 16}, editingString)

			DrawSetGUI(false)
		}
	}
}

func drawObjectPreview(inst object.ObjectType, pos geom.Vec, fitToSize geom.Vec) {
	baseObj := inst.BaseObject()
	sprite := baseObj.SpriteIndex()
	spriteSize := sprite.Size()
	scale := geom.Vec{fitToSize.X / float64(spriteSize.X), fitToSize.Y / float64(spriteSize.Y)}
	if scale.X > 1 {
		scale.X = 1
	}
	if scale.Y > 1 {
		scale.Y = 1
	}
	DrawSpriteScaled(sprite, 0, pos, scale)
}

func drawObject(inst object.ObjectType, pos geom.Vec) {
	baseObj := inst.BaseObject()
	//baseObj.Vec = pos
	DrawSprite(baseObj.SpriteIndex(), 0, pos)
}

func drawRoomObject(roomObject *room.RoomObject, pos geom.Vec) {
	inst := roomEditorObjectIndexToData(roomObject.ObjectIndex)
	if inst == nil {
		return
	}
	drawObject(inst, pos)
}

func drawTextButton(pos geom.Vec, text string) bool {
	// Config
	paddingH := 32.0
	size := geom.Vec{StringWidth(text) + paddingH, 24}

	// Handle mouse over
	isMouseOver := isMouseScreenOver(pos, size)

	// Draw highlight bg
	if isMouseOver {
		DrawRectangle(pos, size, color.White)
	}

	// Draw Text
	pos.X += paddingH * 0.5
	pos.Y += 16
	if isMouseOver {
		DrawTextColor(pos, text, color.Black)
	} else {
		DrawTextColor(pos, text, color.White)
	}
	return MouseCheckPressed(MbLeft) && isMouseOver
}

func (roomEditor *roomEditor) getSnapPosition(pos geom.Vec, brushSize geom.Size, layer *room.RoomLayerSprite) (geom.Vec, bool) {
	offsetX := float64(brushSize.X)
	offsetY := float64(brushSize.Y)
	var horizTarget, vertTarget geom.Rect
	var horizSide, vertSide int
	//fmt.Printf("Reset\n")
	if horizSide == 0 {
		targetPos := pos
		targetPos.X += offsetX

		brushRect := geom.Rect{}
		brushRect.Vec = targetPos
		brushRect.Size = brushSize

		for _, obj := range layer.Sprites {
			spriteName := obj.SpriteName
			sprite := sprite.SpriteLoadByName(spriteName)

			other := geom.Rect{}
			other.X = float64(obj.X)
			other.Y = float64(obj.Y)
			other.Size = sprite.Size()

			if brushRect.CollisionRectangle(other) {
				// Make sure mouse pos is on the left-side of record
				if pos.X <= other.Left() {
					if horizSide == 0 ||
						(other.DistancePoint(pos) < horizTarget.DistancePoint(pos)) {
						horizTarget = other
						horizSide = 1
					}
				}
			}
		}
	}
	{
		targetPos := pos
		targetPos.X -= offsetX

		brushRect := geom.Rect{}
		brushRect.Vec = targetPos
		brushRect.Size = brushSize

		for _, obj := range layer.Sprites {
			spriteName := obj.SpriteName
			sprite := sprite.SpriteLoadByName(spriteName)

			other := geom.Rect{}
			other.X = float64(obj.X)
			other.Y = float64(obj.Y)
			other.Size = sprite.Size()

			if brushRect.CollisionRectangle(other) {
				// Make sure mouse pos is on the right-side of record
				if pos.X > other.Right() {
					if horizSide == 0 ||
						(other.DistancePoint(pos) < horizTarget.DistancePoint(pos)) {
						horizTarget = other
						horizSide = -1
					}
				}
			}
		}
	}
	{
		targetPos := pos
		targetPos.Y += offsetY

		brushRect := geom.Rect{}
		brushRect.Vec = targetPos
		brushRect.Size = brushSize

		for _, obj := range layer.Sprites {
			spriteName := obj.SpriteName
			sprite := sprite.SpriteLoadByName(spriteName)

			other := geom.Rect{}
			other.X = float64(obj.X)
			other.Y = float64(obj.Y)
			other.Size = sprite.Size()

			if brushRect.CollisionRectangle(other) {
				// Make sure mouse pos is on the top-side of record
				if pos.Y < other.Top() {
					if vertSide == 0 ||
						(other.DistancePoint(pos) < vertTarget.DistancePoint(pos)) {
						vertTarget = other
						vertSide = 1
					}
				}
			}
		}
	}
	{
		targetPos := pos
		targetPos.Y -= offsetY

		brushRect := geom.Rect{}
		brushRect.Vec = targetPos
		brushRect.Size = brushSize

		for _, obj := range layer.Sprites {
			spriteName := obj.SpriteName
			sprite := sprite.SpriteLoadByName(spriteName)

			other := geom.Rect{}
			other.X = float64(obj.X)
			other.Y = float64(obj.Y)
			other.Size = sprite.Size()

			if brushRect.CollisionRectangle(other) {
				// Make sure mouse pos is on the bottom-side of record
				if pos.Y > other.Bottom() {
					if vertSide == 0 ||
						(other.DistancePoint(pos) < vertTarget.DistancePoint(pos)) {
						vertTarget = other
						vertSide = -1
					}
				}
			}
		}
	}
	if horizSide != 0 || vertSide != 0 {
		//brushRect := geom.Rect{}
		//brushRect.Vec = pos
		//brushRect.Size = brushSize

		var horizDist, vertDist float64
		horizDist = 9999
		vertDist = 9999
		if horizSide != 0 {
			horizDist = horizTarget.DistancePoint(pos)
		}
		if vertSide != 0 {
			vertDist = vertTarget.DistancePoint(pos)
		}
		if horizDist < vertDist {
			switch horizSide {
			case 0:
				// no-op
			case -1:
				pos.X = horizTarget.Right()
				pos.Y = horizTarget.Top()
			case 1:
				pos.X = horizTarget.Left() - float64(brushSize.X)
				pos.Y = horizTarget.Top()
			default:
				panic(fmt.Sprintf("unknown horiz side: %d", horizSide))
			}
		} else {
			switch vertSide {
			case 0:
				// no-op
			case -1:
				pos.X = vertTarget.Left()
				pos.Y = vertTarget.Bottom()
			case 1:
				pos.X = vertTarget.Left()
				pos.Y = vertTarget.Top() - float64(brushSize.Y)
			default:
				panic(fmt.Sprintf("unknown vert side: %d", vertSide))
			}
		}
		return pos, true
	}
	return pos, false
}

func (roomEditor *roomEditor) calculateRoomBounds() {
	// Reset room size
	editingRoom := roomEditor.editingRoom
	editingRoom.Left = 0
	editingRoom.Right = 0
	editingRoom.Top = 0
	editingRoom.Bottom = 0
	for _, layer := range roomEditor.layers() {
		switch layer := layer.(type) {
		case *room.RoomLayerInstance:
			for _, obj := range layer.Instances {
				inst := roomEditorObjectIndexToData(obj.ObjectIndex)
				if inst == nil {
					continue
				}
				x := int32(obj.X)
				y := int32(obj.Y)
				size := inst.BaseObject().Size
				width := int32(size.X)
				height := int32(size.Y)
				if x < editingRoom.Left {
					editingRoom.Left = x
				}
				if right := x + width; right > editingRoom.Right {
					editingRoom.Right = right
				}
				if y < editingRoom.Top {
					editingRoom.Top = y
				}
				if bottom := y + height; bottom > editingRoom.Bottom {
					editingRoom.Bottom = bottom + height
				}
			}
		case *room.RoomLayerSprite:
			for _, obj := range layer.Sprites {
				spriteName := obj.SpriteName
				sprite := sprite.SpriteLoadByName(spriteName)

				x := int32(obj.X)
				y := int32(obj.Y)
				width := int32(sprite.Size().X)
				height := int32(sprite.Size().Y)
				if x < editingRoom.Left {
					editingRoom.Left = x
				}
				if right := x + width; right > editingRoom.Right {
					editingRoom.Right = right
				}
				if y < editingRoom.Top {
					editingRoom.Top = y
				}
				if bottom := y + height; bottom > editingRoom.Bottom {
					editingRoom.Bottom = bottom + height
				}
			}
		}
	}
}

func (roomEditor *roomEditor) newRoom(name string) {
	if roomEditor.hasUnsavedChanges {
		roomEditor.setStatusText("Cannot create new room with unsaved changes. Restart the game and editor, then press CTRL+N")
		return
	}
	roomFilepath := roomEditor.roomDirectory + reditor.UUID()
	if _, err := os.Stat(roomFilepath); !os.IsNotExist(err) {
		roomEditor.setStatusText("Cannot create room \"" + filepath.Base(roomFilepath) + "\". That room name is already taken and exists.")
		return
	}

	editingRoom := &room.Room{
		Config: &room.RoomConfig{
			UUID: reditor.UUID(),
			Name: name,
		},
	}
	if roomEditor.editingRoom != nil {
		// Copy configs of current rooms layers
		for _, layer := range roomEditor.editingRoom.InstanceLayers {
			var copyLayer room.RoomLayerInstance
			copyLayer.Config = new(room.RoomLayerConfig)
			*copyLayer.Config = *layer.Config
			editingRoom.InstanceLayers = append(editingRoom.InstanceLayers, &copyLayer)
		}
		for _, layer := range roomEditor.editingRoom.SpriteLayers {
			var copyLayer room.RoomLayerSprite
			copyLayer.Config = new(room.RoomLayerConfig)
			*copyLayer.Config = *layer.Config
			editingRoom.SpriteLayers = append(editingRoom.SpriteLayers, &copyLayer)
		}
		for _, layer := range roomEditor.editingRoom.BackgroundLayers {
			var copyLayer room.RoomLayerBackground
			copyLayer.Config = new(room.RoomLayerConfig)
			*copyLayer.Config = *layer.Config
			editingRoom.BackgroundLayers = append(editingRoom.BackgroundLayers, &copyLayer)
		}
	}
	roomEditor.editorChangeRoom(editingRoom)
}

func (roomEditor *roomEditor) loadRoom(name string) {
	if roomEditor.hasUnsavedChanges {
		roomEditor.setStatusText("Cannot load room with unsaved changes. Restart the game and editor.")
		return
	}
	roomFilepath := roomEditor.roomDirectory + name
	if _, err := os.Stat(roomFilepath); os.IsNotExist(err) {
		roomEditor.setStatusText("Room \"" + name + "\" does not exist.")
		return
	}
	editingRoom := LoadRoom(name)
	if editingRoom == nil {
		roomEditor.setStatusText("Room \"" + name + "\" could not be loaded.")
		return
	}
	roomEditor.editorChangeRoom(editingRoom)
}

func (roomEditor *roomEditor) newLayerAndSelected(editingRoom *room.Room, text string, kind room.RoomLayerKind) {
	text = strings.TrimSpace(text)
	config := &room.RoomLayerConfig{
		Kind: kind,
		UUID: reditor.UUID(),
		Name: text,
	}
	switch kind {
	case room.RoomLayerKind_Instance:
		layer := &room.RoomLayerInstance{Config: config}
		editingRoom.InstanceLayers = append(editingRoom.InstanceLayers, layer)
		roomEditor.editingLayer = layer
	case room.RoomLayerKind_Background:
		layer := &room.RoomLayerBackground{Config: config}
		editingRoom.BackgroundLayers = append(editingRoom.BackgroundLayers, layer)
		roomEditor.editingLayer = layer
	case room.RoomLayerKind_Sprite:
		layer := &room.RoomLayerSprite{Config: config}
		editingRoom.SpriteLayers = append(editingRoom.SpriteLayers, layer)
		roomEditor.editingLayer = layer
	default:
		panic("Unhandled room layer kind: %s" + kind.String())
	}
	roomEditor.editorConfigSave()
}

func (roomEditor *roomEditor) setStatusText(text string) {
	roomEditor.statusText = text
	roomEditor.statusTimer = time.Now()
	println(text)
}

func inputSelectPressed() bool {
	return KeyboardCheckPressed(VkEnter) || KeyboardCheckPressed(VkNumpadEnter)
}

func hasFilterMatch(s string, filterBy string) bool {
	return filterBy == "" ||
		strings.Index(s, filterBy) >= 0
}

func (roomEditor *roomEditor) editorConfigLoad() {
	roomEditor.editingLayer = nil
	configPath := debugConfigPath("room_editor")
	fileData, err := file.OpenFile(configPath)
	if err == nil {
		bytes, err := ioutil.ReadAll(fileData)
		if err != nil {
			panic("Error loading " + configPath + "\n" + "Error: " + err.Error())
		}
		editorConfig := roomEditorConfig{}
		if err := json.Unmarshal(bytes, &editorConfig); err != nil {
			panic("Error unmarshalling " + configPath + "\n" + "Error: " + err.Error())
		}
		// Set layer from config
		for _, layer := range roomEditor.layers() {
			if layer.GetConfig().UUID == editorConfig.LayerSelected {
				roomEditor.editingLayer = layer
				break
			}
		}
		// Set brush from config
		switch roomEditor.editingLayer.(type) {
		case *room.RoomLayerSprite:
			obj := sprite.SpriteLoadByName(editorConfig.BrushSelected)
			if obj.Name() == editorConfig.BrushSelected {
				roomEditor.spriteSelected = obj
				break
			}
		case *room.RoomLayerInstance:
			for _, obj := range roomEditor.objectIndexToData {
				if obj != nil &&
					obj.ObjectName() == editorConfig.BrushSelected {
					roomEditor.objectSelected = obj
					break
				}
			}
		}
	} else {
		println("No editor config exists: " + configPath)
	}

}

func (roomEditor *roomEditor) editorConfigSave() {
	var editorConfig roomEditorConfig
	if roomEditor.editingLayer != nil {
		editorConfig.LayerSelected = roomEditor.editingLayer.GetConfig().UUID
		switch roomEditor.editingLayer.(type) {
		case *room.RoomLayerInstance:
			if roomEditor.objectSelected != nil {
				editorConfig.BrushSelected = roomEditor.objectSelected.ObjectName()
			}
		case *room.RoomLayerSprite:
			if roomEditor.spriteSelected.IsValid() {
				editorConfig.BrushSelected = roomEditor.spriteSelected.Name()
			}
		case *room.RoomLayerBackground:
			// no-op
		}
	}
	if roomEditor.editingRoom != nil {
		editorConfig.RoomSelected = roomEditor.editingRoom.Config.UUID
	}

	json, _ := json.MarshalIndent(editorConfig, "", "\t")
	configPath := debugConfigPath("room_editor")
	err := ioutil.WriteFile(configPath, json, 0644)
	if err != nil {
		println("Failed to write room editor config: " + configPath + "\n" + "Error: " + err.Error())
	}
}

func editorSave() {
	roomEditor := gRoomEditor
	editingRoom := roomEditor.editingRoom
	if editingRoom == nil {
		return
	}
	roomDirectory := editingRoom.Filepath()
	if !roomEditor.hasUnsavedChanges {
		roomEditor.setStatusText("No changes made to:" + roomDirectory)
		return
	}

	// Create room directory
	if _, err := os.Stat(roomDirectory); os.IsNotExist(err) {
		os.Mkdir(roomDirectory, 0700)
	}

	// Write config
	{
		json, _ := json.MarshalIndent(editingRoom.Config, "", "\t")
		err := ioutil.WriteFile(roomDirectory+"/config.json", json, 0644)
		if err != nil {
			panic("Failed to write room config.json file: " + roomDirectory + "\n" + "Error: " + err.Error())
		}
	}

	// NOTE(Jake): 2018-07-21
	//
	// Not sure how I should handle errors yet...
	// will need to see how it plays out in practice.
	//
	didErrorOccur := false

	// Handle instance layers
	{
		for _, layer := range roomEditor.layers() {
			config := layer.GetConfig()
			layerDirectory := roomDirectory + "/" + config.UUID

			// Create layer directory if it doesn't exist.
			if _, err := os.Stat(layerDirectory); os.IsNotExist(err) {
				os.Mkdir(layerDirectory, 0700)
			}

			// Write config
			{
				json, _ := json.MarshalIndent(config, "", "\t")
				err := ioutil.WriteFile(layerDirectory+"/config.json", json, 0644)
				if err != nil {
					break
				}
			}

			switch layer := layer.(type) {
			case nil:
				// no-op
			case *room.RoomLayerInstance:
				// Deleted instances
				if deletedInstances := layer.DeletedInstances; len(deletedInstances) > 0 {
					// NOTE(Jake): 2018-06-10
					//
					// Clear out 'DeletedInstances' before we save
					// the map file as data.
					//
					// We don't want this information
					// serialized.
					//
					layer.DeletedInstances = nil
					for _, obj := range deletedInstances {
						uuid := obj.UUID
						filepath := layerDirectory + "/" + uuid + ".txt"
						if _, err := os.Stat(filepath); os.IsNotExist(err) {
							// Ignore if it does not exist
							continue
						}
						err := os.Remove(filepath)
						if err != nil {
							println("Error deleting instance:", err.Error())
							didErrorOccur = true
							continue
						}
					}
				}

				// Write instances
				if instances := layer.Instances; len(instances) > 0 {
					for _, obj := range instances {
						data := roomEditorObjectIndexToData(obj.ObjectIndex)
						uuid := obj.UUID
						filepath := layerDirectory + "/" + uuid + ".txt"

						file, err := os.Create(filepath)
						if err != nil {
							println("Error writing file:", err.Error())
							didErrorOccur = true
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
						// but I'm going to err' on the side of caution here until I decide to investigate
						// this further.
						//
						for i := 0; i < 4; i++ {
							w.WriteString(uuid)
						}
						w.Flush()

						file.Close()
					}
				}
			case *room.RoomLayerBackground:
				// Write everything
				{
					// NOTE(Jake): 2018-07-29
					//
					// Ideally I'd want to disable the config data being saved into this
					// as we already get that information from config.json.
					//
					json, _ := json.MarshalIndent(layer, "", "\t")
					err := ioutil.WriteFile(layerDirectory+"/background.json", json, 0644)
					if err != nil {
						break
					}
				}
			case *room.RoomLayerSprite:
				// Deleted instances
				if deletedRecords := layer.DeletedSprites; len(deletedRecords) > 0 {
					// NOTE(Jake): 2018-06-10
					//
					// Clear out 'DeletedInstances' before we save
					// the map file as data.
					//
					// We don't want this information
					// serialized.
					//
					layer.DeletedSprites = nil
					for _, obj := range deletedRecords {
						uuid := obj.UUID
						filepath := layerDirectory + "/" + uuid + ".txt"
						if _, err := os.Stat(filepath); os.IsNotExist(err) {
							// Ignore if it does not exist
							continue
						}
						err := os.Remove(filepath)
						if err != nil {
							println("Error deleting instance:", err.Error())
							didErrorOccur = true
							continue
						}
					}
				}

				// Write instances
				if records := layer.Sprites; len(records) > 0 {
					for _, obj := range records {
						uuid := obj.UUID
						filepath := layerDirectory + "/" + uuid + ".txt"

						file, err := os.Create(filepath)
						if err != nil {
							println("Error writing sprite object file:", err.Error())
							didErrorOccur = true
							// todo(jake): 2018-06-10
							//
							// Error recovery here?
							//

							continue
						}

						name := obj.SpriteName
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
						// but I'm going to err' on the side of caution here until I decide to investigate
						// this further.
						//
						for i := 0; i < 4; i++ {
							w.WriteString(uuid)
						}
						w.Flush()

						file.Close()
					}
				}
			default:
				panic(fmt.Sprintf("Unimplemented save logic for layer type: %T", layer))
			}
		}
	}

	// Handle deletion of layers
	if deletedLayers := editingRoom.DeletedLayers; len(deletedLayers) > 0 {
		// NOTE(Jake): 2018-08-05
		//
		// Clear deleted layers out so they aren't serialized into the binary map
		//
		editingRoom.DeletedLayers = editingRoom.DeletedLayers[:0]
		for _, layerUUID := range deletedLayers {
			layerDirectory := roomDirectory + "/" + layerUUID
			os.RemoveAll(layerDirectory)
		}
	}

	// Save data copy of file
	editingRoom.DebugWriteDataFile(editingRoom.Filepath())

	// Save room!
	baseName := filepath.Base(editingRoom.Filepath())
	if didErrorOccur {
		roomEditor.setStatusText(baseName + " saved but with errors... Check console logs.")
	} else {
		roomEditor.setStatusText(baseName + " saved successfully!")
	}

	roomEditor.hasUnsavedChanges = false
}
