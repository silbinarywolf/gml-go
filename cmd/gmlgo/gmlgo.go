package main

import (
	"fmt"
	"log"
	"os"

	"github.com/silbinarywolf/gml-go/cmd/gmlgo/cmd/fix"
	"github.com/silbinarywolf/gml-go/cmd/gmlgo/cmd/generate"
	"github.com/silbinarywolf/gml-go/cmd/gmlgo/cmd/serve"
	"github.com/spf13/cobra"
)

var (
	Directory string
	Tags      string
)

var rootCmd = &cobra.Command{
	Use:   "gmlgo",
	Short: "A tool for building gmlgo projects",
	Long:  ``,
	Run:   Run,
}

var generateCmd = &cobra.Command{
	Use:   "generate [dir]",
	Short: "Generate code so that assets and objects can be referenced by constant IDs",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		dir := ""
		if len(args) > 0 {
			dir = args[0]
		}
		generate.Run(generate.Arguments{
			Directory: dir,
		})
	},
}

var serveCmd = &cobra.Command{
	Use:   "serve [dir]",
	Short: "Serve a build of your game for playing in a web browser, defaults to port 8080",
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

var fixCmd = &cobra.Command{
	Use:   "fix [dir]",
	Short: fix.ShortDescription,
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		dir := ""
		if len(args) > 0 {
			dir = args[0]
		}
		fix.Run(fix.Arguments{
			Directory: dir,
		})
	},
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("gmlgo: ")

	rootCmd.AddCommand(generateCmd)
	rootCmd.AddCommand(serveCmd)
	rootCmd.AddCommand(fixCmd)
	serveCmd.Flags().StringVar(&Tags, "tags", "", "a list of build tags to consider satisfied during the build")
	//rootCmd.PersistentFlags().StringVarP(&Directory, "dir", "d", ".", "directory")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func Run(cmd *cobra.Command, args []string) {
	panic(args)
}
