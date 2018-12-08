package gml

import (
	"sort"
	"strconv"

	"github.com/silbinarywolf/gml-go/gml/internal/geom"
	"github.com/silbinarywolf/gml-go/gml/internal/object"
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
	globalInstances            *instanceManager
	roomInstances              []RoomInstance
	instancesMarkedForDelete   []object.ObjectType
	isCreatingRoomInstance     bool
	gWidth                     int
	gHeight                    int
	frameBudgetNanosecondsUsed int64
}

func newState() *state {
	return &state{
		globalInstances: newInstanceManager(),
		roomInstances:   make([]RoomInstance, 1, 10),
	}
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

// IsCreatingRoomInstance returns whether this instance was created by a room or not, rather
// than programmatically.
func IsCreatingRoomInstance() bool {
	return gState.isCreatingRoomInstance
}

func (state *state) createNewRoomInstance(room *room.Room) *RoomInstance {
	state.roomInstances = append(state.roomInstances, RoomInstance{
		used: true,
		room: room,
	})
	state.isCreatingRoomInstance = true
	defer func() {
		state.isCreatingRoomInstance = false
	}()
	index := len(state.roomInstances) - 1
	roomInst := &state.roomInstances[index]
	roomInst.index = index

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
					instanceCreateLayer(geom.Vec{float64(obj.X), float64(obj.Y)}, layer, roomInst, object.ObjectIndex(obj.ObjectIndex))
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

func (state *state) deleteRoomInstance(roomInst *RoomInstance) {
	for _, layer := range roomInst.instanceLayers {
		// NOTE(Jake): 2018-08-21
		//
		// Running Destroy() on each rather than InstanceDestroy()
		// for speed purposes
		//
		for _, inst := range layer.manager.instances {
			//InstanceDestroy()
			inst.Destroy()
			cameraInstanceDestroy(inst)
		}
		layer.manager.reset()
	}

	roomInst.used = false
	*roomInst = RoomInstance{}
}

func (state *state) update(animationUpdate bool) {
	// Simulate global instances
	state.globalInstances.update(animationUpdate)

	// Simulate each instance in each room instance
	for i := 1; i < len(state.roomInstances); i++ {
		roomInst := &state.roomInstances[i]
		if !roomInst.used {
			continue
		}
		roomInst.update(animationUpdate)
	}

	// Remove deleted entities
	for _, inst := range state.instancesMarkedForDelete {
		instanceRemove(inst)
	}
	state.instancesMarkedForDelete = state.instancesMarkedForDelete[:0]
}
