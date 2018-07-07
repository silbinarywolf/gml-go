package gml

import (
	"github.com/silbinarywolf/gml-go/gml/internal/object"
)

var (
	gState *state = newState()
)

type state struct {
	globalInstances *instanceManager
	roomInstances   []RoomInstance
}

func newState() *state {
	return &state{
		globalInstances: newInstanceManager(),
		roomInstances:   make([]RoomInstance, 1, 10),
	}
}

func (state *state) createNewRoomInstance(room *Room) *RoomInstance {
	state.roomInstances = append(state.roomInstances, RoomInstance{
		used: true,
		room: room,
	})
	index := len(state.roomInstances) - 1
	roomInst := &state.roomInstances[index]
	roomInst.index = index

	// If non-blank room instance, use room data to create
	if roomInst.room != nil {
		// Instantiate instances for this room
		for _, obj := range roomInst.room.Instances {
			roomInst.InstanceCreate(V(float64(obj.X), float64(obj.Y)), object.ObjectIndex(obj.ObjectIndex))
		}
	}
	return roomInst
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
}

func (state *state) draw() {
	// Render global instances
	state.globalInstances.draw()

	// Render each instance in each room instance
	for i := 1; i < len(state.roomInstances); i++ {
		roomInst := &state.roomInstances[i]
		if !roomInst.used {
			continue
		}
		roomInst.draw()
	}
}
