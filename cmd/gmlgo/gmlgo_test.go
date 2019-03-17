package main_test

import (
	"os"
	"testing"

	"github.com/silbinarywolf/gml-go/cmd/gmlgo/internal/build"
	"github.com/silbinarywolf/gml-go/cmd/gmlgo/internal/generate"
)

// TestBuild runs the tests in testdata/*.*.
func TestBuild(t *testing.T) {
	build.Cmd.Run(build.Cmd, []string{"./testdata/game_with_one_object"})
	if _, err := os.Stat("./testdata/game_with_one_object/game/gmlgo_gen.go"); os.IsNotExist(err) {
		t.Errorf("%s", err)
		return
	}
	os.Remove("./testdata/game_with_one_object/game/gmlgo_gen.go")
}

// TestGenerate runs the tests in testdata/*.*.
func TestGenerate(t *testing.T) {
	generate.Cmd.Run(generate.Cmd, []string{"./testdata/game_with_one_object"})
	if _, err := os.Stat("./testdata/game_with_one_object/game/gmlgo_gen.go"); os.IsNotExist(err) {
		t.Errorf("%s", err)
		return
	}
	os.Remove("./testdata/game_with_one_object/game/gmlgo_gen.go")
}
