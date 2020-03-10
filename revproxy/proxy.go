package main

import (
	"io"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	ur, err := url.Parse("http://127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(ur)
	go http.ListenAndServeTLS(":8083", "./keys/3572252_cloud.shuoyu.wang.pem", "./keys/3572252_cloud.shuoyu.wang.key", proxy)
	go http.ListenAndServe(":8082", http.HandlerFunc(redirectTLS))
	forwarder()
}

func redirectTLS(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://cloud.shuoyu.wang:57777"+r.RequestURI, http.StatusMovedPermanently)
}

func forwarder() {
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}
	for {
		local, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		sniff := make([]byte, 3)
		n, err := local.Read(sniff)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
		s := string(sniff[:n])

		var addr string
		if s == "GET"[:n] || s == "POS"[:n] {
			addr = ":8082"
		} else {
			addr = ":8083"
		}

		remote, err := net.Dial("tcp", addr)
		if err != nil {
			log.Fatal(err)
		}
		go io.Copy(local, remote)
		go func() {
			remote.Write(sniff[:n])
			io.Copy(remote, local)
		}()
	}
}
