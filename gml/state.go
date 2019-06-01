package gml

import (
	"strconv"
)

var (
	gState *state = newState()
)

type state struct {
	instanceManager          instanceManager
	instancesMarkedForDelete []InstanceIndex
	isCreatingRoomInstance   bool
	//gWidth                     int
	gHeight                    int
	frameBudgetNanosecondsUsed int64
}

func newState() *state {
	s := new(state)
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

func (state *state) update() {
	// Simulate each active instance
	for i := 0; i < len(state.instanceManager.instances); i++ {
		inst := state.instanceManager.instances[i]
		baseObj := inst.BaseObject()
		if baseObj.internal.IsDestroyed {
			continue
		}
		// NOTE(Jake): 2019-04-03
		// Tested against Game Maker Studio 2, 2.2.2.326
		// It updates ImageIndex by ImageSpeed *before* the Begin Step
		baseObj.SpriteState.ImageUpdate()

		inst.Update()
	}
}

func (state *state) removeDeletedEntities() {
	manager := &state.instanceManager
	for _, instanceIndex := range state.instancesMarkedForDelete {
		dataIndex, ok := manager.instanceIndexToIndex[instanceIndex]
		if !ok {
			continue
		}

		// Remove from room instance draw list
		{
			roomInstanceIndex := manager.instances[dataIndex].BaseObject().RoomInstanceIndex()
			roomInst := roomGetInstance(roomInstanceIndex)
			if roomInst != nil {
				instances := roomInst.instances
				if len(instances) == 1 {
					if instanceIndex == instances[0] {
						instances = instances[:len(instances)-1]
					}
				} else {
					for dataIndex, otherInstanceIndex := range instances {
						if instanceIndex == otherInstanceIndex {
							lastEntry := instances[len(instances)-1]
							instances[dataIndex] = lastEntry
							instances = instances[:len(instances)-1]
							break
						}
					}
				}
				roomInst.instances = instances
			}
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