package gml

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var (
	assetsDirectory string = "▲not-set▲"
)

func AssetsDirectory() string {
	return assetsDirectory
}

func ReadFileAsString(path string) (string, error) {
	path = AssetsDirectory() + "/" + path
	fileData, err := ebitenutil.OpenFile(path)
	if err != nil {
		return "", err
	}
	bytesData, err := ioutil.ReadAll(fileData)
	fileData.Close()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(bytesData)), nil
}

func init() {
	// NOTE(Jake): 2018-06-03
	//
	// Allow setting asset dir via environment variable for `go test` support
	//
	assetsDirectory = os.Getenv("GML_ASSET_DIR")
	if assetsDirectory != "" {
		assetsDirectory = assetsDirectory + "/assets"
	} else {
		assetsDirectory = ProgramDirectory() + "/assets"
	}
}
