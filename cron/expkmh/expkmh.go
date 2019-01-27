package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
	"sync"
	"time"

	"github.com/OuSatoru/fuckgo/common"
)

var (
	wg sync.WaitGroup
)

func main() {
	tablesf, err := ioutil.ReadFile("/fr/data/xms/kmhtables.txt")
	if err != nil {
		log.Fatal(err)
	}
	tables := strings.Split(string(tablesf), " ")
	fmt.Println(tables)
	wg.Add(1)
	for _, table := range tables {
		go exptable(strings.Trim(table, " "))
	}
	wg.Wait()
}

func exptable(tablename string) {
	ticker := time.NewTicker(10 * time.Second)
	for {
		<-ticker.C
		status := fmt.Sprintf("select count(1) from bcas.%s where ETLDT = '%s'", tablename, time.Now().AddDate(0, 0, -1).Format("2006-01-02"))
		zeros := fmt.Sprintf("select count(1) from bcas.%s where ETLDT = '%s'", tablename, time.Now().AddDate(0, 0, 1).Format("2006-01-02"))
		if common.Exists(fmt.Sprintf("/fr/data/xms/kmhexp/%s.%s", common.Yesterday(), tablename)) {
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
		if string(num) != string(zero) {
			out, err := exec.Command("./expkmh.sh", tablename).Output()
			if err != nil {
				log.Println(err)
			}
			fmt.Printf("%s\n", out)
		}
	}
}
