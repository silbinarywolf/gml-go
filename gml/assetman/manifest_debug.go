// +build debug

package assetman

import (
	"encoding/json"
	"io/ioutil"

	"github.com/silbinarywolf/gml-go/gml/internal/file"
)

func debugWriteManifest() {
	manifest := make(map[string]map[string]string)
	for _, manager := range assetManagers {
		name, data := manager.ManifestJSON()
		manifest[name] = data
	}

	data, err := json.MarshalIndent(manifest, "", "  ")
	if err != nil {
		panic(err)
	}

	if err := ioutil.WriteFile(file.AssetDirectory+"/manifest.json", data, 0755); err != nil {
		panic(err)
	}
}
