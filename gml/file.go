package gml

import (
	"io/ioutil"
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
	assetsDirectory = ProgramDirectory() + "/assets"
}
