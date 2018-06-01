// +build js

package gml

func LoadRoom(name string) *Map {
	manager := gRoomManager

	// Use already loaded asset
	if res, ok := manager.assetMap[name]; ok {
		return res
	}

	// Load from *.data file if it exists
	result, err := loadRoomFromDataFile(name)
	if err != nil {
		panic(err)
	}

	return result
}
