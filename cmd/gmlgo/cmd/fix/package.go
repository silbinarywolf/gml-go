package fix

import (
	"go/ast"
	"go/types"
)

type Package struct {
	dir      string
	name     string
	files    []*ast.File
	typesPkg *types.Package
}

// File holds a single parsed file and associated package / type data.
type File struct {
	pkg  *Package  // Package to which this file belongs.
	file *ast.File // Parsed AST.
}
