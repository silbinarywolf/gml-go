package file

import (
	"io"
)

var (
	AssetDirectory string = "▲not-set▲"

	// todo(Jake): 2018-11-24
	// Think of a better name? ProgramPath?
	// The name should work as both a full URL (web output) and full directory path.
	ProgramDirectory string
)

const (
	AssetDirectoryBase = "asset"
)

// ReadSeekCloser is io.ReadSeeker and io.Closer.
type readSeekCloser interface {
	io.ReadSeeker
	io.Closer
}

func InitAssetDir() {
	if ProgramDirectory == "" {
		ProgramDirectory = computeProgramDirectory()
	}
	if AssetDirectory == "" {
		AssetDirectory = ProgramDirectory + "/" + AssetDirectoryBase
	}
}

func SetAssetDir(dir string) {
	AssetDirectory = dir
}
