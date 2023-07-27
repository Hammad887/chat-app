package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitializeDB() {
	var err error
	DB, err = sql.Open("mysql", "hammad:Hammad_887@/chatapp")
	if err != nil {
		panic(err)
	}
}
