package main

import (
	"fmt"
	"time"
)

func main() {
	prime := primes()

	for {
		fmt.Println(<-prime)
		time.Sleep(time.Second)
	}
}

func generate() chan int {
	in := make(chan int)
	go func() {
		for i := 2; ; i++ {
			in <- i
		}
	}()
	return in
}

func filter(in chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				out <- i
			}
		}
	}()
	return out
}

func primes() chan int {
	out := make(chan int)
	go func() {
		ch := generate()
		for {
			prime := <-ch
			ch = filter(ch, prime)
			out <- prime
		}
	}()
	return out
}
