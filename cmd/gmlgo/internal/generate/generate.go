package generate

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/build"
	"go/format"
	"go/token"
	"go/types"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"unicode"

	"github.com/silbinarywolf/gml-go/cmd/gmlgo/internal/base"
	"golang.org/x/tools/go/packages"
)

var Cmd = &base.Command{
	UsageLine: "generate [dir]",
	Short:     `generate code required by the game framework`,
	Long:      `generate code so that assets and objects can be referenced by constant IDs`,
	Flag:      flag.NewFlagSet("generate", flag.ExitOnError),
	Run:       run,
}

var tags *string

var verbose bool

func init() {
	tags = Cmd.Flag.String("tags", "", "a list of build tags to consider satisfied during the build")
	Cmd.Flag.BoolVar(&verbose, "v", false, "verbose")
	Cmd.Flag.BoolVar(&verbose, "verbose", false, "verbose")
}

const (
	genFile    = "gmlgo_gen.go"
	objectPath = "github.com/silbinarywolf/gml-go/gml.Object"
	version    = "0.1.0"
)

type Arguments struct {
	Directory string
	Verbose   bool
}

// Generator holds the state of the analysis. Primarily used to buffer
// the output for format.Source.
type Generator struct {
	buf              bytes.Buffer // Accumulated output.
	header           bytes.Buffer // to prepend before buf
	hasSerialization bool
	hasAudio         bool
}

func (g *Generator) Printf(format string, args ...interface{}) {
	fmt.Fprintf(&g.buf, format, args...)
}

func (g *Generator) Headerf(format string, args ...interface{}) {
	fmt.Fprintf(&g.header, format, args...)
}

type Parser struct {
	pkgs []*Package // Package we are scanning.
}

// File holds a single parsed file and associated data.
type File struct {
	pkg  *Package  // Package to which this file belongs.
	file *ast.File // Parsed AST.
}

type Package struct {
	path     string
	name     string
	defs     map[*ast.Ident]types.Object
	files    []*File
	fileSet  *token.FileSet
	typesPkg *types.Package
}

func run(cmd *base.Command, args []string) error {
	cmd.Flag.Parse(args)
	if !cmd.Flag.Parsed() {
		cmd.Flag.PrintDefaults()
		os.Exit(1)
	}
	args = cmd.Flag.Args()
	dir := ""
	if len(args) > 0 {
		dir = args[0]
	}
	return Run(Arguments{
		Directory: dir,
		Verbose:   verbose,
	})
}

func Run(args Arguments) (err error) {
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
		// Generate assets
		generateAssets(dir)

		// Parse packages
		packageList := make([]*Package, 0, 10)
		{
			projectDir, err := filepath.Abs(dir)
			if err != nil {
				panic(err)
			}
			projectDir = filepath.Join(projectDir, "...")
			cfg := &packages.Config{
				Mode: packages.LoadFiles | packages.NeedImports | packages.LoadSyntax,
				// NOTE(Jae): 2019-05-17
				// Tests shouldn't have game objects in them, so set to false
				Tests: false,
				// NOTE(Jae): 2019-05-17
				// Build tags probably don't matter for game object structure parsing.
				//BuildFlags: []string{fmt.Sprintf("-tags=%s", strings.Join(tags, " "))},
			}
			pkgs, err := packages.Load(cfg, projectDir)
			if err != nil {
				log.Fatal(err)
			}
			for _, pkg := range pkgs {
				//if len(pkg.Syntax) != len(pkg.GoFiles) {
				//	panic("Unexpected error")
				//}
				if len(pkg.GoFiles) == 0 {
					continue
				}
				pkgInfo := &Package{
					path:     filepath.Dir(pkg.GoFiles[0]),
					name:     pkg.Name,
					defs:     pkg.TypesInfo.Defs,
					typesPkg: pkg.Types,
					fileSet:  pkg.Fset,
					files:    make([]*File, len(pkg.Syntax)),
				}
				for i, file := range pkg.Syntax {
					pkgInfo.files[i] = &File{
						file: file,
						pkg:  pkgInfo,
						//trimPrefix:  p.trimPrefix,
						//lineComment: p.lineComment,
					}
				}
				packageList = append(packageList, pkgInfo)
			}
		}

		// Generation objects for each package (if any parsed)
		for _, pkg := range packageList {

			structsUsingObject := inspectGameObjectStructs(pkg)
			if len(structsUsingObject) == 0 {
				continue
			}
			fmt.Printf("Path: %s\n", pkg.path)
			// Run generate
			generateGameObject(pkg.path, pkg.name, structsUsingObject)
		}

		/*{
			gameDir := filepath.Join(dir, "game")

			// Run parser
			p := new(Parser)
			p.parseGamePackageDir(gameDir, []string{})
			structsUsingObject := p.parseGameObjectStructs()
			if len(structsUsingObject) > 0 {
				// Run generate
				generateGameObject(gameDir, p.pkg.name, structsUsingObject)
			}
		}
		{
			gameDir := filepath.Join(dir, "game", "object")

			// Run parser
			p := new(Parser)
			p.parseGamePackageDir(gameDir, []string{})
			structsUsingObject := p.parseGameObjectStructs()
			if len(structsUsingObject) > 0 {
				// Run generate
				generateGameObject(gameDir, p.pkg.name, structsUsingObject)
			}
		}*/
	}
	return
}

// isDirectory reports whether the named file is a directory.
func isDirectory(name string) bool {
	info, err := os.Stat(name)
	if err != nil {
		log.Fatal(err)
	}
	return info.IsDir()
}

func buildContext(tags []string) *build.Context {
	ctx := build.Default
	ctx.BuildTags = tags
	return &ctx
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

type Struct struct {
	Name    string
	Struct  *types.Struct
	ID      string
	Node    *ast.TypeSpec
	FileSet *token.FileSet
}

type AssetKind struct {
	Name   string
	Assets []Asset
}

type Asset struct {
	Name string
	Path string
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

func inspectGameObjectStructs(pkg *Package) []Struct {
	var structsUsingGMLObject []Struct
	for _, file := range pkg.files {
		if file.file == nil {
			continue
		}
		ast.Inspect(file.file, func(n ast.Node) bool {
			switch n := n.(type) {
			// type XXXX struct
			case *ast.TypeSpec:
				structName := n.Name.Name
				if pkg.typesPkg == nil {
					return false
				}
				typeName := pkg.typesPkg.Scope().Lookup(structName)
				if typeName == nil {
					// Skip if cannot lookup struct name
					return false
				}
				typeInfo, ok := typeName.Type().(*types.Named)
				if !ok {
					// Skip if can't determine type
					return false
				}
				structTypeInfo, ok := typeInfo.Underlying().(*types.Struct)
				if !ok {
					return false
				}
				if hasEmbeddedObjectRecursive(structTypeInfo) {
					structsUsingGMLObject = append(structsUsingGMLObject, Struct{
						Name:    structName,
						Struct:  structTypeInfo,
						ID:      typeInfo.Obj().Pkg().Path() + "." + typeInfo.Obj().Name(),
						Node:    n,
						FileSet: pkg.fileSet,
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
	return structsUsingGMLObject
}

// generate produces the code for object indexes
func generateGameObject(dir string, packageName string, structsUsingGMLObject []Struct) {
	// get filename
	outputName := filepath.Join(dir, genFile)

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
		if !strings.Contains(string(input), "// Code generated by \"gmlgo generate\". DO NOT EDIT.") {
			log.Fatalf("cannot generate %s file as it's not using gmlgo generated code. rename or delete your %s file.\n", outputName, outputName)
		}
	}

	// check for duplicates and provide a nice error message
	{
		var duplicatesFound []*Struct
		dup := make(map[string]*Struct, len(structsUsingGMLObject))
		for i, _ := range structsUsingGMLObject {
			record := &structsUsingGMLObject[i]
			if firstOccuranceRecord, ok := dup[record.ID]; ok {
				hasAddedFirst := false
				for _, other := range duplicatesFound {
					hasAddedFirst = hasAddedFirst || firstOccuranceRecord == other
				}
				if !hasAddedFirst {
					duplicatesFound = append(duplicatesFound, firstOccuranceRecord)
				}
				duplicatesFound = append(duplicatesFound, record)
				continue
			}
			dup[record.ID] = record
		}
		if len(duplicatesFound) > 0 {
			dir, err := os.Getwd()
			if err != nil {
				log.Fatal(err)
			}
			for _, record := range duplicatesFound {
				pos := record.FileSet.Position(record.Node.Pos())
				errorLoc := pos.String()
				relPath, err := filepath.Rel(dir, errorLoc)
				if err != nil {
					// if error getting relative path, ignore and show full path
					relPath = errorLoc
				}
				log.Printf("%s: duplicate object name %s\n", relPath, record.Name)
			}
			os.Exit(1)
		}
	}

	g := Generator{}
	g.generateObjectIndexes(structsUsingGMLObject)
	g.generateObjectMetaAndMethods(structsUsingGMLObject)
	g.generateSerialization(structsUsingGMLObject)

	// Header
	{
		// Print the header and package clause.
		g.generateCodeGenHeader()
		g.Headerf(`package ` + packageName + `
`)

		// Import
		{
			g.Headerf(`
import (`)
			if g.hasSerialization {
				g.Headerf(`
	"bytes"
	"encoding/binary"
`)
			}
			g.Headerf(`
	"github.com/silbinarywolf/gml-go/gml"
`)

			g.Headerf(")\n")
		}
	}

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

	// Don't write to file if there are no changes
	if bytes.Equal(input, src) {
		return
	}

	// Write to file.
	if err := ioutil.WriteFile(outputName, src, 0644); err != nil {
		log.Fatalf("error writing output: %s\n", err)
	}
	if verbose {
		log.Printf("%s\n", outputName)
	}
}

func (g *Generator) generateSerialization(structsUsingGMLObject []Struct) {
	for _, record := range structsUsingGMLObject {
		pkg := record.Struct.Field(0).Pkg()
		isSerializable := false
		for i := 0; i < record.Struct.NumFields(); i++ {
			field := record.Struct.Field(i)
			if field.Id() == "ObjectSerialize" {
				switch fieldType := field.Type().(type) {
				case *types.Named:
					if fieldType.Obj().Pkg().Path() == "github.com/silbinarywolf/gml-go/gml" {
						isSerializable = true
					}
				}
			}
		}
		if isSerializable &&
			record.Struct.NumFields() > 0 {
			g.Printf("func (self %s) UnsafeSnapshotMarshalBinaryRoot(buf *bytes.Buffer) error {", record.Name)
			for i := 0; i < record.Struct.NumFields(); i++ {
				field := record.Struct.Field(i)
				g.generateType(pkg, field, "", true)
			}
			g.Printf("\n	return nil\n}\n")

			g.Printf("func (self *%s) UnsafeSnapshotUnmarshalBinaryRoot(buf *bytes.Buffer) error {", record.Name)
			for i := 0; i < record.Struct.NumFields(); i++ {
				field := record.Struct.Field(i)
				g.generateType(pkg, field, "", false)
			}
			g.Printf("\n	return nil\n}\n")
			g.hasSerialization = true
		}
	}
}

func (g *Generator) generateType(pkg *types.Package, field *types.Var, prefix string, isWrite bool) {
	isExportedOrSamePackage := field.Exported() || pkg.Path() == field.Pkg().Path()
	if !isExportedOrSamePackage {
		panic("self." + prefix + field.Name() + " not exported. Cannot generate serialization code if using struct with unexported fields.")
	}
	switch fieldType := field.Type().(type) {
	case *types.Basic:
		switch fieldType.Kind() {
		case types.Int:
			if isWrite {
				g.Printf(`
if err := binary.Write(buf, binary.LittleEndian, int64(self.%s%s)); err != nil {
	return err
}`, prefix, field.Name())
			} else {
				g.Printf(`
{
var d int64
if err := binary.Read(buf, binary.LittleEndian, &d); err != nil {
	return err
}
self.%s%s = int(d)
}`, prefix, field.Name())
			}
		default:
			if isWrite {
				g.Printf(`
if err := binary.Write(buf, binary.LittleEndian, self.%s%s); err != nil {
	return err
}`, prefix, field.Name())
			} else {
				g.Printf(`
if err := binary.Read(buf, binary.LittleEndian, &self.%s%s); err != nil {
	return err
}`, prefix, field.Name())
			}
		}
	case *types.Named:
		hasMarshalMethod := false
		for i := 0; i < fieldType.NumMethods(); i++ {
			method := fieldType.Method(i)
			if isWrite && method.Name() == "UnsafeSnapshotMarshalBinary" {
				typeInfo, ok := method.Type().(*types.Signature)
				if !ok {
					panic("Expected method to have type signature")
				}

				// Validate params
				{
					values := typeInfo.Params()
					if values.Len() != 1 {
						panic("Expected UnsafeSnapshotMarshalBinary to only have 1 parameter")
					}
					param, ok := values.At(0).Type().(*types.Pointer)
					if !ok {
						panic("Expected parameter 1 to be pointer " + param.String())
					}
					underlyingType, ok := param.Elem().(*types.Named)
					if !ok {
						panic("Expected parameter 1 to be named")
					}
					isBytesBuf := underlyingType.Obj().Pkg().Path() == "bytes" &&
						underlyingType.Obj().Name() == "Buffer"
					if !isBytesBuf {
						panic("Expected parameter 1 to be bytes.Buffer")
					}
				}

				// Validate return
				{
					values := typeInfo.Results()
					if values.Len() != 1 {
						panic("Expected UnsafeSnapshotMarshalBinary to only have 1 return value")
					}

					value := values.At(0).Type().(*types.Named)
					if value.Obj().Id() != "_.error" {
						// NOTE: Jake: 2019-03-27
						// Haven't checked if "_.error" is the correct way to confirm
						// the error type is correct. Guessing!
						panic("Expected return type error, not " + value.Obj().Name())
					}
				}

				hasMarshalMethod = true
			}
			if !isWrite && method.Name() == "UnsafeSnapshotUnmarshalBinary" {
				typeInfo, ok := method.Type().(*types.Signature)
				if !ok {
					panic("Expected method to have type signature")
				}

				// Validate params
				{
					values := typeInfo.Params()
					if values.Len() != 1 {
						panic("Expected UnsafeSnapshotUnmarshalBinary to only have 1 parameter")
					}
					param, ok := values.At(0).Type().(*types.Pointer)
					if !ok {
						panic("Expected parameter 1 to be pointer " + param.String())
					}
					underlyingType, ok := param.Elem().(*types.Named)
					if !ok {
						panic("Expected parameter 1 to be named")
					}
					isBytesBuf := underlyingType.Obj().Pkg().Path() == "bytes" &&
						underlyingType.Obj().Name() == "Buffer"
					if !isBytesBuf {
						panic("Expected parameter 1 to be bytes.Buffer")
					}
				}

				// Validate return
				{
					values := typeInfo.Results()
					if values.Len() != 1 {
						panic("Expected UnsafeSnapshotMarshalBinary to only have 1 return value")
					}

					value := values.At(0).Type().(*types.Named)
					if value.Obj().Id() != "_.error" {
						// NOTE: Jake: 2019-03-27
						// Haven't checked if "_.error" is the correct way to confirm
						// the error type is correct. Guessing!
						panic("Expected return type error, not " + value.Obj().Name())
					}
				}

				hasMarshalMethod = true
			}
		}

		if hasMarshalMethod {
			if isWrite {
				g.Printf(`
if err := self.%s.UnsafeSnapshotMarshalBinary(buf); err != nil {
	return err
}`, field.Name())
			} else {
				g.Printf(`
if err := self.%s.UnsafeSnapshotUnmarshalBinary(buf); err != nil {
	return err
}`, field.Name())
			}
		} else {
			switch fieldType := fieldType.Underlying().(type) {
			case *types.Struct:
				prefix := prefix + field.Name() + "."
				for i := 0; i < fieldType.NumFields(); i++ {
					field := fieldType.Field(i)
					g.generateType(pkg, field, prefix, isWrite)
				}
			case *types.Basic:
				if isWrite {
					g.Printf(`
if err := binary.Write(buf, binary.LittleEndian, self.%s%s); err != nil {
	return err
}`, prefix, field.Name())
				} else {
					g.Printf(`
if err := binary.Read(buf, binary.LittleEndian, &self.%s%s); err != nil {
	return err
}`, prefix, field.Name())
				}
			default:
				panic(fmt.Sprintf("Unhandled field type: %T\n", fieldType))
			}
		}
	default:
		fmt.Printf("default: %s, %T\n", fieldType.String(), fieldType)
	}
}

func (g *Generator) generateCodeGenHeader() {
	// NOTE(Jake): 2019-09-14
	// Changed the first line to friendly to lint tools using the regex:
	// "^\/\/ Code generated .* DO NOT EDIT\.$" <- tested on https://regexr.com/
	// https://github.com/golang/go/issues/13560#issuecomment-288457920
	g.Headerf("// Code generated by \"gmlgo generate\". DO NOT EDIT.\n")
	g.Headerf("// DO NOT COMMIT TO YOUR VCS REPOSITORY.\n")
	g.Headerf("\n")
}

func getFilesRecursively(assetDir string, assetTypeDir string, assetNamesUsed map[string]string) []Asset {
	rootDir := filepath.Clean(assetDir + "/" + assetTypeDir)
	filepathSet := make([]Asset, 0, 50)
	dirs := make([]string, 0, 50)
	dirs = append(dirs, rootDir)
	for len(dirs) > 0 {
		dir := dirs[len(dirs)-1]
		dirs = dirs[:len(dirs)-1]
		files, err := ioutil.ReadDir(dir)
		if err != nil {
			log.Fatal(err)
		}
		isAsset := false
		for _, f := range files {
			name := f.Name()
			path := dir + "/" + name
			if f.IsDir() {
				dirs = append(dirs, path)
				continue
			}
			// NOTE(Jake): 2019-04-27
			// Maybe change this "isAsset" check into a function callback
			// so each asset type has explicit rules (except for "custom", whose rule would be any non-folder file)
			isAsset = isAsset ||
				(len(name) >= 2 && name[0] == '0' && name[1] == '.') || // ie. "0.png"
				name == "config.json" ||
				name == "sound.mp3" || name == "sound.wav"
		}
		if isAsset {
			name := filepath.Base(dir)
			relativeAssetPath := dir[len(assetDir)+1:]

			// Check if asset name is valid Go
			isExported := false
			for _, c := range name {
				isExported = unicode.IsUpper(c)
				break
			}
			if !isExported {
				panic(fmt.Errorf("Asset names must begin with a capital letter: %s", name))
			}

			// Check if duplicate
			if otherPath, ok := assetNamesUsed[name]; ok {
				panic(fmt.Errorf("Cannot have duplicate asset names:\n- %s\n- %s", relativeAssetPath, otherPath))
			}

			filepathSet = append(filepathSet, Asset{
				Name: name,
				Path: relativeAssetPath[len(assetTypeDir)+1:],
			})
			assetNamesUsed[name] = relativeAssetPath
		}
	}
	sort.Slice(filepathSet, func(i, j int) bool {
		return filepathSet[i].Name < filepathSet[j].Name
	})
	return filepathSet
}

func generateAssets(dir string) {
	g := Generator{}

	// Generate header
	{
		// Print the header and package clause.
		g.generateCodeGenHeader()
		g.Headerf(`package asset
`)

		// Import
		{
			g.Headerf(`
import (`)
			if g.hasSerialization {
				g.Headerf(`
	"bytes"
	"encoding/binary"
`)
			}
			g.Headerf(`
	"github.com/silbinarywolf/gml-go/gml"
	"github.com/silbinarywolf/gml-go/gml/audio"
`)

			g.Headerf(")\n")
		}

		g.Headerf(`
// Silence errors if audio is unused
var _ = audio.InitSoundGeneratedData
`)
	}
	assetDir := filepath.Join(dir, "asset")
	if _, err := os.Stat(assetDir); os.IsNotExist(err) {
		// Skip if we have no asset folder
		return
	}
	files, err := ioutil.ReadDir(assetDir)
	if err != nil {
		log.Fatal(err)
	}
	var assetKinds []AssetKind
	assetNamesUsed := make(map[string]string, len(files))
	for _, f := range files {
		switch assetTypeFolderName := f.Name(); assetTypeFolderName {
		case "font",
			"sprite",
			"sound",
			"custom":
			filepathSet := getFilesRecursively(assetDir, assetTypeFolderName, assetNamesUsed)
			if len(filepathSet) == 0 {
				continue
			}
			assetKinds = append(assetKinds, AssetKind{
				Name:   assetTypeFolderName,
				Assets: filepathSet,
			})
		default:
			if !f.IsDir() {
				// Ignore files
				continue
			}
			log.Fatal(fmt.Errorf("Unexpected asset directory type: %s, create and use a \"custom/%s\" folder for custom asset systems.", assetTypeFolderName, assetTypeFolderName))
		}
	}
	// Generate asset indexes
	for _, assetKind := range assetKinds {
		if len(assetKind.Assets) == 0 {
			continue
		}
		var kind, gotype string
		switch assetKind.Name {
		case "font":
			kind = "Fnt"
			gotype = "gml.FontIndex"
		case "sprite":
			kind = "Spr"
			gotype = "gml.SpriteIndex"
		case "sound":
			kind = "Snd"
			gotype = "audio.SoundIndex"
		case "custom":
			kind = "Cus"
			gotype = "gml.CustomAssetIndex"
			{
				g.Printf("const (\n")
				for i, asset := range assetKind.Assets {
					// ie. Player    gml.SpriteIndex = 1
					g.Printf("	%s %s = %d\n", asset.Name, gotype, i+1)
				}
				g.Printf("\n)\n\n")
			}
			{
				g.Printf("var _gen_Cus_index_to_path = []string{\n")
				for _, asset := range assetKind.Assets {
					// ie. Player: "objects/Player"
					g.Printf("	%s: \"%s\",\n", asset.Name, asset.Path)
				}
				g.Printf("\n}\n\n")
			}
			g.Printf(`
func init() {
	gml.InitCustomAsset(_gen_Cus_index_to_path)
}

`)
			continue
		default:
			panic("Unimplemented asset kind: " + assetKind.Name)
		}

		{
			g.Printf("const (\n")
			for i, asset := range assetKind.Assets {
				// ie. Player    gml.SpriteIndex = 1
				g.Printf("	%s %s = %d\n", asset.Name, gotype, i+1)
			}
			g.Printf("\n)\n\n")
		}

		// todo(Jake): 2019-04-27
		// Deprecate providing name, the filepath should
		// be all thats required / the unique key to the asset
		{
			g.Printf("var _gen_%s_index_to_name = []string{\n", kind)
			for _, asset := range assetKind.Assets {
				// ie. Player: "Player"
				g.Printf("	%s: \"%s\",\n", asset.Name, asset.Name)
			}
			g.Printf("\n}\n\n")
		}
		{
			g.Printf("var _gen_%s_index_to_path = []string{\n", kind)
			for _, asset := range assetKind.Assets {
				// ie. Player: "objects/Player"
				g.Printf("	%s: \"%s\",\n", asset.Name, asset.Path)
			}
			g.Printf("\n}\n\n")
		}
		{
			g.Printf("var _gen_%s_name_to_index = map[string]%s{\n", kind, gotype)
			for _, asset := range assetKind.Assets {
				// ie. "Player": SprPlayer
				g.Printf("	\"%s\": %s,\n", asset.Name, asset.Name)
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
	gml.InitSpriteGeneratedData(_gen_Spr_index_to_name, _gen_Spr_name_to_index, _gen_Spr_index_to_path)
}

`)
		case "sound":
			g.Printf(`
func init() {
	audio.InitSoundGeneratedData(_gen_Snd_index_to_name, _gen_Snd_name_to_index)
}

`)
		case "custom":
			// no-op
		default:
			panic("Unimplemented asset kind: " + assetKind.Name)
		}
	}

	// Load existing asset code-gen file
	var input []byte
	outputName := filepath.Join(assetDir, genFile)
	{
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
	}

	src := g.format()

	if bytes.Equal(input, src) {
		// Don't write to file if no changes
		return
	}

	// Write to file.
	if err := ioutil.WriteFile(outputName, src, 0644); err != nil {
		log.Fatalf("error writing output: %s\n", err)
	}
	if verbose {
		log.Printf("%s\n", outputName)
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
	var buf bytes.Buffer
	buf.Write(g.header.Bytes())
	buf.Write(g.buf.Bytes())
	src, err := format.Source(buf.Bytes())
	if err != nil {
		// Should never happen, but can arise when developing this code.
		// The user can compile the output to see the error.
		log.Fatalf("invalid Go generated: %s\n%s", err, buf.Bytes())
	}
	return src
}
