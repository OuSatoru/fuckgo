package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"time"
)

func main() {
	ticker := time.NewTicker(10 * time.Second)
	for {
		<-ticker.C
		if exists(fmt.Sprintf("/fr/data/xms/%s.bl", yesterday())) {
			continue
		}
		_, err := exec.Command("db2", "connect to jsbods user ods using ods@98").Output()
		if err != nil {
			log.Println(err)
		}
		out, err := exec.Command("db2", "SELECT SJRQ, SJZT FROM REPORT.FDM_SJRQ").Output()
		if err != nil {
			log.Println(err)
		}
		match, _ := regexp.MatchString(time.Now().Format("20060102")+`\s+1`, string(out))
		// fmt.Printf("%s\n", out)
		if match {
			out, err := exec.Command("./expbl.sh").Output()
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
