package room

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/silbinarywolf/gml-go/gml/internal/file"
)

const (
	RoomDirectoryBase = "room"
)

var (
	gRoomManager = newRoomManager()
)

type RoomManager struct {
	assetMap map[string]*Room
	assetDir string
}

func newRoomManager() *RoomManager {
	return &RoomManager{
		assetMap: make(map[string]*Room),
	}
}

func loadRoomFromDataFile(name string) (*Room, error) {
	start := time.Now()
	roomDataPath := file.AssetDirectory + "/" + RoomDirectoryBase + "/" + name + ".data"
	dataFile, err := file.OpenFile(roomDataPath)
	if err != nil {
		return nil, err
	}
	roomData, err := ioutil.ReadAll(dataFile)
	dataFile.Close()
	if err != nil {
		panic(fmt.Errorf("Unable to read map data file into bytes: %s", name))
	}
	result := new(Room)
	err = result.Unmarshal(roomData)
	if err != nil {
		panic(fmt.Errorf("Unable to load map data: %s", err))
	}
	elapsed := time.Since(start)
	println("Room \"" + name + "\" took " + elapsed.String() + " to load.")
	return result, nil
}
