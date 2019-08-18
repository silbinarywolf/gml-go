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

	// Get WASM files
	{
		var err error
		indexHTMLData, err = shared.ReadDefaultIndexHTML(dir)
		if err != nil {
			return err
		}
		wasmJSData, err = shared.ReadDefaultWasmJS(dir)
		if err != nil {
			return err
		}
	}

	// Generate unique folder name
	distFolder := dir + "/dist/" + time.Now().Format("2006-01-02_15-04-05")
	if err := os.MkdirAll(distFolder, os.ModePerm); err != nil {
		return err
	}

	// Run "go generate"
	generate.Run(generate.Arguments{
		Directory: dir,
		Verbose:   *verbose || *verboseShort,
	})

	// Build
	if err := compile(dir, distFolder, args); err != nil {
		return err
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
	//panic(fmt.Sprintf("%v", args))
	if err := build.Build(outputDir, argsNew, []string{"GOOS=" + GOOS, "GOARCH=" + GOARCH}); err != nil {
		return err
	}
	if err := asset.CopyAssetDirectory(gameDir+"/asset", outputDir+"/asset"); err != nil {
		return err
	}
	return nil
}
