// +build js

package gml

import (
	"fmt"
	"testing"

	"github.com/gopherjs/gopherjs/js"
)

// NOTE(Jake): 2018-09-09
// Can't find the documentation but this was needed for GopherJS
func TestMain(m *testing.M) {
	i := m.Run()

	js.Global.Call("eval", fmt.Sprintf("window.$GopherJSTestResult = %v", i))
}
