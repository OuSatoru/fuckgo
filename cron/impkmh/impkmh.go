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

var wg sync.WaitGroup

func main() {
	tablesf, err := ioutil.ReadFile("/fr/data/xms/kmhtables.txt")
	if err != nil {
		log.Fatal(err)
	}
	tables := strings.Split(string(tablesf), " ")
	fmt.Println(tables)
	wg.Add(1)
	for _, table := range tables {
		go task(strings.Trim(table, " "))
	}
	wg.Wait()
}

func task(tablename string) {
	ticker := time.NewTicker(10 * time.Second)
	for {
		<-ticker.C
		if common.Exists(fmt.Sprintf("/fr/data/xms/kmhexp/%s.end%s", common.Yesterday(), tablename)) {
			continue
		}
		if common.Exists(fmt.Sprintf("/fr/data/xms/kmhexp/%s.%s", common.Yesterday(), tablename)) {
			out, err := exec.Command("./impkmh.sh", tablename).Output()
			if err != nil {
				log.Println(err)
			}
			fmt.Printf("%s\n", out)
		}
	}
}
