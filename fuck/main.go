package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s - %v", r.RemoteAddr, r.Header)
	})
	http.ListenAndServe(":2333", nil)
}
