package build

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/silbinarywolf/gml-go/cmd/gmlgo/cmd/generate"
	"github.com/silbinarywolf/gml-go/cmd/gmlgo/internal/base"
)

var Cmd = &base.Command{
	UsageLine: "gmlgo build [dir]",
	Short:     `Run "gmlgo generate" and "go build"`,
	Flag:      flag.NewFlagSet("build", flag.ExitOnError),
	Long:      ``,
	Run:       run,
}

var tags *string

var verbose *bool

func init() {
	tags = Cmd.Flag.String("tags", "", "a list of build tags to consider satisfied during the build")
	verbose = Cmd.Flag.Bool("verbose", false, "verbose")
}

func run(cmd *base.Command, args []string) {
	dir := ""
	if len(args) > 0 {
		dir = args[0]
	}

	// Run "go generate"
	generate.Run(generate.Arguments{
		Directory: dir,
		Verbose:   *verbose,
	})

	// Run "go build"
	{
		//panic(fmt.Sprintf("%v", os.Args[2:]))
		args := make([]string, 0, len(os.Args[2:])+1)
		args = append(args, "build")
		args = append(args, os.Args[2:]...)
		cmd := exec.Command("go", args...)
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
			panic(errOutput)
		}
		fmt.Printf("%s", stdOutput)
	}
}
