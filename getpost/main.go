package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Reply struct {
	Method string
	Params string
}

func main() {
	http.HandleFunc("/get", hGet)
	http.HandleFunc("/post", hPost)
	http.ListenAndServe(":2333", nil)
}

func hGet(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fstr := "Params: "
	if len(r.Form) != 0 {
		for k, v := range r.Form {
			fstr += k + "=" + strings.Join(v, ", ") + " "
		}
	}
	// fmt.Fprintf(w, "%s", fstr)
	reply := Reply{Method: "Get", Params: fstr}
	b, err := json.MarshalIndent(reply, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "%s", string(b))
}

func hPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fstr := "Params: "
	if len(r.Form) != 0 {
		for k, v := range r.Form {
			fstr += k + "=" + strings.Join(v, ", ") + " "
		}
	}
	// fmt.Fprintf(w, "%s", fstr)
	reply := Reply{Method: "Post", Params: fstr}
	b, err := json.MarshalIndent(reply, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(r.RemoteAddr)
	fmt.Fprintf(w, "%s", string(b))
}
