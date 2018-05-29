package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	go task1("rh", "SELECT max(sjrq) FROM REPORT.ODS_RHRBTJB")
	go task2("yq", "select jssj from report.fdm_sjrq")
	wg.Wait()
}

func task1(suffix, status string) {
	ticker := time.NewTicker(10 * time.Second)
	for {
		<-ticker.C
		if exists(fmt.Sprintf("/fr/data/xms/%s.%s", yesterday(), suffix)) {
			continue
		}
		_, err := exec.Command("db2", "connect to jsbods user ods using ods@98").Output()
		if err != nil {
			log.Println(err)
		}
		out, err := exec.Command("db2", status).Output()
		if err != nil {
			log.Println(err)
		}
		// fmt.Printf("%s\n", out)
		if strings.Contains(string(out), yesterday()) {
			time.Sleep(2 * time.Second)
			out, err := exec.Command(fmt.Sprintf("./exp%s.sh", suffix)).Output()
			if err != nil {
				log.Println(err)
			}
			fmt.Printf("%s\n", out)
		}
	}
}

func task2(suffix, status string) {
	ticker := time.NewTicker(10 * time.Second)
	for {
		<-ticker.C
		if exists(fmt.Sprintf("/fr/data/xms/%s.%s", yesterday(), suffix)) {
			continue
		}
		_, err := exec.Command("db2", "connect to jsbods user ods using ods@98").Output()
		if err != nil {
			log.Println(err)
		}
		out, err := exec.Command("db2", status).Output()
		if err != nil {
			log.Println(err)
		}
		// fmt.Printf("%s\n", out)
		if strings.Contains(string(out), time.Now().Format("2006-01-02")) {
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
