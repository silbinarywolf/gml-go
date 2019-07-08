package gml

import (
	"sort"

	"github.com/silbinarywolf/gml-go/gml/internal/geom"
)

type RoomInstanceIndex int32

const roomUndefined RoomInstanceIndex = 0

type roomInstanceStateManager struct {
	roomInstances          []roomInstance
	lastCreatedRoom        RoomInstanceIndex
	isCreatingRoomInstance bool
}

type roomInstance struct {
	used  bool
	index RoomInstanceIndex
	//room  *room.Room // deprecate room data
	geom.Rect

	instances []InstanceIndex
}

var roomInstanceState = roomInstanceStateManager{
	roomInstances: make([]roomInstance, 1, 10),
}

// RoomInstanceNew create a new empty room instance programmatically
func RoomInstanceNew() RoomInstanceIndex {
	roomInstanceState.roomInstances = append(roomInstanceState.roomInstances, roomInstance{
		used: true,
	})
	roomInstanceState.isCreatingRoomInstance = true
	defer func() {
		roomInstanceState.isCreatingRoomInstance = false
	}()
	index := len(roomInstanceState.roomInstances) - 1
	roomInst := &roomInstanceState.roomInstances[index]
	roomInst.index = RoomInstanceIndex(index)

	// If creating room programmatically, default the room size
	// to the size of the screen
	roomInst.Size = WindowSize()

	roomInstanceState.lastCreatedRoom = roomInst.index

	return roomInst.index
}

func (roomInstanceIndex RoomInstanceIndex) RoomInstanceChangeRoom(baseObj *Object) {
	roomInst := &roomInstanceState.roomInstances[roomInstanceIndex]
	if !roomInst.used {
		return
	}
	oldRoomInstanceIndex := baseObj.RoomInstanceIndex()
	if oldRoomInstanceIndex == 0 {
		baseObj.internal.RoomInstanceIndex = roomInstanceIndex
		roomInst.instances = append(roomInst.instances, baseObj.InstanceIndex())
		return
	}
	if oldRoomInstanceIndex == baseObj.internal.RoomInstanceIndex {
		return
	}
	// NOTE(Jake): 2018-07-22
	// For now instances default to the last instance layer
	//layerIndex := len(roomInst.instanceLayers) - 1
	//layer := &roomInst.instanceLayers[layerIndex]

	//instanceRemove(inst)
	panic("todo: Update this to remove instance index from one room instance list and add it to another")
	// Move entity to new list
	//index := len(manager.instances)
	//moveInstance(inst, index, roomInstanceIndex, layerIndex)
	//manager.instances = append(manager.instances, inst)
}

func (roomInstanceIndex RoomInstanceIndex) InstanceCreate(x, y float64, objectIndex ObjectIndex) ObjectType {
	return instanceCreate(x, y, objectIndex, func(inst *Object) {
		inst.internal.RoomInstanceIndex = roomInstanceIndex
		roomInst := &roomInstanceState.roomInstances[roomInstanceIndex]
		roomInst.instances = append(roomInst.instances, inst.InstanceIndex())
	}, true)
}

// Destroy destroys a room instance
func (roomInstanceIndex RoomInstanceIndex) Destroy() {
	if roomInst := roomGetInstance(roomInstanceIndex); roomInst != nil {
		// NOTE(Jake): 2018-08-21
		// Running Destroy() on each rather than InstanceDestroy()
		// for speed purposes
		for _, instanceIndex := range roomInst.instances {
			if inst := instanceIndex.Get(); inst != nil {
				fastInstanceDestroy(inst)
			}
		}
		if roomInstanceState.lastCreatedRoom == roomInst.index {
			roomInstanceState.lastCreatedRoom = 0
		}
		roomInst.instances = nil
		roomInst.used = false
		gState.instancesMarkedForDelete = gState.instancesMarkedForDelete[:0]
		*roomInst = roomInstance{}
		return
	}
	panic("Invalid roomInstanceIndex given")
}

func (roomInstanceIndex RoomInstanceIndex) SetSize(width, height float64) {
	if roomInst := roomGetInstance(roomInstanceIndex); roomInst != nil {
		roomInst.Size.X = width
		roomInst.Size.Y = height
		return
	}
	panic("Invalid roomInstanceIndex given")
}

// Size returns the size of the given room instance
func (roomInstanceIndex RoomInstanceIndex) Size() geom.Vec {
	if roomInst := roomGetInstance(roomInstanceIndex); roomInst != nil {
		return roomInst.Size
	}
	panic("Invalid roomInstanceIndex given")
}

// WithAll returns a list of instances in the same room as the provided object
func (roomIndex RoomInstanceIndex) WithAll() []InstanceIndex {
	roomInst := roomGetInstance(roomIndex)
	if roomInst == nil {
		panic("Cannot call WithAll() on room that doesn't exist")
	}
	var list []InstanceIndex
	for _, otherIndex := range roomInst.instances {
		other := otherIndex.getBaseObject()
		if other == nil {
			continue
		}
		list = append(list, otherIndex)
	}
	if len(list) == 0 {
		return nil
	}
	return list
}

func roomGetInstance(roomInstanceIndex RoomInstanceIndex) *roomInstance {
	roomInst := &roomInstanceState.roomInstances[roomInstanceIndex]
	if roomInst.used {
		return roomInst
	}
	return nil
}

func roomLastCreated() *roomInstance {
	if roomInstanceState.lastCreatedRoom == 0 {
		return nil
	}
	return roomGetInstance(roomInstanceState.lastCreatedRoom)
}

func (roomInst *roomInstance) draw() {
	// Sort by order
	sort.SliceStable(roomInst.instances, func(i, j int) bool {
		a := roomInst.instances[i].getBaseObject()
		if a == nil {
			return false
		}
		b := roomInst.instances[j].getBaseObject()
		if b == nil {
			return false
		}
		return a.Depth() > b.Depth()
	})
	//log.Printf("Stable sort count: %d, cap: %d\n", len(layer.instances), cap(layer.instances))

	for _, instanceIndex := range roomInst.instances {
		inst := instanceIndex.Get()
		if inst == nil {
			panic("instance index not removed from draw list when destroyed")
		}
		inst.Draw()
	}

	// DrawTextF(16, 16, "Instance Debug Draw Count: %v", len(roomInst.instances))
}
