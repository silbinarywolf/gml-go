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

// buildDirShort is not unused, we use it to silence errors about passing an "-o" flag
// We pass it seamlessly to "go build" by doing nothing
var buildDirShort *string

func init() {
	tags = Cmd.Flag.String("tags", "", "a list of build tags to consider satisfied during the build")
	verboseShort = Cmd.Flag.Bool("v", false, "verbose")
	verbose = Cmd.Flag.Bool("verbose", false, "verbose")
	buildDirShort = Cmd.Flag.String("o", "", "build dir")
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
	if err := Build(dir, args, nil); err != nil {
		return err
	}

	return
}

func Build(dir string, args []string, envVars []string) error {
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
	cmd.Env = append(cmd.Env, envVars...)

	cmdOut, _ := cmd.StdoutPipe()
	cmdErr, _ := cmd.StderrPipe()

	err := cmd.Start()
	if err != nil {
		return err
	}
	errOutput, _ := ioutil.ReadAll(cmdErr)
	stdOutput, _ := ioutil.ReadAll(cmdOut)
	if len(errOutput) > 0 {
		return fmt.Errorf("%s", errOutput)
	}
	fmt.Printf("%s", stdOutput)
	return nil
}
