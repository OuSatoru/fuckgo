package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	_, err := exec.Command("db2", "connect to bcas_dt user bcas using `1qaz").Output()
	if err != nil {
		log.Println(err)
	}
	out, err := exec.Command("db2", "SELECT count(*) FROM BCAS.D_PER_ACHV_2018 WHERE ETLDT = CURRENT_DATE").Output()
	if err != nil {
		log.Println(err)
	}
	zero, err := exec.Command("db2", "SELECT count(*) FROM BCAS.D_PER_ACHV_2018 WHERE ETLDT = '2019-01-01'").Output()
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("%s\n", out)
	fmt.Println(string(out) == string(zero))
}
