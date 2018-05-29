package gml

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var g_mapManager = newMapManager()

type MapManager struct {
	assetMap map[string]*Map
}

func newMapManager() MapManager {
	manager := MapManager{}
	manager.assetMap = make(map[string]*Map)
	return manager
}

func loadMapFromData(name string) (*Map, error) {
	mapDir := currentDirectory() + "/assets/map/" + name
	mapDataFile, err := ebitenutil.OpenFile(mapDir + ".data")
	if err != nil {
		return nil, err
	}
	mapData, err := ioutil.ReadAll(mapDataFile)
	mapDataFile.Close()
	if err != nil {
		panic(fmt.Errorf("Unable to read map data file into bytes: %s", name))
	}
	result := new(Map)
	err = result.Unmarshal(mapData)
	if err != nil {
		panic(fmt.Errorf("Unable to load map data: %s", err))
	}
	log.Println("Loaded map from data file")
	return result, nil
}
