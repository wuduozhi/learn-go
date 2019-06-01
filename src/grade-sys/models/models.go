package models

import (
	"grade-sys/pkg/logging"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
)

var Db *sql.DB

func init(){
	var err error
	Db,err = sql.Open("mysql","weihuda:7Hy9ZsN8eLmZDhX6@tcp(qnxg.net:3306)/weihuda?parseTime=true")
	if err != nil {
		logging.Fatal(err)
	}

	return
}

func checkErr(err error){
	if err != nil {
		logging.Debug(err)
	}
}
