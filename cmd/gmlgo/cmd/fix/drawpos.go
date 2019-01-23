package fix

import (
	"fmt"
	"go/ast"
	"go/types"
	"strings"
)

func init() {
	register(drawPosFix)
}

// drawPosFix is the first fix rule created and was mainly built as a PoC before the 1st official 1.0 release
// it can probably be deleted once other fixes are introduced.
var drawPosFix = fix{
	name: "drawPos",
	date: "2019-01-20",
	f:    drawposfix,
	desc: `Change all Draw* functions to pass "x, y" instead of geom.Vec for positions. Github Issue #81`,
}

func drawposfix(f *File) bool {
	astFile := f.file
	ast.Inspect(astFile, func(n ast.Node) bool {
		switch n := n.(type) {
		case *ast.CallExpr:
			switch fun := n.Fun.(type) {
			case *ast.SelectorExpr:
				switch x := fun.X.(type) {
				case *ast.Ident:
					scope := f.pkg.typesPkg.Scope().Innermost(x.NamePos)
					scope, obj := scope.LookupParent(x.Name, x.NamePos)
					pkgName, ok := obj.(*types.PkgName)
					if !ok {
						break
					}
					switch pkgName.Imported().Path() {
					case "github.com/silbinarywolf/gml-go/gml":
						if strings.HasPrefix(fun.Sel.Name, "DrawSelf") {
							fmt.Printf("%s\n", fun)
							for _, arg := range n.Args {
								arg, ok := arg.(*ast.UnaryExpr)
								if !ok {
									panic(fmt.Sprintf("%T\n", arg))
								}
								x := arg.X.(*ast.SelectorExpr)
								panic(fmt.Sprintf("%s\n", x.Sel.String()))
							}
							panic(fmt.Sprintf("continue this: %s %s", fun.Sel.Name, n.Args))
						}
					}
				case *ast.SelectorExpr:
					// this is for receiver calls only I *think*
				default:
					panic(fmt.Sprintf("Unhandled call expr using: %T\n", x))
				}
			case *ast.Ident:
				// ignore "new()"
			default:
				panic(fmt.Sprintf("Unhandled type at root: %T", fun))
			}
		}
		return true
	})
	return false
}
