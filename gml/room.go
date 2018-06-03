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
	for _, obj := range room.Instances {
		roomInst.InstanceCreate(V(float64(obj.X), float64(obj.Y)), ObjectIndex(obj.ObjectIndex))
	}
	return roomInst
}
