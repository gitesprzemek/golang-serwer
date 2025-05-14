package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3" 
)

var DB *sql.DB

func InitDatabase() {
	var err error

	DB, err = sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		log.Fatal("Błąd przy otwieraniu bazy danych:", err)
	}
}
