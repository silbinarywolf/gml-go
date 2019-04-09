package main_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/silbinarywolf/gml-go/cmd/gmlgo/internal/generate"
)

// Golden represents a test case.
type Golden struct {
	name   string
	output string
}

var golden = []Golden{
	{"one_object", one_object_out},
	{"deep_one", deep_one_out},
	{"deep_two", deep_two_out},
	{"serialization_simple", serialization_simple_out},
	{"serialization_private_error", serialization_private_error_out},
	{"duplicate_asset_name", duplicate_asset_name},
}

const one_object_out = `// Code generated by "gmlgo"
// 0.1.0
// DO NOT EDIT. DO NOT COMMIT TO YOUR VCS REPOSITORY.

package game

import (
	"github.com/silbinarywolf/gml-go/gml"
)

const (
	ObjPlayer gml.ObjectIndex = 1
)

var _gen_Obj_index_to_name = []string{
	ObjPlayer: "Player",
}

var _gen_Obj_name_to_index = map[string]gml.ObjectIndex{
	"Player": ObjPlayer,
}

var _gen_Obj_index_to_data = []gml.ObjectType{
	ObjPlayer: new(Player),
}

func init() {
	gml.InitObjectGeneratedData(_gen_Obj_index_to_name, _gen_Obj_name_to_index, _gen_Obj_index_to_data)
}
`

const deep_one_out = `// Code generated by "gmlgo"
// 0.1.0
// DO NOT EDIT. DO NOT COMMIT TO YOUR VCS REPOSITORY.

package game

import (
	"github.com/silbinarywolf/gml-go/gml"
)

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

const deep_two_out = `// Code generated by "gmlgo"
// 0.1.0
// DO NOT EDIT. DO NOT COMMIT TO YOUR VCS REPOSITORY.

package game

import (
	"github.com/silbinarywolf/gml-go/gml"
)

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

const serialization_simple_out = `// Code generated by "gmlgo"
// 0.1.0
// DO NOT EDIT. DO NOT COMMIT TO YOUR VCS REPOSITORY.

package game

import (
	"bytes"
	"encoding/binary"

	"github.com/silbinarywolf/gml-go/gml"
)

const (
	ObjSerializablePlayer gml.ObjectIndex = 1
)

var _gen_Obj_index_to_name = []string{
	ObjSerializablePlayer: "SerializablePlayer",
}

var _gen_Obj_name_to_index = map[string]gml.ObjectIndex{
	"SerializablePlayer": ObjSerializablePlayer,
}

var _gen_Obj_index_to_data = []gml.ObjectType{
	ObjSerializablePlayer: new(SerializablePlayer),
}

func init() {
	gml.InitObjectGeneratedData(_gen_Obj_index_to_name, _gen_Obj_name_to_index, _gen_Obj_index_to_data)
}
func (self SerializablePlayer) UnsafeSnapshotMarshalBinaryRoot(buf *bytes.Buffer) error {
	if err := self.Object.UnsafeSnapshotMarshalBinary(buf); err != nil {
		return err
	}
	if err := binary.Write(buf, binary.LittleEndian, int64(self.defaultInt)); err != nil {
		return err
	}
	return nil
}
func (self *SerializablePlayer) UnsafeSnapshotUnmarshalBinaryRoot(buf *bytes.Buffer) error {
	if err := self.Object.UnsafeSnapshotUnmarshalBinary(buf); err != nil {
		return err
	}
	{
		var d int64
		if err := binary.Read(buf, binary.LittleEndian, &d); err != nil {
			return err
		}
		self.defaultInt = int(d)
	}
	return nil
}
`

const serialization_private_error_out = `self.Embed.privateInt32 not exported. Cannot generate serialization code if using struct with unexported fields.`

const duplicate_asset_name = `Cannot have duplicate asset names:
- sprite/folder_a/SprHero
- sprite/folder_b/SprHero`

// TestBuild runs the tests in testdata/*.*.
func TestBuild(t *testing.T) {
	for _, n := range golden {
		n := n // this is needed or t.Parallel() will be buggy - https://gist.github.com/posener/92a55c4cd441fc5e5e85f27bca008721
		t.Run(n.name, func(t *testing.T) {
			t.Parallel()
			projectDir := "./testdata/" + n.name
			genFile := projectDir + "/game/gmlgo_gen.go"
			expected := n.output

			os.Remove(genFile)

			skipBecauseGotExpectedError := false
			func() {
				defer func() {
					if r := recover(); r != nil {
						if got := fmt.Sprintf("%v", r); got != expected {
							t.Errorf("%s: got\n====\n%s\n====\nexpected\n====\n%s", genFile, got, expected)
							return
						}
						skipBecauseGotExpectedError = true
						return
					}
				}()
				generate.Cmd.Run(generate.Cmd, []string{projectDir})
			}()
			if skipBecauseGotExpectedError {
				return
			}

			bytes, err := ioutil.ReadFile(genFile)
			if err != nil {
				t.Errorf("%s", err)
				return
			}
			got := string(bytes)
			if got != expected {
				t.Errorf("%s: got\n====\n%s\n====\nexpected\n====\n%s", genFile, got, expected)
				return
			}
			os.Remove(genFile)
		})
	}
}

// TestGenerate runs the tests in testdata/*.*.
/*func TestGenerate(t *testing.T) {
	generate.Cmd.Run(generate.Cmd, []string{"./testdata/game_with_one_object"})
	if _, err := os.Stat("./testdata/game_with_one_object/game/gmlgo_gen.go"); os.IsNotExist(err) {
		t.Errorf("%s", err)
		return
	}
	os.Remove("./testdata/game_with_one_object/game/gmlgo_gen.go")
}
*/
