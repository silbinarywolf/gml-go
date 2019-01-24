package generate

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/build"
	"go/format"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

const (
	Use              = "generate [dir]"
	ShortDescription = "Generate code so that assets and objects can be referenced by constant IDs"
)

const (
	genFile    = "gmlgo_gen.go"
	objectPath = "github.com/silbinarywolf/gml-go/gml.Object"
	version    = "0.1.0"
)

type Arguments struct {
	Directory string
}

func Run(args Arguments) {
	if args.Directory == "" {
		args.Directory = "."
	}
	dir := args.Directory

	// Support ./...
	var dirs []string
	if filepath.Base(dir) == "..." {
		dir = filepath.Dir(dir)
		files, err := ioutil.ReadDir(dir)
		if err != nil {
			log.Fatal(err)
		}
		for _, info := range files {
			if !info.IsDir() {
				continue
			}
			dirs = append(dirs, dir+"/"+info.Name())
		}
	} else {
		dirs = []string{dir}
	}

	for _, dir := range dirs {
		// todo(Jake): 2018-12-03 - #33
		// Replace "game" with scanning each sub-package, throw an error if multiple packages
		// have multiple objects. Constraint for now will be all object types need to be in the same package
		gameDir := filepath.Join(dir, "game")

		// get filename
		baseName := fmt.Sprintf(genFile)
		outputName := filepath.Join(gameDir, strings.ToLower(baseName))

		// check existing file
		var input []byte
		if _, err := os.Stat(outputName); !os.IsNotExist(err) {
			input, err = ioutil.ReadFile(outputName)
			if err != nil {
				log.Fatalf("reading file: %s", err)
			}
			if len(input) == 0 {
				log.Fatalf("cannot generate %s as it's empty. rename or delete your %s file.\n", outputName, outputName)
			}
			if !strings.Contains(string(input), "// Code generated by \"gmlgo") {
				log.Fatalf("cannot generate %s file as it's not using gmlgo generated code. rename your %s file.\n", outputName, outputName)
			}
		}

		// Run generate
		g := Generator{}
		g.parsePackageDir(gameDir, []string{})
		g.generate()
		g.generateAssets(dir)

		// If no generated output, don't write anything
		if g.buf.Len() == 0 {
			log.Fatalf("no gml.Object structs found, no output for %s\n", outputName)
		}

		// Format the output.
		src := g.format()

		// If no generated output, don't write anything
		if len(src) == 0 {
			log.Fatalf("no gml.Object structs found, no output for %s\n", outputName)
		}

		// Check if any changes
		if bytes.Equal(input, src) {
			//log.Printf("no changes to %s\n", outputName)
			continue
		}

		// Write to file.
		err := ioutil.WriteFile(outputName, src, 0644)
		if err != nil {
			log.Fatalf("error writing output: %s\n", err)
		}
		//log.Printf("updated %s\n", outputName)
	}
}

// isDirectory reports whether the named file is a directory.
func isDirectory(name string) bool {
	info, err := os.Stat(name)
	if err != nil {
		log.Fatal(err)
	}
	return info.IsDir()
}

// Generator holds the state of the analysis. Primarily used to buffer
// the output for format.Source.
type Generator struct {
	buf bytes.Buffer // Accumulated output.
	pkg *Package     // Package we are scanning.
}

func (g *Generator) Printf(format string, args ...interface{}) {
	fmt.Fprintf(&g.buf, format, args...)
}

// File holds a single parsed file and associated data.
type File struct {
	pkg  *Package  // Package to which this file belongs.
	file *ast.File // Parsed AST.
}

type Package struct {
	dir      string
	name     string
	defs     map[*ast.Ident]types.Object
	files    []*File
	typesPkg *types.Package
}

func buildContext(tags []string) *build.Context {
	ctx := build.Default
	ctx.BuildTags = tags
	return &ctx
}

// parsePackageDir parses the package residing in the directory.
func (g *Generator) parsePackageDir(directory string, tags []string) {
	pkg, err := buildContext(tags).ImportDir(directory, 0)
	if err != nil {
		log.Fatalf("parsePackageDir: cannot parse %s: %s", directory, err)
	}
	var names []string
	names = append(names, pkg.GoFiles...)
	names = prefixDirectory(directory, names)
	g.parsePackage(directory, names, nil)
}

// prefixDirectory places the directory name on the beginning of each name in the list.
func prefixDirectory(directory string, names []string) []string {
	if directory == "." {
		return names
	}
	ret := make([]string, len(names))
	for i, name := range names {
		ret[i] = filepath.Join(directory, name)
	}
	return ret
}

// parsePackage analyzes the single package constructed from the named files.
// If text is non-nil, it is a string to be used instead of the content of the file,
// to be used for testing. parsePackage exits if there is an error.
func (g *Generator) parsePackage(directory string, names []string, text interface{}) {
	var files []*File
	var astFiles []*ast.File
	g.pkg = new(Package)
	fs := token.NewFileSet()
	for _, name := range names {
		if !strings.HasSuffix(name, ".go") ||
			filepath.Base(name) == genFile {
			continue
		}
		parsedFile, err := parser.ParseFile(fs, name, text, parser.ParseComments)
		if err != nil {
			log.Fatalf("parsing package: %s: %s\n", name, err)
		}
		astFiles = append(astFiles, parsedFile)
		files = append(files, &File{
			file: parsedFile,
			pkg:  g.pkg,
		})
	}
	if len(astFiles) == 0 {
		log.Fatalf("%s: no buildable Go files\n", directory)
	}
	g.pkg.name = astFiles[0].Name.Name
	g.pkg.files = files
	g.pkg.dir = directory
	g.pkg.typeCheck(fs, astFiles)
}

// check type-checks the package so we can evaluate contants whose values we are printing.
func (pkg *Package) typeCheck(fs *token.FileSet, astFiles []*ast.File) {
	pkg.defs = make(map[*ast.Ident]types.Object)
	config := types.Config{
		IgnoreFuncBodies:         true,               // We only need to evaluate constants.
		Importer:                 importer.Default(), // func defaultImporter() types.Importer
		FakeImportC:              true,
		DisableUnusedImportCheck: true,
	}
	info := &types.Info{
		Defs: pkg.defs,
	}
	typesPkg, err := config.Check(pkg.dir, fs, astFiles, info)
	if err != nil {
		log.Fatalf("checking package: %s", err)
	}
	pkg.typesPkg = typesPkg
}

type Struct struct {
	Name string
}

type AssetKind struct {
	Name   string
	Assets []string
}

// hasEmbeddedObjectRecursive checks to see if "gml.Object" has been embedded
// into this struct, it will search each embedded struct to see if that struct
// also contains the "gml.Object" struct. This is to allow people to create
// base struct objects wherein all other objects can inherit that object.
func hasEmbeddedObjectRecursive(structTypeInfo *types.Struct) bool {
	for i := 0; i < structTypeInfo.NumFields(); i++ {
		field := structTypeInfo.Field(i)
		fieldTypeInfo, ok := field.Type().(*types.Named)
		if !ok {
			continue
		}
		structTypeInfo, ok := fieldTypeInfo.Underlying().(*types.Struct)
		if !ok {
			continue
		}
		// Search for embedded "gml.Object" field
		if field.Embedded() {
			if fieldTypeInfo.String() == objectPath {
				return true
			} else if hasEmbeddedObjectRecursive(structTypeInfo) {
				return true
			}
		}
	}
	return false
}

// generate produces the code for object indexes
func (g *Generator) generate() {
	var structsUsingGMLObject []Struct
	for _, file := range g.pkg.files {
		if file.file == nil {
			continue
		}
		//fmt.Printf("file: %s\n---------------\n\n", file.file.Name.String())
		ast.Inspect(file.file, func(n ast.Node) bool {
			switch n := n.(type) {
			// type XXXX struct
			case *ast.TypeSpec:
				structName := n.Name.Name
				typeInfo, ok := g.pkg.typesPkg.Scope().Lookup(structName).Type().(*types.Named)
				if !ok {
					// Skip if can't determine type
					return false
				}
				structTypeInfo := typeInfo.Underlying().(*types.Struct)
				if hasEmbeddedObjectRecursive(structTypeInfo) {
					structsUsingGMLObject = append(structsUsingGMLObject, Struct{
						Name: structName,
					})
				}
				return false
			}
			return true
		})
	}

	// Sort alphabetically
	sort.Slice(structsUsingGMLObject[:], func(i, j int) bool {
		return structsUsingGMLObject[i].Name < structsUsingGMLObject[j].Name
	})

	// Print the header and package clause.
	g.Printf("// Code generated by \"gmlgo\"\n")
	g.Printf("// %s\n", version)
	g.Printf("// DO NOT EDIT. DO NOT COMMIT TO YOUR VCS REPOSITORY.\n")
	g.Printf("\n")
	g.Printf(`package ` + g.pkg.name + `
`)
	g.Printf(`
import (
	"github.com/silbinarywolf/gml-go/gml"
)
`)
	g.generateObjectIndexes(structsUsingGMLObject)
	g.generateObjectMetaAndMethods(structsUsingGMLObject)
}

func (g *Generator) generateAssets(dir string) {
	// Read asset names
	assetDir := filepath.Join(dir, "asset")
	files, err := ioutil.ReadDir(assetDir)
	if err != nil {
		log.Fatal(err)
	}
	var assetKinds []AssetKind
	for _, f := range files {
		switch name := f.Name(); name {
		case "font",
			"sprite":
			files, err := ioutil.ReadDir(filepath.Join(assetDir, name))
			if err != nil {
				log.Fatal(err)
			}
			var assetNames []string
			for _, f := range files {
				if f.IsDir() {
					assetName := f.Name()
					if assetName == "data" &&
						name == "font" {
						// Ignore special "data" folder that's used
						// by "font"
						continue
					}
					assetNames = append(assetNames, f.Name())
				}
			}
			if len(assetNames) > 0 {
				assetKinds = append(assetKinds, AssetKind{
					Name:   name,
					Assets: assetNames,
				})
			}
		//case "layer":
		/*files, err := ioutil.ReadDir(filepath.Join(assetDir, name))
		if err != nil {
			log.Fatal(err)
		}
		var assetNames []string
		for _, f := range files {
			if !f.IsDir() {
				assetName := f.Name()
				assetNames = append(assetNames, f.Name())
			}
		}
		*/
		default:
			if !f.IsDir() {
				// Ignore files
				continue
			}
			log.Fatal(fmt.Errorf("Unexpected asset kind directory: %s", name))
		}
	}
	// Generate asset indexes
	for _, assetKind := range assetKinds {
		if len(assetKind.Assets) == 0 {
			continue
		}
		var prefix, gotype string
		switch assetKind.Name {
		case "font":
			prefix = "Fnt"
			gotype = "gml.FontIndex"
		case "sprite":
			prefix = "Spr"
			gotype = "gml.SpriteIndex"
		/*case "layer":
		prefix = "Lay"
		gotype = "gml.LayerIndex"*/
		default:
			panic("Unimplemented asset kind: " + assetKind.Name)
		}

		{
			g.Printf("const (\n")
			for i, assetName := range assetKind.Assets {
				// ie. SprPlayer    gml.SpriteIndex = 1
				g.Printf("	%s%s %s = %d\n", prefix, assetName, gotype, i+1)
			}
			g.Printf("\n)\n\n")
		}
		{
			g.Printf("var _gen_%s_index_to_name = []string{\n", prefix)
			for _, assetName := range assetKind.Assets {
				// ie. SprPlayer: "Player"
				g.Printf("	%s%s: \"%s\",\n", prefix, assetName, assetName)
			}
			g.Printf("\n}\n\n")
		}
		{
			g.Printf("var _gen_%s_name_to_index = map[string]%s{\n", prefix, gotype)
			for _, assetName := range assetKind.Assets {
				// ie. "Player": SprPlayer
				g.Printf("	\"%s\": %s%s,\n", assetName, prefix, assetName)
			}
			g.Printf("\n}\n")
		}
		switch assetKind.Name {
		case "font":
			g.Printf(`
func init() {
	gml.InitFontGeneratedData(_gen_Fnt_index_to_name, _gen_Fnt_name_to_index)
}

`)
		case "sprite":
			g.Printf(`
func init() {
	gml.InitSpriteGeneratedData(_gen_Spr_index_to_name, _gen_Spr_name_to_index)
}

`)
		default:
			panic("Unimplemented asset kind: " + assetKind.Name)
		}
	}
}

func (g *Generator) generateObjectIndexes(structsUsingGMLObject []Struct) {
	g.Printf(`
const (
`)
	for i, record := range structsUsingGMLObject {
		g.Printf("	Obj" + record.Name + " gml.ObjectIndex = " + strconv.Itoa(i+1) + "\n")
	}
	g.Printf(`)

`)
}

func (g *Generator) generateObjectMetaAndMethods(structsUsingGMLObject []Struct) {
	var prefix, gotype string
	prefix = "Obj"
	gotype = "gml.ObjectIndex"

	{
		g.Printf("var _gen_%s_index_to_name = []string{\n", prefix)
		for _, record := range structsUsingGMLObject {
			assetName := record.Name
			// ie. ObjPlayer: "Player"
			g.Printf("	%s%s: \"%s\",\n", prefix, assetName, assetName)
		}
		g.Printf("\n}\n\n")
	}
	{
		g.Printf("var _gen_%s_name_to_index = map[string]%s{\n", prefix, gotype)
		for _, record := range structsUsingGMLObject {
			assetName := record.Name
			// ie. "Player": ObjPlayer
			g.Printf("	\"%s\": %s%s,\n", assetName, prefix, assetName)
		}
		g.Printf("\n}\n\n")
	}
	{
		g.Printf("var _gen_%s_index_to_data = []gml.ObjectType{\n", prefix)
		for _, record := range structsUsingGMLObject {
			assetName := record.Name
			// ie. ObjPlayer: new(Player)
			g.Printf("	%s%s: new(%s),\n", prefix, assetName, assetName)
		}
		g.Printf("\n}\n")
	}

	{
		// Write Object types
		/*for _, record := range structsUsingGMLObject {
			//g.Printf("func (inst *" + record.Name + ") ObjectIndex() gml.ObjectIndex { return Obj" + record.Name + " }\n")
			g.Printf("func (inst *" + record.Name + ") ObjectName() string { return \"" + record.Name + "\" }\n")
			g.Printf("\n")
		}
		.ObjectIndex{*/
		g.Printf(`

func init() {
	gml.InitObjectGeneratedData(_gen_Obj_index_to_name, _gen_Obj_name_to_index, _gen_Obj_index_to_data)
}
`)
	}
}

// format returns the gofmt-ed contents of the Generator's buffer.
func (g *Generator) format() []byte {
	src, err := format.Source(g.buf.Bytes())
	if err != nil {
		// Should never happen, but can arise when developing this code.
		// The user can compile the output to see the error.
		log.Printf("warning: internal error: invalid Go generated: %s", err)
		log.Printf("warning: compile the package to analyze the error")
		return g.buf.Bytes()
	}
	return src
}
