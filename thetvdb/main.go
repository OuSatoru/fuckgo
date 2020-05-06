package main

import (
	"io"
	"log"
	"net/http"
)

var client = http.DefaultClient

func main() {
	http.HandleFunc("/", tvHandler)
	if err := http.ListenAndServe(":19999", nil); err != nil {
		log.Fatal(err)
	}
}

func tvHandler(w http.ResponseWriter, r *http.Request) {
	body := r.Body
	req, err := http.NewRequest(r.Method, "https://www.thetvdb.com"+r.RequestURI, body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	for k, vv := range r.Header {
		for _, v := range vv {
			req.Header.Add(k, v)
		}
	}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	for k, vv := range resp.Header {
		for _, v := range vv {
			w.Header().Add(k, v)
		}
	}
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}
