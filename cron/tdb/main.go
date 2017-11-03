package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func main() {
	_, err := exec.Command("db2", "connect to jsbods user ods using ods@98").Output()
	if err != nil {
		log.Println(err)
	}
	out, err := exec.Command("db2", "SELECT * FROM REPORT.ODS_ZBTJB WHERE SJRQ = to_char(CURRENT_DATE - 1 DAY, 'YYYYMMDD') AND ZBDH = '0001'").Output()
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("%s\n", out)
	fmt.Println(!strings.Contains(string(out), "0 record(s) selected"))
}
