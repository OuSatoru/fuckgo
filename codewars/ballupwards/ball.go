// https://www.codewars.com/kata/ball-upwards/train/go
package main

import "fmt"

func MaxBall(v0 int) int {
	h := 0.0
	var i int
	for i = 1; ; i++ {
		if hight(v0, i) > h {
			h = hight(v0, i)
		} else {
			break
		}
	}
	return i - 1
}

func hight(v, t int) float64 {
	return float64(v)/3.6*float64(t)*0.1 - 0.5*9.81*float64(t)*float64(t)*0.1*0.1
}

func main() {
	fmt.Println(MaxBall(25))
}
