package serve

import (
	"bytes"
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

func init() {
	tags = Cmd.Flag.String("tags", "", "a list of build tags to consider satisfied during the build")
	verbose = Cmd.Flag.Bool("verbose", false, "verbose")
}

const indexHTML = `<html>
	<head>
		<meta charset="utf-8">
		<style>
			body {
				background-color: #000;
				color: #fff;
				font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Helvetica, Arial, sans-serif, "Apple Color Emoji", "Segoe UI Emoji", "Segoe UI Symbol";
			}

			.error-container {
				background-color: #F47F7F;
				border: 1px solid #531212;
				color: #000;
				padding: 10px 10px;
			}
			.container {
				width: 1280px;
				max-width: 100%;
				margin: 0 auto;
			}
			.progress-bar-wrapper {
				width: 100%;
				height: 10px;
				border: 1px solid #266926;
			}
			.progress-bar {
				width: 0%;
				height: inherit;
				background-color: #46C346;
			}
		</style>
	</head>
	<body>
		<div class="container">
			<p>Loading...</p>
			<div class="progress-bar-wrapper">
				<div class="progress-bar"></div>
			</div>
		</div>
		<script src="wasm_exec.js"></script>
		<script>
			window.onerror = function(message, source, lineno, colno, error) {
				let el = document.createElement("div");
				el.classList.add("error-container");
				let newContent = document.createTextNode(message); 
				el.appendChild(newContent);
				document.body.appendChild(el);
			}
			let goRequest;
			let progressBar = document.body.querySelector(".progress-bar");

			if (!progressBar) {
				throw new Error("Missing .progress-bar.")
			}
			function setProgressBar(percent) {
				progressBar.style.width = String(percent) + "%";
			}
			let filesProgress = {};
			function updateProgressBar() {
				if (!progressBar) {
					return;
				}
				let percent = 0;
				let i = 0;
				for (let key in filesProgress) {
					if (!filesProgress.hasOwnProperty(key)) {
						continue;
					}
					percent += filesProgress[key]
					i++;
				}
				const totalPercent = percent / i;
				setProgressBar(totalPercent);
				if (totalPercent < 100) {
					return;
				}
				const go = new Go();
				WebAssembly.instantiate(goRequest.response, go.importObject).then((result) => {
					while (document.body.hasChildNodes()) {
						document.body.removeChild(document.body.childNodes[0]);
					}
					go.run(result.instance);
				});
			}
			function preloadFile(path) {
				const preloadLink = document.createElement("link");
				preloadLink.href = path;
				preloadLink.rel = "preload";
				preloadLink.as = "fetch";
				filesProgress[path] = 0;
				preloadLink.addEventListener("progress", function (e) {
					if (e.lengthComputable) {
						const percent = (e.loaded / e.total * 100 | 0);
						filesProgress[path] = percent;
						updateProgressBar();
					}
				});
				preloadLink.addEventListener("load", function (e) {
					filesProgress[path] = 100;
					updateProgressBar();
				});
				document.head.appendChild(preloadLink);
			}
			function getBinary() {
				let fullAssetName = "main.wasm";
				filesProgress[fullAssetName] = 0;
				let request = new XMLHttpRequest();
				goRequest = request;
				request.addEventListener("progress", function (e) {
					if (e.lengthComputable) {
						let percent = (e.loaded / e.total * 100 | 0);
						filesProgress[fullAssetName] = percent;
						updateProgressBar();
					}
				});
				request.addEventListener("load", function () {
					if (request.status !== 200) {
						throw new Error(request.status + " " + request.statusText);
					}
					filesProgress[fullAssetName] = 100;
					updateProgressBar();
				});
				request.responseType = "arraybuffer";
				request.open("GET", fullAssetName);
				request.setRequestHeader("X-Requested-With", "XMLHttpRequest");
				request.send();
				return request;
			}
			function getManifest() {
				let request = new XMLHttpRequest();
				request.overrideMimeType("application/json");
				request.open("GET", "asset/manifest.json");
				request.setRequestHeader("X-Requested-With", "XMLHttpRequest");
				request.addEventListener("load", function() {
					let jsonResponse = request.response;
					let json = JSON.parse(jsonResponse);

					getBinary();
					for (let groupName in json) {
						if (!json.hasOwnProperty(groupName)) {
							continue;
						}
						let group = json[groupName];
						for (let key in group) {
							if (!group.hasOwnProperty(key)) {
								continue;
							}
							let name = group[key];
							let fullAssetName = "asset/" + groupName + "/" + name + ".data";
							preloadFile(fullAssetName);
						}
					}
				});
				request.send();
			}
			getManifest();
		</script>
	</body>
</html>
`

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

	if dir != "." {
		panic("Specifying a custom directory is not currently supported.")
	}

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
		if _, err := os.Stat(fpath); os.IsNotExist(err) {
			http.ServeContent(w, r, "index.html", time.Now(), bytes.NewReader([]byte(indexHTML)))
			return
		}
	case "wasm_exec.js":
		if _, err := os.Stat(fpath); os.IsNotExist(err) {
			f := filepath.Join(runtime.GOROOT(), "misc", "wasm", "wasm_exec.js")
			http.ServeFile(w, r, f)
			return
		}
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
	if args.Port == "" {
		args.Port = ":8080"
	}
	if args.Directory == "" {
		args.Directory = "."
	}
	arguments = args
	log.Printf("listening on %q...", args.Port)
	http.HandleFunc("/", handle)
	log.Fatal(http.ListenAndServe(args.Port, nil))
	return nil
}
