package gml

import "github.com/silbinarywolf/gml-go/gml/internal/object"

type instanceIterator struct {
	data  []object.ObjectType
	index int
	curr  object.ObjectType
}

func (manager *instanceManager) Iterator() instanceIterator {
	return instanceIterator{
		data:  manager.instances,
		index: 0,
	}
}

func (iterator *instanceIterator) Next() bool {
	if iterator.index >= len(iterator.data) {
		return false
	}
	iterator.curr = iterator.data[iterator.index]
	iterator.index++
	return true
}

func (iterator *instanceIterator) Current() object.ObjectType {
	return iterator.curr
}
