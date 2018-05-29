// +build js

package gml

func LoadMap(name string) *Map {
	manager := g_mapManager

	// Use already loaded asset
	if res, ok := manager.assetMap[name]; ok {
		return res
	}

	// Load from *.data file if it exists
	result, err := loadMapFromData(name)
	if err != nil {
		panic(err)
	}

	return result
}
