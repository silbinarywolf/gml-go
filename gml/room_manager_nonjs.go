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
	"github.com/silbinarywolf/gml-go/gml/internal/object"
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
	x64, err := strconv.ParseInt(strings.TrimSpace(scanner.Text()), 10, 32)
	x := int32(x64)
	if err != nil {
		log.Printf("Error parsing Y of entity %s.\n", entityName)
		return
	}

	// Y
	scanner.Scan()
	y64, err := strconv.ParseInt(strings.TrimSpace(scanner.Text()), 10, 32)
	y := int32(y64)
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

	// Set room dimensions
	{
		// NOTE(Jake): 2018-06-02
		//
		// Probably a slow hack to get the entity size
		// for building map data on-fly, but whatever!
		//
		inst := object.NewRawInstance(objectIndex, 0, 0)
		inst.Create()

		baseObj := inst.BaseObject()
		x := int32(x)
		y := int32(y)
		width := int32(baseObj.Size.X)
		height := int32(baseObj.Size.Y)

		if x < room.Left {
			room.Left = x
		}
		right := x + width
		if right > room.Right {
			room.Right = right
		}
		if y < room.Top {
			room.Top = y
		}
		bottom := y + height
		if bottom > room.Bottom {
			room.Bottom = bottom + height
		}
	}

	basename := filepath.Base(instancePath)
	filename := strings.TrimSuffix(basename, filepath.Ext(basename))

	// Increase entity counter
	{
		filenameParts := strings.Split(filename, "_")
		if len(filenameParts) == 3 {
			id := filenameParts[len(filenameParts)-1]
			count, err := strconv.ParseInt(id, 10, 64)
			if err == nil {
				if count > room.UserEntityCount {
					username := filenameParts[len(filenameParts)-2]
					if username == roomEditorUsername() {
						room.UserEntityCount = count
					}
				}
			} else {
				println(filename, ": Skipping, Error parsing the last part (entity ID) after splitting by _")
			}
		} else {
			println(filename, ": Expected to split into 3 parts, not", len(filenameParts))
		}
	}

	room.Instances = append(room.Instances, &RoomObject{
		Filename:    filename,
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

	roomPath := AssetsDirectory() + "/room/" + name

	// Read entities
	instancePathList := make([]string, 0, 1000)
	err := filepath.Walk(roomPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			println("prevent panic by handling failure accessing a path " + roomPath + ": " + err.Error())
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
	room.Filepath = roomPath
	room.Instances = make([]*RoomObject, 0, len(instancePathList))
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
	//
	if debugMode {
		go func() {
			err := room.writeDataFile(roomPath)
			if err != nil {
				panic("Failed writing " + roomPath + ", error: " + err.Error())
			}
		}()
	}
	return room
}
