package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now().Add(-9*time.Hour).AddDate(0, 0, -1))
}
