package publish

import (
	"flag"
	"io/ioutil"
	"os"
	"path/filepath"
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

func init() {
	tags = Cmd.Flag.String("tags", "", "a list of build tags to consider satisfied during the build")
	verboseShort = Cmd.Flag.Bool("v", false, "verbose")
	verbose = Cmd.Flag.Bool("verbose", false, "verbose")
}

func run(cmd *base.Command, args []string) (err error) {
	cmd.Flag.Parse(args)
	if !cmd.Flag.Parsed() {
		cmd.Flag.PrintDefaults()
		os.Exit(1)
	}
	dir := "."
	if dirArgs := cmd.Flag.Args(); len(dirArgs) > 0 {
		dir = dirArgs[0]
	}

	indexHTMLData, err := shared.ReadDefaultIndexHTML()
	if err != nil {
		return err
	}
	wasmJSData, err := shared.ReadDefaultWasmJS()
	if err != nil {
		return err
	}

	// Generate unique folder name
	distFolder, err := filepath.Abs("dist/" + time.Now().Format("2006-01-02_15-04-05"))
	if err != nil {
		return err
	}
	if err := os.MkdirAll(distFolder, os.ModePerm); err != nil {
		return err
	}

	// Run "go generate"
	generate.Run(generate.Arguments{
		Directory: dir,
		Verbose:   *verbose || *verboseShort,
	})

	// Build web
	{
		outputDir := distFolder + "/web"
		if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
			return err
		}
		args := args
		args = append(args, "-o", outputDir+"/main.wasm")
		build.Build(outputDir, args, []string{"GOOS=js", "GOARCH=wasm"})
		if err := ioutil.WriteFile(outputDir+"/index.html", indexHTMLData, os.ModePerm); err != nil {
			return err
		}
		if err := ioutil.WriteFile(outputDir+"/wasm_exec.js", wasmJSData, os.ModePerm); err != nil {
			return err
		}
		if err := asset.CopyAssetDirectory(dir+"/asset", outputDir+"/asset"); err != nil {
			return err
		}
	}

	return nil
}
