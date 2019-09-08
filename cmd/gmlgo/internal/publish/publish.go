package publish

import (
	"flag"
	"io/ioutil"
	"os"
	"time"

	"github.com/silbinarywolf/gml-go/cmd/gmlgo/internal/asset"
	"github.com/silbinarywolf/gml-go/cmd/gmlgo/internal/base"
	"github.com/silbinarywolf/gml-go/cmd/gmlgo/internal/build"
	"github.com/silbinarywolf/gml-go/cmd/gmlgo/internal/generate"
	"github.com/silbinarywolf/gml-go/cmd/gmlgo/internal/shared"
	"golang.org/x/xerrors"
)

var Cmd = &base.Command{
	UsageLine: "publish [dir]",
	Short:     `create distributables in "dist" folder`,
	Long:      `executes "gmlgo generate" and "go build" for creating distributables in "dist"`,
	Flag:      flag.NewFlagSet("publish", flag.ExitOnError),
	Run:       run,
}

var tags *string

var verboseShort *bool

var verbose *bool

var (
	indexHTMLData []byte
	wasmJSData    []byte
)

func init() {
	tags = Cmd.Flag.String("tags", "", "a list of build tags to consider satisfied during the build")
	verboseShort = Cmd.Flag.Bool("v", false, "verbose")
	verbose = Cmd.Flag.Bool("verbose", false, "verbose")
}

func run(cmd *base.Command, args []string) error {
	cmd.Flag.Parse(args)
	if !cmd.Flag.Parsed() {
		cmd.Flag.PrintDefaults()
		os.Exit(1)
	}
	dir := "."
	if dirArgs := cmd.Flag.Args(); len(dirArgs) > 0 {
		dir = dirArgs[0]
	}

	if _, err := os.Stat(dir); err != nil {
		if os.IsNotExist(err) {
			return xerrors.Errorf("Directory does not exist: %s", dir)
		}
		return xerrors.Errorf("Error opening directory: %w", err)
	}

	// Get WASM files
	{
		var err error
		indexHTMLData, err = shared.ReadDefaultIndexHTML(dir)
		if err != nil {
			return xerrors.Errorf("reading index.html failed: %w", err)
		}
		wasmJSData, err = shared.ReadDefaultWasmJS(dir)
		if err != nil {
			return xerrors.Errorf("reading wasm_exec.js failed: %w", err)
		}
	}

	// Generate unique folder name
	relativeFolder := "dist/" + time.Now().Format("2006-01-02_15-04-05")
	distFolder := dir + "/" + relativeFolder
	if err := os.MkdirAll(distFolder, os.ModePerm); err != nil {
		return xerrors.Errorf("creating \""+relativeFolder+"\" folder failed: %w", err)
	}

	// Run "go generate"
	generate.Run(generate.Arguments{
		Directory: dir,
		Verbose:   *verbose || *verboseShort,
	})

	// Build
	{
		// NOTE(Jake): 2019-08-16
		// I currently can't get publishing working with cross-platform
		// compiling as it currently relies on CGo code (ie. glfw)
		if err := compileWeb(dir, distFolder, args); err != nil {
			return xerrors.Errorf("error compiling web: %w", err)
		}
		if isLinux {
			if err := compileLinux(dir, distFolder, args); err != nil {
				return xerrors.Errorf("error compiling linux: %w", err)
			}
		}
		if isWindows {
			if err := compileWindows(dir, distFolder, args); err != nil {
				return xerrors.Errorf("error compiling windows: %w", err)
			}
		}
		if isMac {
			if err := compileMac(dir, distFolder, args); err != nil {
				return xerrors.Errorf("error compiling mac: %w", err)
			}
		}
	}

	return nil
}

func compileWeb(dir string, distFolder string, args []string) error {
	distFolder = distFolder + "/web"
	if err := compileBinary(
		dir,
		distFolder,
		"main.wasm",
		"js",
		"wasm",
		args,
	); err != nil {
		return err
	}
	if err := ioutil.WriteFile(distFolder+"/index.html", indexHTMLData, os.ModePerm); err != nil {
		return err
	}
	if err := ioutil.WriteFile(distFolder+"/wasm_exec.js", wasmJSData, os.ModePerm); err != nil {
		return err
	}
	return nil
}

func compileWindows(dir string, distFolder string, args []string) error {
	if err := compileBinary(
		dir,
		distFolder+"/windows",
		"game.exe",
		"windows",
		"amd64",
		args,
	); err != nil {
		return err
	}
	return nil
}

func compileLinux(dir string, distFolder string, args []string) error {
	if err := compileBinary(
		dir,
		distFolder+"/linux",
		"game",
		"linux",
		"amd64",
		args,
	); err != nil {
		return err
	}
	return nil
}

func compileMac(dir string, distFolder string, args []string) error {
	if err := compileBinary(
		dir,
		distFolder+"/mac",
		"game",
		"darwin",
		"amd64",
		args,
	); err != nil {
		return err
	}
	return nil
}

func compileBinary(gameDir string, outputDir string, binaryName string, GOOS string, GOARCH string, args []string) error {
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		return err
	}
	argsNew := make([]string, 0, 3)
	argsNew = append(argsNew, "-o", outputDir+"/"+binaryName)
	argsNew = append(argsNew, args...)
	//panic(xerrors.Sprintf("%v", args))
	if err := build.Build(outputDir, argsNew, []string{"GOOS=" + GOOS, "GOARCH=" + GOARCH}); err != nil {
		return err
	}
	if err := asset.CopyAssetDirectory(gameDir+"/asset", outputDir+"/asset"); err != nil {
		return xerrors.Errorf("error copying asset directory: %w", err)
	}
	return nil
}
