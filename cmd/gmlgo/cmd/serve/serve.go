package serve

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

const indexHTML = `<!DOCTYPE html>
<head>
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<style>
		* {
			box-sizing: border-box;
		}

		body,
		html {
			height: 100%;
			margin: 0;
			padding: 0;
		}

		body {
			position: relative;
			z-index: 0;
			color: #fff; 
			background-color: #000; 
		}

		pre {
			margin: 5px;
		}

		.load-info {
			position: absolute;
			left: calc((100% - 640px) / 2); 
			top: calc((100% - 480px) / 2);
			z-index: 0;
			width: 640px;
			height: 480px;
			font-size: 32px;
			text-align: center; 
			vertical-align: middle;
		}
	</style>
</head>
<body>
	<div class="load-info">
		<p>Loading...</p>
	</div>
	<script src="wasm_exec.js"></script>
	<script>
	(async () => {
	  const infoEl = document.body.querySelector('.load-info');
	  const textEl = infoEl.querySelector('p');
	  textEl.textContent = 'Compiling...';
	  const resp = await fetch('main.wasm');
	  if (infoEl) {
	  	infoEl.style.display = "none";
	  }
	  if (!resp.ok) {
	    const pre = document.createElement('pre');
	    pre.innerText = await resp.text();
	    document.body.appendChild(pre);
	    return;
	  }
	  const src = await resp.arrayBuffer();
	  const go = new Go();
	  const result = await WebAssembly.instantiate(src, go.importObject);
	  go.run(result.instance);
	})();
	</script>
</body>
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

func Run(args Arguments) {
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
}
