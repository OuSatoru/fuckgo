package main

import (
	"crypto/sha256"
	"fmt"
	"strconv"
)

func main() {

}

func gen(sha string) string {
	// nonce
	i := 900000000
	ch := make(chan int)
	ans := make(chan int)
	go func() {
		for {
			ch <- i
			i++
		}
	}()

	for count := 0; count < 10; count++ {
		go func(c, a chan int) {
			for {
				sh := sha256.New()
				num := <-ch
				sh.Write([]byte(strconv.Itoa(num)))
				shCalc := fmt.Sprintf("%x", sh.Sum(nil))
				if shCalc == sha {
					a <- num
					break
				}
				if num%10000 == 0 {
					fmt.Println(num)
				}
			}

		}(ch, ans)
	}
	return strconv.Itoa(<-ans)
}
