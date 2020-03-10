package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/a", ahandler)
	http.HandleFunc("/b", bhandler)
	http.ListenAndServe(":8080", nil)
}

func ahandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/b", http.StatusSeeOther)
}

func bhandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", "from b")
}
