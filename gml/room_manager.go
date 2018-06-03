package gml

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/hajimehoshi/ebiten/ebitenutil"
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
	roomDataPath := AssetsDirectory() + "/room/" + name + ".data"
	dataFile, err := ebitenutil.OpenFile(roomDataPath)
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
	log.Println("Loaded map from data file")
	return result, nil
}
