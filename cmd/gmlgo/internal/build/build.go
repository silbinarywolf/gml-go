package build

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/silbinarywolf/gml-go/cmd/gmlgo/internal/base"
	"github.com/silbinarywolf/gml-go/cmd/gmlgo/internal/generate"
)

var Cmd = &base.Command{
	UsageLine: "build [dir]",
	Short:     `compile game executable`,
	Long:      `run "gmlgo generate" and "go build"`,
	Flag:      flag.NewFlagSet("build", flag.ExitOnError),
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
	dir := ""
	if args := cmd.Flag.Args(); len(args) > 0 {
		dir = args[0]
	}

	// Run "go generate"
	generate.Run(generate.Arguments{
		Directory: dir,
		Verbose:   *verbose || *verboseShort,
	})

	// Run "go build"
	Build(dir, args)

	return
}

func Build(dir string, args []string) {
	var buildArgs []string
	if len(args) > 1 {
		buildArgs = make([]string, 0, len(args[2:])+1)
		buildArgs = append(buildArgs, "build")
		buildArgs = append(buildArgs, args...)
	} else {
		buildArgs = []string{"build"}
		if dir != "" {
			buildArgs = append(buildArgs, dir)
		}
	}
	cmd := exec.Command("go", buildArgs...)
	cmd.Env = os.Environ()

	cmdOut, _ := cmd.StdoutPipe()
	cmdErr, _ := cmd.StderrPipe()

	err := cmd.Start()
	if err != nil {
		panic(err)
	}
	errOutput, _ := ioutil.ReadAll(cmdErr)
	stdOutput, _ := ioutil.ReadAll(cmdOut)
	if len(errOutput) > 0 {
		fmt.Printf(string(errOutput))
		os.Exit(1)
	}
	fmt.Printf("%s", stdOutput)
}
