package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/asifjalil/cli"
)

func main() {
	db, err := sql.Open("cli", "DATABASE=JSBODS; HOSTNAME=; PORT=60000; PROTOCOL=TCPIP; UID=; PWD=;")
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
}
