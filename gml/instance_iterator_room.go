package gml

// Iterate over instances in the provided room
func InstancesIteratorRoom(roomInstanceIndex int) instanceIteratorObjectState {
	return instanceIteratorObjectState{
		instanceIndex:     -1,
		roomInstanceIndex: roomInstanceIndex,
	}
}
