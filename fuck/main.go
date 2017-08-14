package main

import (
	"encoding/json"
	"fmt"
)

type Branches struct {
	Branch []Branch
}

type Branch struct {
	Id       int     `json:"id"`
	Distance float64 `json:"distance"`
	Name     string  `json:"name"`
}

func main() {
	var branch Branch
	branch.Id = 1
	branch.Distance = 3.4
	branch.Name = "dfsf"
	var branches []Branch
	branches = append(branches, branch)
	var bs Branches
	bs.Branch = branches
	b, _ := json.Marshal(branch)
	fmt.Println(string(b))
	fmt.Println(branch)
}
