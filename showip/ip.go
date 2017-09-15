package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/ip", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, r.RemoteAddr)
	})
	http.ListenAndServe(":2333", nil)
}
