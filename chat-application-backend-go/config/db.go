// config/db.go
package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() {
	var err error
	DB, err = sql.Open("mysql", "root:root1234@tcp(localhost:3306)/chat_app")
	if err != nil {
		log.Fatal(err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("MySQL connected")
}
