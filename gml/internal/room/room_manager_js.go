// +build js

package room

func LoadRoom(name string) *Room {
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
