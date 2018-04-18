package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	now := time.Now()
	for i := 0; i < 100000; i++ {
		go func() {
			wg.Add(1)
			time.Sleep(5 * time.Second)
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Printf("%v", time.Since(now))
}
