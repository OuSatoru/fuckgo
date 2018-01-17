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
		if exists(fmt.Sprintf("/fr/data/xms/%s.zhdq", yesterday())) {
			continue
		}

		_, err := exec.Command("db2", "connect to jsbods user ods using ods@98").Output()
		if err != nil {
			log.Println(err)
		}
		out, err := exec.Command("db2", "SELECT max(ETL_DT) FROM CBOD.TDACNACN").Output()
		if err != nil {
			log.Println(err)
		}
		out2, err := exec.Command("db2", "SELECT max(ETL_DT) FROM CBOD.LNLNSLNS").Output()
		if err != nil {
			log.Println(err)
		}
		if strings.Contains(string(out), yesterday()) && strings.Contains(string(out2), yesterday()) {
			out, err := exec.Command("./expzhdq.sh").Output()
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
