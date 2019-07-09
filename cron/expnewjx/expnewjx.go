package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	go task("jxdzyhrb", "select count(*) from E_DZYH_ZTLDJSCS where ETLDT = %s")
	wg.Wait()
}

func task(suffix, statusy string) {
	ticker := time.NewTicker(20 * time.Second)
	for {
		<-ticker.C
		status := fmt.Sprintf(statusy, "current_date - 1 day")
		zeros := fmt.Sprintf(statusy, "'"+time.Now().AddDate(1, 0, 0).Format("2006-01-02")+"'")
		if exists(fmt.Sprintf("/fr/data/xms/%s.end%s", yesterday(), suffix)) {
			continue
		}
		_, err := exec.Command("db2", "connect to bcas user bcas using bcas").Output()
		if err != nil {
			log.Println(err)
		}
		num, err := exec.Command("db2", status).Output()
		if err != nil {
			log.Println(err)
		}
		log.Println(status)
		zero, err := exec.Command("db2", zeros).Output()
		if err != nil {
			log.Println(err)
		}
		log.Println(zeros)
		// fmt.Printf("%s\n", out)
		if string(num) != string(zero) {
			out, err := exec.Command(fmt.Sprintf("./exp%s.sh", suffix)).Output()
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
