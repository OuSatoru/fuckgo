package main

import (
	"net/http"
	"regexp"
)

const (
	dir = "F:\\Work\\Rst\\hello-world"
)

var (
	wasmFile = regexp.MustCompile(`\.wasm$`)
)

func main() {
	http.HandleFunc("/wasm/", serveFile)
	http.ListenAndServe(":8080", nil)
}

func serveFile(w http.ResponseWriter, r *http.Request) {
	uri := r.RequestURI
	if wasmFile.MatchString(uri) {
		w.Header().Set("Content-Type", "application/wasm")
	}
	http.StripPrefix("/wasm/", http.FileServer(http.Dir(dir))).ServeHTTP(w, r)
}
