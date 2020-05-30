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
	{"serialization_private_error", serialization_private_error_out},
	{"duplicate_asset_name", duplicate_asset_name},
}

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
