package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"text/tabwriter"

	"github.com/silbinarywolf/gml-go/cmd/gmlgo/internal/base"
	"github.com/silbinarywolf/gml-go/cmd/gmlgo/internal/build"
	"github.com/silbinarywolf/gml-go/cmd/gmlgo/internal/generate"
	"github.com/silbinarywolf/gml-go/cmd/gmlgo/internal/publish"
	"github.com/silbinarywolf/gml-go/cmd/gmlgo/internal/serve"
	"github.com/silbinarywolf/gml-go/cmd/gmlgo/internal/shared"
)

var cmds = []*base.Command{
	build.Cmd,
	generate.Cmd,
	serve.Cmd,
	publish.Cmd,
	// todo(Jake): 2019-08-10
	//
	// assetpack	[todo] make this build asset files
}

func main() {
	log.SetFlags(0)
	log.SetPrefix(shared.RootCmd + ": ")

	if len(os.Args) < 2 {
		fmt.Printf(`
GmlGo is a tool for building games using the GmlGo library

Usage:

        gmlgo <command> [arguments]

The commands are:

`)
		cmdsSorted := append([]*base.Command{}, cmds...)
		sort.Slice(cmdsSorted[:], func(i, j int) bool {
			// sort alphabetically
			return cmdsSorted[i].UsageLine < cmdsSorted[j].UsageLine
		})
		w := tabwriter.NewWriter(os.Stdout, 4, 4, 4, '\t', tabwriter.AlignRight)
		for _, cmd := range cmdsSorted {
			cmdName := strings.SplitN(cmd.UsageLine, " ", 2)[0]
			fmt.Fprintln(w, "\t"+cmdName+"\t"+cmd.Short)
		}
		w.Flush()
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
	fmt.Println("Unable to find command: " + os.Args[1])
	os.Exit(1)
}
