package model

import (
	"database/sql"
	"fmt"
)

var DB *sql.DB

func GetDBConnection() {
	db, err := sql.Open("sqlite3", "./DB/foo.db")

	if err != nil {
		fmt.Println("DB connection failed: ", err)
	}
	DB = db
}
