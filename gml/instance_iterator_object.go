package gml

/*type instanceIteratorObjectState struct {
	roomInstanceIndex int
	layerIndex        int
	instanceIndex     int
}

// Iterate over instances that belong to the same room as provided argument
func InstancesIteratorObject(inst object.ObjectType) instanceIteratorObjectState {
	if inst == nil ||
		reflect.ValueOf(inst).IsNil() {
		return instanceIteratorObjectState{
			roomInstanceIndex: -1,
		}
	}
	baseObj := inst.BaseObject()
	roomInstanceIndex := baseObj.RoomInstanceIndex()
	return instanceIteratorObjectState{
		instanceIndex:     -1,
		roomInstanceIndex: roomInstanceIndex,
	}
}

func (iterator *instanceIteratorObjectState) Next() bool {
	if iterator.roomInstanceIndex == -1 {
		return false
	}
	roomInst := &gState.roomInstances[iterator.roomInstanceIndex]
	if iterator.layerIndex >= len(roomInst.instanceLayers) {
		return false
	}
loop:
	iterator.instanceIndex++
	layer := &roomInst.instanceLayers[iterator.layerIndex]
	for iterator.instanceIndex < len(layer.manager.instances) {
		if !object.IsDestroyed(layer.manager.instances[iterator.instanceIndex].BaseObject()) {
			return true
		}
		iterator.instanceIndex++
	}
	iterator.instanceIndex = 0
	iterator.layerIndex++
	if iterator.layerIndex < len(roomInst.instanceLayers) {
		goto loop
	}
	return false
}

func (iterator *instanceIteratorObjectState) Value() object.ObjectType {
	return gState.roomInstances[iterator.roomInstanceIndex].instanceLayers[iterator.layerIndex].manager.instances[iterator.instanceIndex]
}*/
