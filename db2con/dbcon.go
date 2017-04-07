package main

import (
	"database/sql"
	"golang.org/x/text/encoding/simplifiedchinese"
	"fmt"
	"log"
	"golang.org/x/text/transform"
	"strings"
	"io/ioutil"

	_ "github.com/alexbrainman/odbc"
)

func main() {
	db, err := sql.Open("odbc", fmt.Sprintf("driver={%s};hostname=%s;port=%d;database=%s;uid=%s;pwd=%s",
		"IBM DATA SERVER DRIVER for ODBC - F:/Work/db/clidriver", "32.185.20.54", 60000, "JSBODS", "ods", "ods"))
	if err != nil {
		log.Fatal(err)
	}
	//rows, err := db.Query("SELECT hm FROM REPORT.ODS_HQCKZD WHERE ZJHM = '320924199107260271'")
	//if err != nil {
	//	fmt.Println("aa")
	//	log.Fatal(err)
	//}
	//defer rows.Close()
	//for rows.Next() {
	//	var item string
	//	if err := rows.Scan(&item); err != nil {
	//		log.Fatal(err)
	//	}
	//	fmt.Println(gbkToUTF8(item))
	//}
	//if err := rows.Err(); err != nil {
	//	log.Fatal(err)
	//}
	db.Exec("DELETE FROM REPORT.TEST5")
}

func gbkToUTF8(g string) string {
	reader := transform.NewReader(strings.NewReader(g), simplifiedchinese.GBK.NewDecoder())
	d, _ := ioutil.ReadAll(reader)
	return string(d)
}
