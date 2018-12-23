// +build darwin freebsd linux windows
// +build !js
// +build !android
// +build !ios

package room

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/silbinarywolf/gml-go/gml/internal/file"
	"github.com/silbinarywolf/gml-go/gml/internal/geom"
	"github.com/silbinarywolf/gml-go/gml/internal/sprite"
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

//func loadInstanceLayer(layer *RoomLayerInstance) {
//
//}

func LoadRoom(name string) *Room {
	manager := gRoomManager

	// Use already loaded asset
	if res, ok := manager.assetMap[name]; ok {
		return res
	}

	// Load room
	start := time.Now()

	// Load from *.data file if it exists
	/*room, err := loadRoomFromDataFile(name)
	if err != nil {
		panic(err)
	}*/
	room := loadRoomFromDirectoryFiles(name)
	elapsed := time.Since(start)
	println("Room \"" + name + "\" took " + elapsed.String() + " to load.")

	manager.assetMap[name] = room

	// NOTE(Jake): 2018-05-29
	//
	// Hack to write map data out to file from binary clients
	// so that web clients can load it.
	//
	// Write out *.data file (for browsers / fast client loading)
	//
	room.DebugWriteDataFile(room.Filepath())

	return room
}

func loadRoomFromDirectoryFiles(name string) *Room {
	/*objectTypeToInitState := make(map[ObjectIndex]ObjectType)
	defer func() {
		for _, inst := range objectTypeToInitState {
			// NOTE(Jake): 2018-09-15
			// Cleanup entities or else they might stay alive on the server
			// ie. networked entities
			// I had to fix a bug where 26 enemies were "there" because this
			// Destroy() wasn't here.
			inst.Destroy()
		}
	}()*/

	//
	room := new(Room)
	room.Config = new(RoomConfig)
	// NOTE(Jake): 2018-07-21
	//
	// I might want UUID to be an actual UUID in the future
	// however can't do this yet as the folder name needs to be the room
	// name currently.
	//
	room.Config.UUID = name
	room.Config.Name = name
	//room.Instances = make([]*RoomObject, 0, len(instancePathList))
	roomPath := room.Filepath()
	{
		// Read config
		configPath := roomPath + "/config.json"
		fileData, err := file.OpenFile(configPath)
		if err != nil {
			panic("Failed to load config.json for room: " + configPath + "\n" + "Error: " + err.Error())
			//continue
		}
		bytes, err := ioutil.ReadAll(fileData)
		if err != nil {
			panic("Error loading load config.json for room: " + configPath + "\n" + "Error: " + err.Error())
		}
		if err := json.Unmarshal(bytes, room.Config); err != nil {
			panic("Error unmarshalling load config.json for room: " + configPath + "\n" + "Error: " + err.Error())
		}
	}
	{
		// Read layer directories
		var layerPathList []string
		err := filepath.Walk(roomPath, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				println("prevent panic by handling failure accessing a path " + roomPath + ": " + err.Error())
				return err
			}
			if !info.IsDir() {
				// Skip files
				return nil
			}
			if path == roomPath {
				// Skip self
				return nil
			}
			layerPathList = append(layerPathList, path)
			return nil
		})
		if err != nil {
			panic(err)
		}
		// Load each layer config
		for _, pathStr := range layerPathList {
			configPath := pathStr + "/config.json"
			fileData, err := file.OpenFile(configPath)
			if err != nil {
				panic("Failed to load config.json for layer: " + configPath + "\n" + "Error: " + err.Error())
				//continue
			}
			bytesData, err := ioutil.ReadAll(fileData)
			if err != nil {
				panic("Error loading load config.json for layer: " + configPath + "\n" + "Error: " + err.Error())
				//continue
			}
			layerConfig := new(RoomLayerConfig)
			if err := json.Unmarshal(bytesData, layerConfig); err != nil {
				panic("Error unmarshalling load config.json for layer: " + configPath + "\n" + "Error: " + err.Error())
				//continue
			}
			//fmt.Printf("%v", layerConfig.String())
			switch layerConfig.Kind {
			case RoomLayerKind_Instance:
				layer := new(RoomLayerInstance)
				layer.Config = layerConfig
				layerPath := roomPath + "/" + layer.Config.UUID

				// Read instance files
				var pathList []string
				err := filepath.Walk(layerPath, func(pathStr string, info os.FileInfo, err error) error {
					if err != nil {
						println("prevent panic by handling failure accessing a path " + roomPath + ": " + err.Error())
						return err
					}
					// Skip directories and only get *.txt files
					if info.IsDir() ||
						path.Ext(pathStr) != ".txt" {
						return nil
					}
					pathList = append(pathList, pathStr)
					return nil
				})
				if err != nil {
					panic(err)
				}

				// Read each individual file
				for _, path := range pathList {
					//println("Loading ", instancePath, "...")
					instanceFileData, err := file.OpenFile(path)
					if err != nil {
						panic("Unable to find map entity file: " + err.Error())
					}
					bytesData, err := ioutil.ReadAll(instanceFileData)
					instanceFileData.Close()
					if err != nil {
						panic("Unable to find map entity file: Read all: " + err.Error())
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
						println("Error parsing Y of entity", entityName)
						continue
					}

					// Y
					scanner.Scan()
					y64, err := strconv.ParseInt(strings.TrimSpace(scanner.Text()), 10, 32)
					y := int32(y64)
					if err != nil {
						println("Error parsing X of entity", entityName)
						continue
					}
					if err := scanner.Err(); err != nil {
						println("Error parsing entity, error: ", err.Error())
						continue
					}
					objectIndex := 1
					panic("todo: ObjectIndex fix this")
					/*objectIndex, ok := ObjectGetIndex(entityName)
					if !ok {
						println("Missing mapping of name \"" + entityName + "\" to entity ID. Is this name defined in your gml.Init()?")
						continue
					}*/

					// Set room dimensions
					{
						// NOTE(Jake): 2018-06-02
						//
						// Probably a slow hack to get the entity size
						// for building map data on-fly, but whatever!
						//
						// NOTE(Jake): 2018-09-15
						// Should definitely look into an entity having default
						// hitbox etc as Create() can cause bugs, like multiple
						// networked entities because I didn't call gml.Destroy()
						//
						/*inst, ok := objectTypeToInitState[objectIndex]
						if !ok {
							inst = New_Deprecated_RawInstance(objectIndex, 0, 0, 0)
							inst.Create()
							objectTypeToInitState[objectIndex] = inst
						}
						baseObj := inst.BaseObject()
						size := baseObj.Size
						*/
						size := geom.Size{}
						panic("todo: Fix this to get object width/height")

						x := int32(x)
						y := int32(y)
						width := int32(size.X)
						height := int32(size.Y)

						if x < room.Left {
							room.Left = x
						}
						if right := x + width; right > room.Right {
							room.Right = right
						}
						if y < room.Top {
							room.Top = y
						}
						if bottom := y + height; bottom > room.Bottom {
							room.Bottom = bottom + height
						}
					}

					basename := filepath.Base(path)
					uuid := strings.TrimSuffix(basename, filepath.Ext(basename))

					layer.Instances = append(layer.Instances, &RoomObject{
						UUID:        uuid,
						ObjectIndex: int32(objectIndex),
						X:           x,
						Y:           y,
					})
				}

				room.InstanceLayers = append(room.InstanceLayers, layer)
			case RoomLayerKind_Background:
				layer := new(RoomLayerBackground)
				layerPath := roomPath + "/" + layerConfig.UUID

				configPath := layerPath + "/background.json"
				fileData, err := file.OpenFile(configPath)
				if err != nil {
					panic("Failed to load background.json for layer: " + configPath + "\n" + "Error: " + err.Error())
				}
				bytesData, err := ioutil.ReadAll(fileData)
				if err != nil {
					panic("Error loading load background.json for layer: " + configPath + "\n" + "Error: " + err.Error())
				}
				if err := json.Unmarshal(bytesData, layer); err != nil {
					panic("Error unmarshalling load background.json for layer: " + configPath + "\n" + "Error: " + err.Error())
				}
				// NOTE(Jake): 2018-07-29
				//
				// "Config" data should ideally not be written into the above struct
				// so we're just ensuring here that config.json is the "source of truth"
				// for this data.
				//
				layer.Config = layerConfig
				room.BackgroundLayers = append(room.BackgroundLayers, layer)
			case RoomLayerKind_Sprite:
				layer := new(RoomLayerSprite)
				layer.Config = layerConfig
				layerPath := roomPath + "/" + layer.Config.UUID

				// Read instance files
				var pathList []string
				err := filepath.Walk(layerPath, func(pathStr string, info os.FileInfo, err error) error {
					if err != nil {
						println("prevent panic by handling failure accessing a path " + roomPath + ": " + err.Error())
						return err
					}
					// Skip directories and only get *.txt files
					if info.IsDir() ||
						path.Ext(pathStr) != ".txt" {
						return nil
					}
					pathList = append(pathList, pathStr)
					return nil
				})
				if err != nil {
					panic(err)
				}

				// Read each individual file
				for _, path := range pathList {
					//println("Loading ", instancePath, "...")
					instanceFileData, err := file.OpenFile(path)
					if err != nil {
						panic("Unable to find sprite object file: " + err.Error())
					}
					bytesData, err := ioutil.ReadAll(instanceFileData)
					instanceFileData.Close()
					if err != nil {
						panic("Unable to find sprite object file: Read all: " + err.Error())
					}
					bytesReader := bytes.NewReader(bytesData)
					scanner := bufio.NewScanner(bytesReader)

					// Entity name
					scanner.Scan()
					spriteName := strings.TrimSpace(scanner.Text())

					// X
					scanner.Scan()
					x64, err := strconv.ParseInt(strings.TrimSpace(scanner.Text()), 10, 32)
					x := int32(x64)
					if err != nil {
						println("Error parsing Y of sprite object", spriteName)
						continue
					}

					// Y
					scanner.Scan()
					y64, err := strconv.ParseInt(strings.TrimSpace(scanner.Text()), 10, 32)
					y := int32(y64)
					if err != nil {
						println("Error parsing X of sprite object", spriteName)
						continue
					}
					if err := scanner.Err(); err != nil {
						println("Error parsing sprite object, error: ", err.Error())
						continue
					}

					// Set room dimensions
					{
						spriteIndex := sprite.SpriteLoadByName(spriteName)
						if spriteIndex == sprite.SprUndefined {
							println("Error loading sprite sprite \"", spriteName, "\" error: ", err.Error())
							continue
						}
						x := int32(x)
						y := int32(y)
						size := spriteIndex.Size()
						width := int32(size.X)
						height := int32(size.Y)

						if x < room.Left {
							room.Left = x
						}
						if right := x + width; right > room.Right {
							room.Right = right
						}
						if y < room.Top {
							room.Top = y
						}
						if bottom := y + height; bottom > room.Bottom {
							room.Bottom = bottom + height
						}
					}

					basename := filepath.Base(path)
					uuid := strings.TrimSuffix(basename, filepath.Ext(basename))

					layer.Sprites = append(layer.Sprites, &RoomSpriteObject{
						UUID:       uuid,
						SpriteName: spriteName,
						X:          x,
						Y:          y,
					})
				}

				room.SpriteLayers = append(room.SpriteLayers, layer)
			default:
				panic("Unknown or unhandled layer kind: " + layerConfig.Kind.String() + "(" + strconv.Itoa(int(layerConfig.Kind)) + ")")
			}
		}
	}
	{
		// Sort each layer
		sort.Slice(room.InstanceLayers, func(i, j int) bool {
			return room.InstanceLayers[i].Config.Order < room.InstanceLayers[j].Config.Order
		})
		sort.Slice(room.BackgroundLayers, func(i, j int) bool {
			return room.BackgroundLayers[i].Config.Order < room.BackgroundLayers[j].Config.Order
		})
	}
	return room
}
