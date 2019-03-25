package generate

import (
	"strings"
	"testing"
)

// Golden represents a test case.
type Golden struct {
	name   string
	input  string // input; the package clause is provided when running the test.
	output string // exected output.
}

var golden = []Golden{
	{"simple", simple_in, simple_out},
	{"deep_1", deep_1_in, deep_1_out},
	{"deep_2", deep_2_in, deep_2_out},
}

// Simple test: enumeration of type int starting at 0.
const simple_in = `
import (
	"github.com/silbinarywolf/gml-go/gml"
)

type GameObjectA struct {
	gml.Object
}
`

const simple_out = `
// Code generated by "gmlgo"
// 0.1.0
// DO NOT EDIT. DO NOT COMMIT TO YOUR VCS REPOSITORY.

package test

import (
	"github.com/silbinarywolf/gml-go/gml"
	"github.com/silbinarywolf/gml-go/gml/audio"
)

// Silence errors if audio is unused
var _ = audio.InitSoundGeneratedData

const (
	ObjGameObjectA gml.ObjectIndex = 1
)

var _gen_Obj_index_to_name = []string{
	ObjGameObjectA: "GameObjectA",
}

var _gen_Obj_name_to_index = map[string]gml.ObjectIndex{
	"GameObjectA": ObjGameObjectA,
}

var _gen_Obj_index_to_data = []gml.ObjectType{
	ObjGameObjectA: new(GameObjectA),
}

func init() {
	gml.InitObjectGeneratedData(_gen_Obj_index_to_name, _gen_Obj_name_to_index, _gen_Obj_index_to_data)
}
`

// Deep 1: test embedded a struct that is embedding gml.Object
const deep_1_in = `
import (
	"github.com/silbinarywolf/gml-go/gml"
)

type GameObject struct {
	gml.Object
}

type GameObjectA struct {
	GameObject
}
`

const deep_1_out = `
// Code generated by "gmlgo"
// 0.1.0
// DO NOT EDIT. DO NOT COMMIT TO YOUR VCS REPOSITORY.

package test

import (
	"github.com/silbinarywolf/gml-go/gml"
	"github.com/silbinarywolf/gml-go/gml/audio"
)

// Silence errors if audio is unused
var _ = audio.InitSoundGeneratedData

const (
	ObjGameObject  gml.ObjectIndex = 1
	ObjGameObjectA gml.ObjectIndex = 2
)

var _gen_Obj_index_to_name = []string{
	ObjGameObject:  "GameObject",
	ObjGameObjectA: "GameObjectA",
}

var _gen_Obj_name_to_index = map[string]gml.ObjectIndex{
	"GameObject":  ObjGameObject,
	"GameObjectA": ObjGameObjectA,
}

var _gen_Obj_index_to_data = []gml.ObjectType{
	ObjGameObject:  new(GameObject),
	ObjGameObjectA: new(GameObjectA),
}

func init() {
	gml.InitObjectGeneratedData(_gen_Obj_index_to_name, _gen_Obj_name_to_index, _gen_Obj_index_to_data)
}
`

// Deep 2: test embedded a struct that is embedding struct that is embedding gml.Object
const deep_2_in = `
import (
	"github.com/silbinarywolf/gml-go/gml"
)

type GameObject struct {
	gml.Object
}

type SuperGameObject struct {
	GameObject
}

type GameObjectA struct {
	SuperGameObject
}
`

const deep_2_out = `
// Code generated by "gmlgo"
// 0.1.0
// DO NOT EDIT. DO NOT COMMIT TO YOUR VCS REPOSITORY.

package test

import (
	"github.com/silbinarywolf/gml-go/gml"
	"github.com/silbinarywolf/gml-go/gml/audio"
)

// Silence errors if audio is unused
var _ = audio.InitSoundGeneratedData

const (
	ObjGameObject      gml.ObjectIndex = 1
	ObjGameObjectA     gml.ObjectIndex = 2
	ObjSuperGameObject gml.ObjectIndex = 3
)

var _gen_Obj_index_to_name = []string{
	ObjGameObject:      "GameObject",
	ObjGameObjectA:     "GameObjectA",
	ObjSuperGameObject: "SuperGameObject",
}

var _gen_Obj_name_to_index = map[string]gml.ObjectIndex{
	"GameObject":      ObjGameObject,
	"GameObjectA":     ObjGameObjectA,
	"SuperGameObject": ObjSuperGameObject,
}

var _gen_Obj_index_to_data = []gml.ObjectType{
	ObjGameObject:      new(GameObject),
	ObjGameObjectA:     new(GameObjectA),
	ObjSuperGameObject: new(SuperGameObject),
}

func init() {
	gml.InitObjectGeneratedData(_gen_Obj_index_to_name, _gen_Obj_name_to_index, _gen_Obj_index_to_data)
}
`

func TestGolden(t *testing.T) {
	for _, test := range golden {
		g := Generator{}
		input := "package test\n" + test.input
		file := test.name + ".go"
		g.parsePackage(".", []string{file}, input)
		g.generate()
		got := strings.TrimSpace(string(g.format()))
		expected := strings.TrimSpace(test.output)
		if got != expected {
			t.Errorf("%s: got\n====\n%s\n====\nexpected\n====\n%s", test.name, got, expected)
		}
	}
}