package shared

import (
	"errors"
	"io/ioutil"
	"path/filepath"

	"golang.org/x/tools/go/packages"
)

const RootCmd = "gmlgo"

const gmlGoPackage = "github.com/silbinarywolf/gml-go"

var (
	cmdDir string
	cmdErr error
)

func computeCmdSourceDir() (string, error) {
	currentDir, err := filepath.Abs(".")
	if err != nil {
		return "", err
	}
	cfg := &packages.Config{
		Dir: currentDir,
	}
	pkgs, err := packages.Load(cfg, "github.com/silbinarywolf/gml-go/cmd/gmlgo")
	if err != nil {
		return "", err
	}
	if len(pkgs) == 0 {
		return "", errors.New("Unable to find package: " + gmlGoPackage)
	}
	pkg := pkgs[0]
	if len(pkg.GoFiles) == 0 {
		return "", errors.New("Cannot find *.go files in:" + currentDir)
	}
	dir := filepath.Dir(pkg.GoFiles[0])
	return dir, nil
}

// getCmdSourceDir returns a cached copy of computeCmdSourceDir
func getCmdSourceDir() (string, error) {
	if cmdDir != "" || cmdErr != nil {
		return cmdDir, cmdErr
	}
	cmdDir, cmdErr = computeCmdSourceDir()
	return cmdDir, cmdErr
}

func ReadDefaultIndexHTML() ([]byte, error) {
	dir, err := getCmdSourceDir()
	if err != nil {
		return nil, err
	}
	dir = dir + "/files/index.html"
	data, err := ioutil.ReadFile(dir)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func ReadDefaultWasmJS() ([]byte, error) {
	dir, err := getCmdSourceDir()
	if err != nil {
		return nil, err
	}
	dir = dir + "/files/wasm_exec.js"
	data, err := ioutil.ReadFile(dir)
	if err != nil {
		return nil, err
	}
	return data, nil
}
