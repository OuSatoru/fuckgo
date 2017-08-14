package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"sort"
	"strconv"

	"golang.org/x/net/websocket"

	_ "github.com/mattn/go-sqlite3"
)

const (
	earthR = 6378137.0
)

var db *sql.DB

type Branches struct {
	Branch []Branch `json:"branches"`
}

type Branch struct {
	ID        int     `json:"id"`
	Distance  float64 `json:"distance"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
	Name      string  `json:"name"`
}

func init() {
	var err error
	db, err = sql.Open("sqlite3", "./branch.db")
	if err != nil {
		panic(err)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/weapp/nearest", nearestBranch)
	mux.Handle("/ws", websocket.Handler(wxHandler))
	mux2 := http.NewServeMux()
	mux2.HandleFunc("/weapp", handler)
	// go http.ListenAndServeTLS(":443", "./cert/server.pem", "./cert/server.key", mux)
	go http.ListenAndServe(":80", mux)
	http.ListenAndServe(":8080", mux2)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hhhh")
}

func nearestBranch(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	longitude, err := strconv.ParseFloat(r.Form.Get("longitude"), 64)
	if err != nil {
		log.Println("Parse lng err")
		return
	}
	latitude, err := strconv.ParseFloat(r.Form.Get("latitude"), 64)
	if err != nil {
		log.Println("Parse lat err")
		return
	}
	var branches []Branch
	rows, err := db.Query(`SELECT lng_gcj02, lat_gcj02, abbr FROM branches WHERE lat_gcj02 IS NOT NULL`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var lng float64
		var lat float64
		var abbr string
		err = rows.Scan(&lng, &lat, &abbr)
		if err != nil {
			panic(err)
		}
		var branch Branch
		branch.Distance = between(longitude, latitude, lng, lat)
		branch.Name = abbr
		branch.Longitude = lng
		branch.Latitude = lat
		branches = append(branches, branch)
	}
	sort.Slice(branches, func(i, j int) bool { return branches[i].Distance < branches[j].Distance })
	for i := range branches {
		branches[i].ID = i + 1
	}
	var bs Branches
	bs.Branch = branches
	b, err := json.MarshalIndent(bs, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Fprint(w, string(b))
}

func redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://"+r.Host+r.URL.String(), http.StatusMovedPermanently)
}

func between(lng1, lat1, lng2, lat2 float64) float64 {
	pi180 := math.Pi / 180
	arcLat1 := lat1 * pi180
	arcLat2 := lat2 * pi180
	x := math.Cos(arcLat1) * math.Cos(arcLat2) * math.Cos((lng1-lng2)*pi180)
	y := math.Sin(arcLat1) * math.Sin(arcLat2)
	s := x + y
	if s > 1 {
		s = 1
	}
	if s < -1 {
		s = -1
	}
	alpha := math.Acos(s)
	between := alpha * earthR
	return between
}
