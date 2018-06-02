package gml

import (
	"io/ioutil"
	"strings"

	"github.com/hajimehoshi/ebiten/ebitenutil"
)

func ReadFileAsString(path string) (string, error) {
	path = WorkingDirectory() + "/assets/" + path
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
