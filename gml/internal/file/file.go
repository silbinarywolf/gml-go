package file

import (
	"io"
	"os"
)

var (
	AssetDirectory string = "▲not-set▲"

	// todo(Jake): 2018-11-24
	// Think of a better name? ProgramPath?
	// The name should work as both a full URL (web output) and full directory path.
	ProgramDirectory string = computeProgramDirectory()
)

const (
	assetDirectoryBase = "asset"
)

// ReadSeekCloser is io.ReadSeeker and io.Closer.
type readSeekCloser interface {
	io.ReadSeeker
	io.Closer
}

func init() {
	// NOTE(Jake): 2018-06-03
	// Allow setting asset dir via environment variable for `go test` support
	AssetDirectory = os.Getenv("GML_ASSET_DIR")
	if AssetDirectory != "" {
		AssetDirectory = AssetDirectory + "/" + assetDirectoryBase
	} else {
		AssetDirectory = ProgramDirectory + "/" + assetDirectoryBase
	}
}
