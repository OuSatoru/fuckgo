package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/alexbrainman/odbc"
)

func main() {
	db, err := sql.Open("odbc", fmt.Sprintf("driver={%s};hostname=%s;port=%d;database=%s;uid=%s;pwd=%s",
		"IBM DATA SERVER DRIVER for ODBC - F:/Work/clidriver", "ip", 60000, "db", "user", "using"))
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query("SELECT hm FROM REPORT.ODS_HQCKZD WHERE ZJHM = '320924199107260271'")
	if err != nil {
		fmt.Println("aa")
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var item string
		if err := rows.Scan(&item); err != nil {
			log.Fatal(err)
		}
		fmt.Println(item)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	// db.Exec("DELETE FROM REPORT.TEST5")
}

// func gbkToUTF8(g string) string {
// 	reader := transform.NewReader(strings.NewReader(g), simplifiedchinese.GBK.NewDecoder())
// 	d, _ := ioutil.ReadAll(reader)
// 	return string(d)
// }
