package assetman

import (
	"encoding/json"
	"io/ioutil"

	"github.com/silbinarywolf/gml-go/gml/internal/file"
)

var manifest map[string]map[string]string

// UnsafeGetNameFromKey will get an asset name
// For internal use only. No backwards compatibility guaranteed.
func UnsafeGetNameFromKey(groupName string, key string) string {
	res, ok := manifest[groupName]
	if !ok {
		panic("Cannot find resource group: " + groupName)
	}
	name, ok := res[key]
	if !ok {
		panic("Cannot find resource by key: " + key)
	}
	return name
}

func loadManifest() {
	f, err := file.OpenFile(file.AssetDirectory + "/manifest.json")
	if err != nil {
		panic(err)
	}
	bytes, err := ioutil.ReadAll(f)
	f.Close()
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(bytes, &manifest); err != nil {
		panic(err)
	}
}
