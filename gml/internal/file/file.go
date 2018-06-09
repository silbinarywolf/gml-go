package file

import (
	"os"
)

var (
	AssetsDirectory string = "▲not-set▲"
)

func init() {
	// NOTE(Jake): 2018-06-03
	//
	// Allow setting asset dir via environment variable for `go test` support
	//
	AssetsDirectory = os.Getenv("GML_ASSET_DIR")
	if AssetsDirectory != "" {
		AssetsDirectory = AssetsDirectory + "/assets"
	} else {
		AssetsDirectory = ProgramDirectory + "/assets"
	}
}
