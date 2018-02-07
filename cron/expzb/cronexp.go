package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {
	ticker := time.NewTicker(10 * time.Second)
	for {
		<-ticker.C
		if exists(fmt.Sprintf("/fr/data/xms/%s.zb", yesterday())) {
			continue
		}
		_, err := exec.Command("db2", "connect to jsbods user ods using ods@98").Output()
		if err != nil {
			log.Println(err)
		}
		out, err := exec.Command("db2", "SELECT max(SJRQ) FROM REPORT.ODS_HXZZYEB").Output()
		if err != nil {
			log.Println(err)
		}
		// fmt.Printf("%s\n", out)
		if strings.Contains(string(out), yesterday()) {
			out, err := exec.Command("./expzb.sh").Output()
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
