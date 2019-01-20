package fix

import (
	"errors"
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"strings"
)

const (
	ShortDescription = "Fix finds GML-Go programs that use old APIs and rewrites them to use newer ones."
)

var (
	errIgnoreVendor = errors.New(`Cannot process "vendor" folder.`)
)

type Arguments struct {
	Directory string // .
}

const (
	vendorBaseName = "vendor"
	gmlgoGenName   = "gmlgo_gen.go"
)

type dirWalker struct {
	goFiles             []string
	hasFoundMainPackage bool
}

// getValidFilesRecursive will recursively walk the directory structure using the following rules:
// - if a folder contains no *.go files, keep recursively searching
// - if a folder contains a *.go file with "package main", get every file within that project and in sub-directories
// We do this so that we only upgrade example games if this is run in the root of the repository
func (info *dirWalker) getValidFilesRecursive(dir string, isInAMainPackage bool) {
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
			if !strings.HasSuffix(path, ".go") ||
				base == gmlgoGenName {
				continue
			}
			goFiles = append(goFiles, path)
		}
	}

	// If no *.go files found, keep recursively searching
	if len(goFiles) == 0 {
		for _, dir := range dirs {
			info.getValidFilesRecursive(dir, false)
		}
		return
	}

	if len(goFiles) > 0 {
		if isInAMainPackage {
			// If we're underneath a "main" package, just get
			// every file
			info.goFiles = append(info.goFiles, goFiles...)
		} else {
			// Parse each *.go file until we know we're in the context
			// of "package main"
			fileSet := token.NewFileSet()
			for _, path := range goFiles {
				text, err := ioutil.ReadFile(path)
				if err != nil {
					log.Fatalf("cannot open *.go file: %s", err)
				}
				parsedFile, err := parser.ParseFile(fileSet, path, text, parser.PackageClauseOnly)
				if err != nil {
					log.Fatal(err)
				}
				if parsedFile.Name.Name == "main" {
					info.goFiles = append(info.goFiles, path)
					isInAMainPackage = true
				}
			}
		}
		// Only keep scanning if *.go files are found and
		// it's a "main" package
		if isInAMainPackage {
			for _, dir := range dirs {
				info.getValidFilesRecursive(dir, isInAMainPackage)
			}
		}
	}
}

func Run(args Arguments) {
	if args.Directory == "" {
		args.Directory = "."
	}
	dir := args.Directory
	//dir, err := filepath.Abs(dir)
	//if err != nil {
	//	log.Fatal(err)
	//}

	var dirInfo dirWalker
	dirInfo.getValidFilesRecursive(dir, false)
	log.Printf("%v", dirInfo.goFiles)
}
