package file

import (
	"io"
)

const (
	AssetDirectoryBase  = "asset"
	assetDirectoryUnset = "▲not-set▲"
)

var (
	AssetDirectory string = assetDirectoryUnset

	// todo(Jake): 2018-11-24
	// Think of a better name? ProgramPath?
	// The name should work as both a full URL (web output) and full directory path.
	ProgramDirectory string
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
	AssetDirectory = ProgramDirectory + "/" + AssetDirectoryBase
}

// SetAssetDir allows the engine to override the default behaviour
// of
func SetAssetDir(dir string) {
	AssetDirectory = dir
}

func IsAssetDirSet() bool {
	return AssetDirectory != assetDirectoryUnset
}
