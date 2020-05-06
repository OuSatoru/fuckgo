package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	ur, err := url.Parse("https://www.thetvdb.com")
	if err != nil {
		log.Fatal(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(ur)
	if err := http.ListenAndServe(":19999", proxy); err != nil {
		log.Fatal(err)
	}
}
