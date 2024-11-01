package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var database *sql.DB

func InitializeDatabase() {
	db, err := sql.Open("sqlite3", "./chat_history.db")
	if err != nil {
		log.Fatal(err)
	}

	Migrate(db)

	database = db
}

func GetDatabase() *sql.DB {
	if database == nil {
		InitializeDatabase()
	}

	return database
}

func Migrate(db *sql.DB) {
	createTable := `CREATE TABLE IF NOT EXISTS chat_history (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		sender TEXT,
		message TEXT,
		response TEXT
	);`
	_, err := db.Exec(createTable)
	if err != nil {
		log.Fatal(err)
	}
}
