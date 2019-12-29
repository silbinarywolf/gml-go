package shared

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"golang.org/x/tools/go/packages"
)

const RootCmd = "gmlgo"

const gmlGoPackage = "github.com/silbinarywolf/gml-go"

var (
	cmdDir string
	cmdErr error
)

func computeCmdSourceDir(gameDir string) (string, error) {
	currentDir, err := filepath.Abs(gameDir)
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

// OpenBrowsers open web browser with a given url
// (can be http:// or file://)
func OpenBrowser(url string) {
	// Taken from:
	// https://presstige.io/p/Using-Go-instead-of-bash-for-scripts-6b51885c1f6940aeb40476000d0eb0fc
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		panic(err)
	}
}

func GetDefaultWasmJSPath(gameDir string) (string, error) {
	const baseName = "wasm_exec.js"
	// Look for user-override
	{
		dir := gameDir + "/html/" + baseName
		if _, err := os.Stat(dir); !os.IsNotExist(err) {
			return dir, nil
		}
	}
	// Look for engine default
	dir, err := computeCmdSourceDir(gameDir)
	if err != nil {
		return "", err
	}
	dir = dir + "/files/" + baseName
	return dir, nil
}

func GetDefaultIndexHTMLPath(gameDir string) (string, error) {
	const baseName = "index.html"
	// Look for user-override
	{
		dir := gameDir + "/html/" + baseName
		if _, err := os.Stat(dir); !os.IsNotExist(err) {
			return dir, nil
		}
	}
	// Look for engine default
	dir, err := computeCmdSourceDir(gameDir)
	if err != nil {
		return "", err
	}
	dir = dir + "/files/" + baseName
	return dir, nil
}

func ReadDefaultIndexHTML(gameDir string) ([]byte, error) {
	dir, err := GetDefaultIndexHTMLPath(gameDir)
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadFile(dir)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func ReadDefaultWasmJS(gameDir string) ([]byte, error) {
	dir, err := GetDefaultWasmJSPath(gameDir)
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadFile(dir)
	if err != nil {
		return nil, err
	}
	return data, nil
}
