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

func (m *Map) writeDataFile(mapDir string) error {
	data, err := m.Marshal()
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(mapDir+".data", data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (m *Map) readEntity(entityFilepath string) {
	fmt.Printf("Loading %s...\n", entityFilepath)
	entityFileData, err := ebitenutil.OpenFile(entityFilepath)
	if err != nil {
		panic(fmt.Errorf("Unable to find map entity file: %s", err))
	}
	bytesData, err := ioutil.ReadAll(entityFileData)
	entityFileData.Close()
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
	m.Entities = append(m.Entities, &MapEntity{
		ObjectIndex: int32(objectIndex),
		X:           x,
		Y:           y,
	})
}

func LoadMap(name string) *Map {
	manager := g_mapManager

	// Use already loaded asset
	if res, ok := manager.assetMap[name]; ok {
		return res
	}

	// Load from *.data file if it exists
	//if mapDataFile, _ := loadMapFromData(name); mapDataFile != nil {
	//	manager.assetMap[name] = mapDataFile
	//	return mapDataFile
	//}

	mapDir := currentDirectory() + "/assets/map/" + name

	// Read entities
	entityFilepaths := make([]string, 0, 1000)
	err := filepath.Walk(mapDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", mapDir, err)
			return err
		}
		if info.IsDir() {
			return nil
		}
		entityFilepaths = append(entityFilepaths, path)
		return nil
	})
	if err != nil {
		panic(err)
	}

	//
	m := new(Map)
	m.Entities = make([]*MapEntity, 0, 10)
	for _, entity := range entityFilepaths {
		m.readEntity(entity)
	}
	manager.assetMap[name] = m

	// NOTE(Jake): 2018-05-29
	//
	// Hack to write map data out to file from binary clients
	// so that web clients can load it.
	//
	// Write out *.data file (for browsers / fast client loading)
	go func() {
		err := m.writeDataFile(mapDir)
		if err != nil {
			log.Printf("Failed writing %s\nerror: %s", mapDir, err)
		}
	}()
	return m
}
