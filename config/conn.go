package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connection() {
	Db, _ := sql.Open("mysql", "root:@/tesweb1")

	if erro := Db.Ping(); erro != nil {
		log.Fatal(Db)
	}

	log.Println("Connected to the database")
	DB = Db
}
