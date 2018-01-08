package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

func main() {
	ticker := time.NewTicker(30 * time.Second)
	for {
		<-ticker.C
		if exists(fmt.Sprintf("/fr/data/xms/%s.endzb", yesterday())) {
			continue
		}
		if exists(fmt.Sprintf("/fr/data/xms/%s.zb", yesterday())) {
			log.Println("importing zb data into frxms")
			out, err := exec.Command("./impzb.sh").Output()
			if err != nil {
				log.Println(err)
			}
			fmt.Printf("%s\n", out)
		}
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
