package gml

import (
	"strconv"
)

var (
	gState *state = newState()
)

type state struct {
	//globalInstances            *roomInstanceManager
	instanceManager          instanceManager
	roomInstances            []roomInstance
	instancesMarkedForDelete []InstanceIndex
	isCreatingRoomInstance   bool
	//gWidth                     int
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

func (state *state) createNewRoomInstance() *roomInstance {
	state.roomInstances = append(state.roomInstances, roomInstance{
		used: true,
	})
	state.isCreatingRoomInstance = true
	defer func() {
		state.isCreatingRoomInstance = false
	}()
	index := len(state.roomInstances) - 1
	roomInst := &state.roomInstances[index]
	roomInst.index = RoomInstanceIndex(index)

	// Create default instance layer if...
	// - No instance layers exist in the room data
	// - Creating blank room
	roomInst.instanceLayers = make([]roomInstanceLayerInstance, 1)
	roomInst.instanceLayers[0] = roomInstanceLayerInstance{
		index: 0,
	}
	roomInst.drawLayers = append(roomInst.drawLayers, &roomInst.instanceLayers[0])

	// If creating room programmatically, default the room size
	// to the size of the screen
	roomInst.Size = WindowSize()

	return roomInst
}

func (state *state) deleteRoomInstance(roomInst *roomInstance) {
	for _, layer := range roomInst.instanceLayers {
		// NOTE(Jake): 2018-08-21
		// Running Destroy() on each rather than InstanceDestroy()
		// for speed purposes
		for _, instanceIndex := range layer.instances {
			if inst := instanceIndex.Get(); inst != nil {
				inst.Destroy()
				cameraInstanceDestroy(instanceIndex)
			}
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

		// Remove from room instance draw list
		{
			roomInstanceIndex := manager.instances[dataIndex].BaseObject().RoomInstanceIndex()
			roomInstance := roomGetInstance(roomInstanceIndex)
			roomInstanceInstances := roomInstance.instanceLayers[0].instances
			if len(roomInstanceInstances) == 1 {
				if instanceIndex == roomInstanceInstances[0] {
					roomInstanceInstances = roomInstanceInstances[:len(roomInstanceInstances)-1]
				}
			} else {
				for dataIndex, otherInstanceIndex := range roomInstanceInstances {
					if instanceIndex == otherInstanceIndex {
						lastEntry := roomInstanceInstances[len(roomInstanceInstances)-1]
						roomInstanceInstances[dataIndex] = lastEntry
						roomInstanceInstances = roomInstanceInstances[:len(manager.instances)-1]
						break
					}
				}
			}
			roomInstance.instanceLayers[0].instances = roomInstanceInstances
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
