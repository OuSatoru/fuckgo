package main

import (
	"fmt"
	"log"
	"os/exec"
	"time"

	"github.com/OuSatoru/fuckgo/common"
)

// use crontab
func main() {
	ticker := time.NewTicker(10 * time.Second)
	for {
		<-ticker.C
		if common.Exists(fmt.Sprintf("/home/db2inst1/baspdback/czzf_main_%s.del", time.Now().Format("20060102"))) {
			continue
		}

		out, err := exec.Command("/home/db2inst1/baspdback/back.sh", time.Now().Format("20060102")).Output()
		if err != nil {
			log.Println(err)
		}
		fmt.Printf("%s\n", out)
	}
}
