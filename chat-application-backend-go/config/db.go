// config/db.go
package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() {
	var err error
	DB, err = sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/chat_app")
	if err != nil {
		log.Fatal(err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("MySQL connected")
}
