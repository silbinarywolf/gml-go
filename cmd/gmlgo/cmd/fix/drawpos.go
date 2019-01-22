package fix

import (
	"fmt"
	"go/ast"
	"go/types"
	"log"
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
				//var packagePath string
				switch x := fun.X.(type) {
				case *ast.Ident:
					scope := f.pkg.typesPkg.Scope().Innermost(x.NamePos)
					scope, obj := scope.LookupParent(x.Name, x.NamePos)
					pkgName, ok := obj.(*types.PkgName)
					if !ok {
						break
					}
					if pkgName.Imported().Path() == "github.com/silbinarywolf/gml-go/gml" {
						panic("continue this")
					}
					log.Printf("%v -- %v -- %v\n", packageName, pkgName.Imported().Path())
					//log.Printf("hilda: %v\n", x)
				case *ast.SelectorExpr:
					// this is for receiver calls only I *think*
				default:
					panic(fmt.Sprintf("Unhandled call expr using: %T\n", x))
				}
				/*if packageName == "gml" {
					log.Printf("%s %v %s\n", fun.Sel.Name, n.Args, packageName)
				}*/
			default:
				panic(fmt.Sprint("Unhandled type at root: %T", n))
			}
		}
		return true
	})
	return false
}
