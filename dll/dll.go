package main

import (
	"fmt"
	"sync"
	"syscall"
	"time"
)

func main() {
	dll := syscall.NewLazyDLL("embed.dll")
	fmt.Println("Called dll: ", dll.Name)
	rust := time.Now()
	p := dll.NewProc("process")
	p.Call()
	fmt.Println("rust dll: ", time.Since(rust))
	var wg sync.WaitGroup
	gotime := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sum := 0
			for j := 0; j < 5000000; j++ {
				sum++
			}
			fmt.Printf("Goroutine finished with count=%d\n", sum)
		}()
	}
	wg.Wait()
	fmt.Println("go: ", time.Since(gotime))
}
