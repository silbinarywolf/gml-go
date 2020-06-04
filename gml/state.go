package gml

import (
	"strconv"
)

var (
	gState       *state       = newState()
	gGameGlobals *gameGlobals = new(gameGlobals)
)

// todo: rename to world
type state struct {
	instanceManager          instanceManager
	instancesMarkedForDelete []InstanceIndex
	isCreatingRoomInstance   bool
	pauseCallback            func() bool
}

type gameGlobals struct {
	hasGameEnded bool
	// frameCount is how many draw calls have executed since
	// the application started, ie. not skipped by ebiten
	frameCount int
	// tickCount is how many update ticks have executed since
	// the application started
	tickCount uint64
	// frameUpdateBudgetNanosecondsUsed is used to calc a percentage
	// of how much time was spent updating and rendering the frame
	// out of how much time or budget you have
	frameUpdateBudgetNanosecondsUsed int64
}

func newState() *state {
	s := new(state)
	s.instanceManager.instanceIndexToIndex = make(map[InstanceIndex]int)
	return s
}

// DebugFrameUsage returns a string like "1% (55ns)" to tell you how much
// of your frame budget has been utilized.
func DebugFrameUsage() string {
	frameBudgetUsed := gGameGlobals.frameUpdateBudgetNanosecondsUsed
	timePerFrame := (1000000000 / float64(DesignedTPS())) // 1 second divided by 60 TPS, (1 = 1000000000 in nanoseconds)
	timeTaken := float64(frameBudgetUsed) / timePerFrame
	text := strconv.FormatFloat(timeTaken*100, 'f', 6, 64)
	return text + "% (" + strconv.Itoa(int(frameBudgetUsed)) + "ns)"
}

// DebugTickCount is incremented by 1 per update() call
func DebugTickCount() uint64 {
	return gGameGlobals.tickCount
}

// DebugFrameCount is similar to DebugTickCount but doesn't increase if ebiten skips the draw
func DebugFrameCount() int {
	return gGameGlobals.frameCount
}

// IsCreatingRoomInstance is to be used in the Create() event of your objects, this will only
// return true if the object is being created from room data, not code.
func IsCreatingRoomInstance() bool {
	return gState.isCreatingRoomInstance
}

// InstanceSetPauseCallback will call the provided function to check if instances
// should call their Update() method or not.
func InstanceSetPauseCallback(callback func() bool) {
	gState.pauseCallback = callback
}

// InstanceIsPaused will return whether executing the Update() method of instances is disabled.
// Update() method of instances can be disabled by utilizing InstanceSetPauseCallback()
func InstanceIsPaused() bool {
	return gState.pauseCallback != nil && gState.pauseCallback()
}

func (state *state) update() {
	if InstanceIsPaused() {
		return
	}

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
