package gmlgo

// NOTE(Jae): 2020-06-01
// We just have this at the root-level of the package to stop "go get" and etc.
// from complaining on Go 1.14 and onwards.
import (
	_ "github.com/silbinarywolf/gml-go/gml"
)
