package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func main() {
	ticker := time.NewTicker(10 * time.Second)
	for {
		<-ticker.C
		if exists(fmt.Sprintf("/fr/data/xms/%s.jx", yesterday())) || hour() < 7 {
			continue
		}
		_, err := exec.Command("db2", "connect to bcas_dt user bcas using `1qaz").Output()
		if err != nil {
			log.Println(err)
		}
		num, err := exec.Command("db2", "SELECT count(*) FROM BCAS.D_PER_ACHV_2018 WHERE ETLDT = CURRENT_DATE - 1 DAY").Output()
		if err != nil {
			log.Println(err)
		}
		zero, err := exec.Command("db2", "SELECT count(*) FROM BCAS.D_PER_ACHV_2018 WHERE ETLDT = '2019-01-01'").Output()
		if err != nil {
			log.Println(err)
		}
		// fmt.Printf("%s\n", out)
		if string(num) != string(zero) {
			out, err := exec.Command("./expjx.sh").Output()
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

func hour() int {
	h, err := strconv.Atoi(time.Now().Format("15"))
	if err != nil {
		log.Fatal(err)
	}
	return h
}

func yesterday() string {
	now := time.Now()
	return now.AddDate(0, 0, -1).Format("20060102")
}
