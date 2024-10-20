package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() {
	var err error
	DB, err = sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/dbname")
	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MySQL!")
}
