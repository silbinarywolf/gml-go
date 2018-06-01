// +build darwin freebsd linux windows
// +build !js
// +build !android
// +build !ios

package gml

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/hajimehoshi/ebiten/ebitenutil"
)

func (room *Room) writeDataFile(roomPath string) error {
	data, err := room.Marshal()
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(roomPath+".data", data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (room *Room) readInstance(instancePath string) {
	println("Loading ", instancePath, "...")
	instanceFileData, err := ebitenutil.OpenFile(instancePath)
	if err != nil {
		panic(fmt.Errorf("Unable to find map entity file: %s", err))
	}
	bytesData, err := ioutil.ReadAll(instanceFileData)
	instanceFileData.Close()
	if err != nil {
		panic(fmt.Errorf("Unable to find map entity file: Read all: %s\n", err))
	}
	bytesReader := bytes.NewReader(bytesData)
	scanner := bufio.NewScanner(bytesReader)

	// Entity name
	scanner.Scan()
	entityName := strings.TrimSpace(scanner.Text())

	// X
	scanner.Scan()
	x, err := strconv.ParseInt(strings.TrimSpace(scanner.Text()), 10, 64)
	if err != nil {
		log.Printf("Error parsing Y of entity %s.\n", entityName)
		return
	}

	// Y
	scanner.Scan()
	y, err := strconv.ParseInt(strings.TrimSpace(scanner.Text()), 10, 64)
	if err != nil {
		log.Printf("Error parsing X of entity %s.\n", entityName)
		return
	}
	if err := scanner.Err(); err != nil {
		log.Printf("Error parsing entity, error: %s.\n", err)
		return
	}
	objectIndex, ok := ObjectGetIndex(entityName)
	if !ok {
		log.Printf("Missing mapping of name \"%s\" to entity ID. Is this name defined in your gml.Init()?", entityName)
		return
	}
	room.Instances = append(room.Instances, &RoomInstance{
		ObjectIndex: int32(objectIndex),
		X:           x,
		Y:           y,
	})
}

func LoadRoom(name string) *Room {
	manager := gRoomManager

	// Use already loaded asset
	if res, ok := manager.assetMap[name]; ok {
		return res
	}

	// Load from *.data file if it exists
	//if mapDataFile, _ := loadMapFromData(name); mapDataFile != nil {
	//	manager.assetMap[name] = mapDataFile
	//	return mapDataFile
	//}

	roomPath := WorkingDirectory() + "/assets/room/" + name

	// Read entities
	instancePathList := make([]string, 0, 1000)
	err := filepath.Walk(roomPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", roomPath, err)
			return err
		}
		if info.IsDir() {
			return nil
		}
		instancePathList = append(instancePathList, path)
		return nil
	})
	if err != nil {
		panic(err)
	}

	//
	room := new(Room)
	room.Instances = make([]*RoomInstance, 0, 10)
	for _, instance := range instancePathList {
		room.readInstance(instance)
	}
	manager.assetMap[name] = room

	// NOTE(Jake): 2018-05-29
	//
	// Hack to write map data out to file from binary clients
	// so that web clients can load it.
	//
	// Write out *.data file (for browsers / fast client loading)
	go func() {
		err := room.writeDataFile(roomPath)
		if err != nil {
			log.Printf("Failed writing %s\nerror: %s", roomPath, err)
		}
	}()
	return room
}
