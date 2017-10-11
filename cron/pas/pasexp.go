package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

func main() {
	ticker := time.NewTicker(10 * time.Second)
	for {
		<-ticker.C
		yes := yesterday()
		if exists(fmt.Sprintf("/fr/data/xms/%s.pas", yes)) {
			continue
		}

		out, err := exec.Command("./pasexp.sh").Output()
		if err != nil {
			log.Println(err)
		}
		fmt.Printf("%s\n", out)

	}

}

func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		// os.IsNotExist(err)
		return false
	}
	return true
}

func yesterday() string {
	now := time.Now()
	return now.AddDate(0, 0, -1).Format("20060102")
}
