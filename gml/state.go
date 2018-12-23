package gml

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/silbinarywolf/gml-go/gml/internal/geom"
	"github.com/silbinarywolf/gml-go/gml/internal/room"
	"github.com/silbinarywolf/gml-go/gml/internal/sprite"
)

var (
	gState *state = newState()
	g_game gameState
)

type gameState struct {
	hasGameRestarted bool
}

func GameRestart() {
	g_game.hasGameRestarted = true
}

type state struct {
	//globalInstances            *roomInstanceManager
	instanceManager            instanceManager
	roomInstances              []roomInstance
	instancesMarkedForDelete   []InstanceIndex
	isCreatingRoomInstance     bool
	gWidth                     int
	gHeight                    int
	frameBudgetNanosecondsUsed int64
}

func newState() *state {
	s := new(state)
	s.roomInstances = make([]roomInstance, 1, 10)
	s.instanceManager.instanceIndexToIndex = make(map[InstanceIndex]int)
	return s
}

// FrameUsage returns a string like "1% (55ns)" to tell you how much
// of your frame budget has been utilized. (Assumes 60FPS)
func FrameUsage() string {
	frameBudgetUsed := gState.frameBudgetNanosecondsUsed
	timeTaken := float64(frameBudgetUsed) / 16000000.0
	//fmt.Printf("Time used: %v / 16000000.0\n", frameBudgetUsed)
	text := strconv.FormatFloat(timeTaken*100, 'f', 6, 64)
	return text + "% (" + strconv.Itoa(int(gState.frameBudgetNanosecondsUsed)) + "ns)"
}

// IsCreatingRoomInstance is to be used in the Create() event of your objects, this will only
// return true if the object is being created from room data, not code.
func IsCreatingRoomInstance() bool {
	return gState.isCreatingRoomInstance
}

func (state *state) createNewRoomInstance(room *room.Room) *roomInstance {
	state.roomInstances = append(state.roomInstances, roomInstance{
		used: true,
		room: room,
	})
	state.isCreatingRoomInstance = true
	defer func() {
		state.isCreatingRoomInstance = false
	}()
	index := len(state.roomInstances) - 1
	roomInst := &state.roomInstances[index]
	roomInst.index = RoomInstanceIndex(index)
	roomInst.size = geom.Size{
		X: int32(WindowWidth()),
		Y: int32(WindowHeight()),
	}

	if room == nil ||
		len(room.InstanceLayers) == 0 {
		// Create default instance layer if...
		// - No instance layers exist in the room data
		// - Creating blank room
		roomInst.instanceLayers = make([]roomInstanceLayerInstance, 1)
		roomInst.instanceLayers[0] = roomInstanceLayerInstance{
			index: 0,
		}
		roomInst.drawLayers = append(roomInst.drawLayers, &roomInst.instanceLayers[0])
	}

	// If non-blank room instance, use room data to create
	if roomInst.room != nil {
		roomInst.size = geom.Size{
			X: roomInst.room.Right - roomInst.room.Left,
			Y: roomInst.room.Bottom - roomInst.room.Top,
		}

		// Instance layers
		if len(room.InstanceLayers) > 0 {
			roomInst.instanceLayers = make([]roomInstanceLayerInstance, len(room.InstanceLayers))
			for i := 0; i < len(room.InstanceLayers); i++ {
				layerData := room.InstanceLayers[i]
				roomInst.instanceLayers[i] = roomInstanceLayerInstance{
					index: i,
				}
				layer := &roomInst.instanceLayers[i]
				layer.drawOrder = layerData.Config.Order
				for _, obj := range layerData.Instances {
					InstanceCreate(float64(obj.X), float64(obj.Y), roomInst.index, ObjectIndex(obj.ObjectIndex))
					fmt.Printf("todo(Jake): 2018-12-19: Fix room instance creation to create on correct layer.\n")
				}
				roomInst.drawLayers = append(roomInst.drawLayers, layer)
			}
		}
		// Background layers
		for i := 0; i < len(room.BackgroundLayers); i++ {
			layerData := room.BackgroundLayers[i]
			spriteName := layerData.SpriteName
			if spriteName == "" {
				continue
			}
			layer := new(roomInstanceLayerBackground)
			layer.x = float64(layerData.X)
			layer.y = float64(layerData.Y)
			layer.roomLeft = float64(room.Left)
			layer.roomRight = float64(room.Right)
			layer.sprite = sprite.SpriteLoadByName(spriteName)
			layer.drawOrder = layerData.Config.Order
			roomInst.drawLayers = append(roomInst.drawLayers, layer)
		}
		// Sprite layers
		for i := 0; i < len(room.SpriteLayers); i++ {
			layerData := room.SpriteLayers[i]
			hasCollision := layerData.Config.HasCollision
			layer := roomInstanceLayerSprite{}
			layer.hasCollision = hasCollision
			layer.sprites = make([]roomInstanceLayerSpriteObject, 0, len(layerData.Sprites))
			for _, sprObj := range layerData.Sprites {
				// Add draw sprite
				spr := sprite.SpriteLoadByName(sprObj.SpriteName)
				record := roomInstanceLayerSpriteObject{
					sprite: spr,
				}
				record.X = float64(sprObj.X)
				record.Y = float64(sprObj.Y)
				layer.sprites = append(layer.sprites, record)
			}
			layer.drawOrder = layerData.Config.Order
			roomInst.spriteLayers = append(roomInst.spriteLayers, layer)
			roomInst.drawLayers = append(roomInst.drawLayers, &roomInst.spriteLayers[len(roomInst.spriteLayers)-1])
		}
		// Sort draw layers by order
		sort.Slice(roomInst.drawLayers, func(i, j int) bool {
			return roomInst.drawLayers[i].order() < roomInst.drawLayers[j].order()
		})
	}
	return roomInst
}

func (state *state) deleteRoomInstance(roomInst *roomInstance) {
	for _, layer := range roomInst.instanceLayers {
		// NOTE(Jake): 2018-08-21
		//
		// Running Destroy() on each rather than InstanceDestroy()
		// for speed purposes
		//
		for _, instanceIndex := range layer.instances {
			inst := InstanceGet(instanceIndex)
			if inst == nil {
				continue
			}
			inst.Destroy()
			cameraInstanceDestroy(instanceIndex)
		}
		layer.instances = nil
	}

	roomInst.used = false
	*roomInst = roomInstance{}
}

func (state *state) update() {
	// Simulate each active instance
	for i := 0; i < len(state.instanceManager.instances); i++ {
		inst := state.instanceManager.instances[i]
		baseObj := inst.BaseObject()

		inst.Update()
		baseObj.SpriteState.ImageUpdate()
	}

	// Remove deleted entities
	manager := &state.instanceManager
	for _, instanceIndex := range state.instancesMarkedForDelete {
		dataIndex, ok := manager.instanceIndexToIndex[instanceIndex]
		if !ok {
			continue
		}
		if dataIndex == len(manager.instances)-1 {
			// Remove last entry
			delete(manager.instanceIndexToIndex, instanceIndex)
			manager.instances = manager.instances[:len(manager.instances)-1]
			continue
		}
		// Swap deleted entry for last entry
		delete(manager.instanceIndexToIndex, instanceIndex)
		lastEntry := manager.instances[len(manager.instances)-1]
		manager.instances[dataIndex] = lastEntry
		manager.instanceIndexToIndex[lastEntry.BaseObject().InstanceIndex()] = dataIndex
		manager.instances = manager.instances[:len(manager.instances)-1]
	}

	state.instancesMarkedForDelete = state.instancesMarkedForDelete[:0]
}
