package database

import (
	"database/sql"
	"log"

	"github.com/kirboyyy/smartwiki/config"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", config.DBPath)
	if err != nil {
		log.Fatal(err)
	}

	// Run migrations
	createTables()
}

func createTables() {
	query := `
    CREATE TABLE IF NOT EXISTS pages (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT,
        content TEXT,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP
    );
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        username TEXT,
        password TEXT,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP
    );`
	_, err := DB.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}
