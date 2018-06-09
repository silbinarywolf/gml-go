package gml

import (
	"io/ioutil"
	"strings"

	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/silbinarywolf/gml-go/gml/internal/file"
)

func AssetsDirectory() string {
	return file.AssetsDirectory
}

func ProgramDirectory() string {
	return file.ProgramDirectory
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
