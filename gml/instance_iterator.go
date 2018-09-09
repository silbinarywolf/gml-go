package gml

import (
	"github.com/silbinarywolf/gml-go/gml/internal/object"
)

type instanceIteratorState struct {
	roomInstanceIndex int
	layerIndex        int
	instanceIndex     int
}

func InstancesIterator(inst object.ObjectType) instanceIteratorState {
	roomInstanceIndex := object.RoomInstanceIndex(inst.BaseObject())
	return instanceIteratorState{
		instanceIndex:     -1,
		roomInstanceIndex: roomInstanceIndex,
	}
}

func (iterator *instanceIteratorState) Next() bool {
	roomInst := &gState.roomInstances[iterator.roomInstanceIndex]
	if iterator.layerIndex >= len(roomInst.instanceLayers) {
		return false
	}
	layer := &roomInst.instanceLayers[iterator.layerIndex]
	iterator.instanceIndex++
	for {
		if iterator.instanceIndex < len(layer.manager.instances) {
			return true
		}
		iterator.instanceIndex = 0
		iterator.layerIndex++
		if iterator.layerIndex < len(roomInst.instanceLayers) {
			continue
		}
		return false
	}
}

func (iterator *instanceIteratorState) Value() object.ObjectType {
	return gState.roomInstances[iterator.roomInstanceIndex].instanceLayers[iterator.layerIndex].manager.instances[iterator.instanceIndex]
}
