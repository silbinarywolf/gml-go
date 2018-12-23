package gml

/*type roomInstanceIteratorState struct {
	roomInstanceIndex int
}

func RoomInstanceIterator() roomInstanceIteratorState {
	return roomInstanceIteratorState{
		roomInstanceIndex: 0,
	}
}

func (iterator *roomInstanceIteratorState) Next() bool {
	for {
		roomInst := &gState.roomInstances[iterator.roomInstanceIndex]
		if iterator.roomInstanceIndex >= len(gState.roomInstances) {
			return false
		}
		if !roomInst.used {
			iterator.roomInstanceIndex++
			continue
		}
		iterator.roomInstanceIndex++
		return true
	}
}

func (iterator *roomInstanceIteratorState) Value() int {
	return iterator.roomInstanceIndex
}
*/
