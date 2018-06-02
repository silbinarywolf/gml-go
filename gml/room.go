package gml

var (
	roomInstances []*RoomInstance
)

//type RoomInstanceID int

func (room *Room) Create() *RoomInstance {
	roomInst := &RoomInstance{
		room: room,
	}
	//roomIndex := RoomInstanceID(len(roomInstances))
	roomInstances = append(roomInstances, roomInst)

	// Instantiate instances for this room
	for _, inst := range room.Instances {
		roomInst.InstanceCreate(V(float64(inst.X), float64(inst.Y)), ObjectIndex(inst.ObjectIndex))
	}
	return roomInst
}
