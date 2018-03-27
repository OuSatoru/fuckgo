package main

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(gen("52a5d4a071d82caea87329868d22f6b8390ac3d227c6fde0d4525e69510ec479"))
	fmt.Println(time.Since(start))
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
				// if num%10000 == 0 {
				// 	fmt.Println(num)
				// }
			}

		}(ch, ans)
	}
	return strconv.Itoa(<-ans)
}
