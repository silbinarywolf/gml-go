package fix

import (
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

const (
	Use              = "fix [dir]"
	ShortDescription = "Fix finds GML-Go programs that use old APIs and rewrites them to use newer ones."
)

type Arguments struct {
	Directory string // .
}

const (
	genFile        = "gmlgo_gen.go"
	packageName    = "gml"
	importPath     = "\"github.com/silbinarywolf/gml-go/gml\""
	vendorBaseName = "vendor"
	parserMode     = parser.ParseComments
)

type dirWalker struct {
	fileSet *token.FileSet
	goFiles []*File
}

func Run(args Arguments) {
	if args.Directory == "" {
		args.Directory = "."
	}
	dir := args.Directory
	fileSet := token.NewFileSet()
	//dir, err := filepath.Abs(dir)
	//if err != nil {
	//	log.Fatal(err)
	//}

	// getValidAndSortFiles
	var astFiles []*File
	{
		var dirInfo dirWalker

		dirInfo.fileSet = fileSet
		dirInfo.getValidFilesRecursive(dir)
		astFiles = dirInfo.goFiles
	}

	// sortFilesAlphabetically
	sort.Slice(astFiles[:], func(i, j int) bool {
		return astFiles[i].file.Name.Name < astFiles[j].file.Name.Name
	})

	// sortFixesByDate
	sort.Slice(fixes[:], func(i, j int) bool {
		return fixes[i].date < fixes[j].date
	})

	// typeCheck
	/*{
		config := types.Config{
			Importer:                 importer.Default(),
			FakeImportC:              true,
			DisableUnusedImportCheck: true,
		}
		for _, astFile := range astFiles {
			pkg, err := config.Check(dir, fileSet, []*ast.File{astFile}, nil)
			if err != nil {
				// todo: Catch all errors and` report nicely
				log.Printf("ignoring file due to parse error: %s\n", err)
				continue
			}
			lookup := pkg.Scope().Lookup("DrawSetFont")
			if lookup == nil {
				continue
			}

			log.Fatalf("%v\n", lookup.Type())
		}
		log.Fatal("todo: Keep workin on typeCheck")
	}*/

	//
	/*for _, astFile := range astFiles {
		for _, fix := range fixes {
			if fix.f(astFile) {
				// todo: Handle what happens if file is fixed
			}
		}
	}*/
}

// getValidFilesRecursive will recursively walk the directory structure using the following rules:
// - get all *.go files that import this library
// - ignore "vendor" directories
func (dirInfo *dirWalker) getValidFilesRecursive(dir string) {
	var dirs []string
	var goFiles []string

	// Get files and directories in two seperate lists
	{
		files, err := ioutil.ReadDir(dir)
		if err != nil {
			log.Fatal(err)
		}
		for _, info := range files {
			base := info.Name()
			path := dir + "/" + base
			if info.IsDir() {
				// Skip code that a user would not want to
				// upgrade by default such as third-party "vendor" folder
				if base == vendorBaseName {
					continue
				}
				dirs = append(dirs, path)
				continue
			}
			if !strings.HasSuffix(path, ".go") {
				continue
			}
			goFiles = append(goFiles, path)
		}
	}

	// Parse each *.go file in this dir
	if len(goFiles) > 0 {
		var astFiles []*ast.File

		//
		goPackage := &Package{}
		goPackage.dir = dir

		// Parse files in package
		usesThisLibrary := false
		for _, path := range goFiles {
			text, err := ioutil.ReadFile(path)
			if err != nil {
				log.Fatalf("cannot open *.go file: %s", err)
			}
			astFile, err := parser.ParseFile(dirInfo.fileSet, path, text, parserMode)
			if err != nil {
				log.Fatal(err)
			}
			usesThisLibrary := false
			for _, goImport := range astFile.Imports {
				usesThisLibrary = usesThisLibrary || goImport.Path.Value == importPath
			}
			astFiles = append(astFiles, astFile)
		}

		if !usesThisLibrary {
			// Exit early.
			// Keep parsing directories recursively
			for _, dir := range dirs {
				dirInfo.getValidFilesRecursive(dir)
			}
			return
		}

		//
		for _, astFile := range astFiles {
			dirInfo.goFiles = append(dirInfo.goFiles, &File{
				pkg:  goPackage,
				file: astFile,
			})
			goPackage.files = append(goPackage.files, astFile)
		}

		// Typecheck package
		{
			config := types.Config{
				Importer: importer.Default(),
				// NOTE(Jake): 2019-01-22: Might need to support this later
				//FakeImportC: true,
			}
			info := &types.Info{
				Defs: make(map[*ast.Ident]types.Object),
			}
			typesPkg, err := config.Check(goPackage.dir, dirInfo.fileSet, astFiles, info)
			if err != nil {
				log.Fatal(err)
			}
			goPackage.typesPkg = typesPkg
		}
	}

	// Keep parsing directories recursively
	for _, dir := range dirs {
		dirInfo.getValidFilesRecursive(dir)
	}
}