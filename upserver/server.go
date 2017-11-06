package main

import (
	"net/http"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))
	http.Handle("/tmp/", http.StripPrefix("/tmp/", http.FileServer(http.Dir("G:\\"))))
	http.HandleFunc("/", handler)
	http.HandleFunc("/upload", upload)
	http.ListenAndServe(":2333", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
}

func upload(w http.ResponseWriter, r *http.Request) {

}
