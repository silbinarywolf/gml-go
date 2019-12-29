package serve

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/silbinarywolf/gml-go/cmd/gmlgo/internal/base"
	"github.com/silbinarywolf/gml-go/cmd/gmlgo/internal/shared"
)

var Cmd = &base.Command{
	UsageLine: "serve [dir]",
	Short:     `build and run your game in a browser`,
	Long:      `serve a build of your game for playing in a web browser, defaults to port 8080`,
	Flag:      flag.NewFlagSet("serve", flag.ExitOnError),
	Run:       run,
}

var tags *string

var verbose *bool

var wasmJSPath string

var indexHTMLPath string

func init() {
	tags = Cmd.Flag.String("tags", "", "a list of build tags to consider satisfied during the build")
	verbose = Cmd.Flag.Bool("verbose", false, "verbose")
}

var (
	arguments    Arguments
	tmpOutputDir = ""
)

type Arguments struct {
	Port      string // :8080
	Directory string // .
	Tags      string // ie. "debug"
}

func handle(w http.ResponseWriter, r *http.Request) {
	output, err := ensureTmpOutputDir()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	dir := arguments.Directory
	tags := arguments.Tags

	// Get path and package
	upath := r.URL.Path[1:]
	pkg := filepath.Dir(upath)
	fpath := filepath.Join(".", filepath.Base(upath))
	if strings.HasSuffix(r.URL.Path, "/") {
		fpath = filepath.Join(fpath, "index.html")
	}

	parts := strings.Split(upath, "/")
	isAsset := len(parts) > 0 && parts[0] == "asset"

	if isAsset {
		// Load asset
		log.Print("serving asset: " + upath)

		// todo(Jake): 2018-12-30
		// Improve this so when "data" folder support
		// is added, this allows any filetype from the "data" folder.
		switch ext := filepath.Ext(upath); ext {
		case ".ttf",
			".data",
			".json":
			http.ServeFile(w, r, upath)
		}
		return
	}

	switch filepath.Base(fpath) {
	case "index.html":
		log.Print("serving index.html: " + indexHTMLPath)
		http.ServeFile(w, r, indexHTMLPath)
	case "wasm_exec.js":
		log.Print("serving index.html: " + wasmJSPath)
		http.ServeFile(w, r, wasmJSPath)
		return
	case "main.wasm":
		if _, err := os.Stat(fpath); os.IsNotExist(err) {
			// go build
			args := []string{"build", "-o", filepath.Join(output, "main.wasm")}
			if tags != "" {
				args = append(args, "-tags", tags)
			}
			args = append(args, pkg)
			log.Print("go ", strings.Join(args, " "))
			cmdBuild := exec.Command(gobin(), args...)
			cmdBuild.Env = append(os.Environ(), "GOOS=js", "GOARCH=wasm")
			cmdBuild.Dir = dir
			out, err := cmdBuild.CombinedOutput()
			if err != nil {
				log.Print(err)
				log.Print(string(out))
				http.Error(w, string(out), http.StatusInternalServerError)
				return
			}
			if len(out) > 0 {
				log.Print(string(out))
			}

			f, err := os.Open(filepath.Join(output, "main.wasm"))
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			defer f.Close()
			http.ServeContent(w, r, "main.wasm", time.Now(), f)
			return
		}
	}
}

func gobin() string {
	return filepath.Join(runtime.GOROOT(), "bin", "go")
}

func ensureTmpOutputDir() (string, error) {
	if tmpOutputDir != "" {
		return tmpOutputDir, nil
	}

	tmp, err := ioutil.TempDir("", "")
	if err != nil {
		return "", err
	}
	tmpOutputDir = tmp
	return tmpOutputDir, nil
}

func run(cmd *base.Command, args []string) error {
	cmd.Flag.Parse(args)
	if !cmd.Flag.Parsed() {
		cmd.Flag.PrintDefaults()
		os.Exit(1)
	}
	args = cmd.Flag.Args()
	dir := ""
	if len(args) > 0 {
		dir = args[0]
	}
	return Run(Arguments{
		Directory: dir,
		Tags:      *tags,
	})
}

func Run(args Arguments) error {
	// Setup globals
	if args.Port == "" {
		args.Port = ":8080"
	}
	if args.Directory == "" {
		args.Directory = "."
	}
	arguments = args

	// Validation of settings
	dir := args.Directory
	if dir != "." {
		panic("Specifying a custom directory is not currently supported.")
	}

	// Get default resources
	var err error
	wasmJSPath, err = shared.GetDefaultWasmJSPath(args.Directory)
	if err != nil {
		panic(err)
	}
	log.Printf("wasm_exec.js: %s\n", wasmJSPath)
	indexHTMLPath, err = shared.GetDefaultIndexHTMLPath(args.Directory)
	if err != nil {
		panic(err)
	}
	log.Printf("index.html: %s\n", indexHTMLPath)

	// Start server
	log.Printf("listening on %q...", args.Port)
	http.HandleFunc("/", handle)
	shared.OpenBrowser("http://localhost" + args.Port)
	log.Fatal(http.ListenAndServe(args.Port, nil))
	return nil
}
