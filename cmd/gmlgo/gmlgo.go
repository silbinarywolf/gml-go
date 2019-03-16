package main

import (
	"fmt"
	"log"
	"os"

	"github.com/silbinarywolf/gml-go/cmd/gmlgo/cmd/generate"
	"github.com/silbinarywolf/gml-go/cmd/gmlgo/cmd/serve"
	"github.com/silbinarywolf/gml-go/cmd/gmlgo/internal/base"
	"github.com/silbinarywolf/gml-go/cmd/gmlgo/internal/build"
	"github.com/silbinarywolf/gml-go/cmd/gmlgo/internal/shared"
	"github.com/spf13/cobra"
)

var cmds = []*base.Command{
	build.Cmd,
}

var (
	Tags    string
	Verbose bool
)

var rootCmd = &cobra.Command{
	Use: "gmlgo",
}

var generateCmd = &cobra.Command{
	Use:   generate.Use,
	Short: generate.ShortDescription,
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		dir := ""
		if len(args) > 0 {
			dir = args[0]
		}
		generate.Run(generate.Arguments{
			Directory: dir,
			Verbose:   Verbose,
		})
	},
}

var serveCmd = &cobra.Command{
	Use:   serve.Use,
	Short: serve.ShortDescription,
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		dir := ""
		if len(args) > 0 {
			dir = args[0]
		}
		serve.Run(serve.Arguments{
			Directory: dir,
			Tags:      Tags,
		})
	},
}

func main() {
	log.SetFlags(0)
	log.SetPrefix(shared.RootCmd + ": ")

	rootCmd.AddCommand(generateCmd)
	rootCmd.AddCommand(serveCmd)
	//rootCmd.AddCommand(buildCmd)
	//rootCmd.AddCommand(fixCmd)
	generateCmd.Flags().BoolVar(&Verbose, "v", false, "verbose")
	serveCmd.Flags().StringVar(&Tags, "tags", "", "a list of build tags to consider satisfied during the build")

	if len(os.Args) < 2 {
		fmt.Print(`
GmlGo is a tool for building games using the GmlGo library

Usage:

        gmlgo <command> [arguments]

The commands are:

        build		run generate, assetpack and compile packages and dependencies
        assetpack	[todo] make this build asset files
        generate	generate Go files by processing gml.Object's and assets
`)
		os.Exit(1)
	}
	args := os.Args[2:]
	for _, cmd := range cmds {
		if cmd.Name() == os.Args[1] {
			cmd.Flag.Parse(args)
			if !cmd.Flag.Parsed() {
				cmd.Flag.PrintDefaults()
				os.Exit(1)
			}
			args = cmd.Flag.Args()
			cmd.Run(cmd, args)
			break
		}
	}
}
