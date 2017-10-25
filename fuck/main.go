package main

import "fmt"

func main() {
	var spawns []int
	m := []int{1, 0, 2, 4}
	for i := 0; i < 4; i++ {
		go func() {
			spawns = append(spawns, m[3])
		}()
	}

	fmt.Println(spawns)
}
