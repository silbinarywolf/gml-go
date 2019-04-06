package main

import (
	"fmt"
	"log"
	"os"

	"github.com/silbinarywolf/gml-go/cmd/gmlgo/internal/base"
	"github.com/silbinarywolf/gml-go/cmd/gmlgo/internal/build"
	"github.com/silbinarywolf/gml-go/cmd/gmlgo/internal/generate"
	"github.com/silbinarywolf/gml-go/cmd/gmlgo/internal/serve"
	"github.com/silbinarywolf/gml-go/cmd/gmlgo/internal/shared"
)

var cmds = []*base.Command{
	build.Cmd,
	generate.Cmd,
	serve.Cmd,
}

func main() {
	log.SetFlags(0)
	log.SetPrefix(shared.RootCmd + ": ")

	if len(os.Args) < 2 {
		fmt.Print(`
GmlGo is a tool for building games using the GmlGo library

Usage:

        gmlgo <command> [arguments]

The commands are:

        build		run generate, assetpack and compile packages and dependencies
        assetpack	[todo] make this build asset files
        generate	generate Go files by processing gml.Object's and assets
        serve		serve a build of your game for playing in a web browser, defaults to port 8080
`)
		os.Exit(1)
	}
	args := os.Args[2:]
	for _, cmd := range cmds {
		if cmd.Name() == os.Args[1] {
			if err := cmd.Run(cmd, args); err != nil {
				panic(err)
			}
			return
		}
	}
	panic("Unable to find command: " + os.Args[1])
}
