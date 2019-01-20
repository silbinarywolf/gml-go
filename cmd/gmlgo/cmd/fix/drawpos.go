package fix

import "go/ast"

func init() {
	register(drawPosFix)
}

var drawPosFix = fix{
	name: "drawPos",
	date: "2019-01-20",
	f:    drawposfix,
	desc: `Change all Draw* functions to pass "x, y" instead of geom.Vec for positions. Github Issue #81`,
}

func drawposfix(f *ast.File) bool {
	return false
}
