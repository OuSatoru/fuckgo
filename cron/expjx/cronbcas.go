package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	go task1("jx", "SELECT count(*) FROM BCAS.D_PER_ACHV_%s WHERE ETLDT = %s")
	go task1("jxsjyh", "SELECT count(*) FROM BCAS.D_PER_ACHV_%s WHERE ETLDT = %s AND FORMULA_CODE LIKE '57%%'")
	go task1("jxdk", "SELECT count(*) FROM BCAS.D_PER_ACHV_%s WHERE ETLDT = %s AND FORMULA_CODE = '11000000'")
	wg.Wait()
}

func task1(suffix, statusy string) {

	ticker := time.NewTicker(10 * time.Second)
	for {
		<-ticker.C
		status := fmt.Sprintf(statusy, time.Now().Format("2006"), "current_date - 1 day")
		zeros := fmt.Sprintf(statusy, time.Now().Format("2006"), "'"+time.Now().AddDate(1, 0, 0).Format("2006-01-02")+"'")
		if exists(fmt.Sprintf("/fr/data/xms/%s.%s", yesterday(), suffix)) || hour() < 7 {
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
