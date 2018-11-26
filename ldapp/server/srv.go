package main

import (
	"fmt"
	"net/http"

	"github.com/OuSatoru/fuckgo/ldapp"
)

func main() {
	http.HandleFunc("/verify", verifyHandler)
	http.ListenAndServe(":7474", nil)
}

func verifyHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	user := r.Form.Get("user")
	pwd := r.Form.Get("password")
	err := ldapp.VerifyUser(user, pwd)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	fmt.Fprint(w, 1)
}
