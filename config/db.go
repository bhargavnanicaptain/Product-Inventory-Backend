package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {

	var err error
	dsn := "root:123Nani321@@tcp(localhost:3306)/productdb?parseTime=true"
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(" Failed to open the DB: ", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("Failed to connect to DB: ", err)
	}
	log.Println("Connected to MySQL")

}
