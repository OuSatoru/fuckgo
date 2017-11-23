package main

import "net/http"

func main() {
	http.HandleFunc("/login", loginHandler)
	http.ListenAndServe(":2333", nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {

}
