package main

import (
	"fmt"
	"log"
	"os/exec"
	"sync"
	"time"

	"github.com/OuSatoru/fuckgo/common"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	go task("rh")
	go task("yq")
	wg.Wait()
}

func task(suffix string) {
	ticker := time.NewTicker(30 * time.Second)
	for {
		<-ticker.C
		if common.Exists(fmt.Sprintf("/fr/data/xms/%s.end%s", common.Yesterday(), suffix)) {
			continue
		}
		if common.Exists(fmt.Sprintf("/fr/data/xms/%s.%s", common.Yesterday(), suffix)) {
			out, err := exec.Command(fmt.Sprintf("./imp%s.sh", suffix)).Output()
			if err != nil {
				log.Println(err)
			}
			fmt.Printf("%s\n", out)
		}
	}
}
