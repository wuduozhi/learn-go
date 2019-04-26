package models

import (
	"log"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
)

var Db *sql.DB

func init(){
	var err error
	Db,err = sql.Open("mysql","root:WudUozhI@tcp(127.0.0.1:3306)/email_sys?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	return
}

