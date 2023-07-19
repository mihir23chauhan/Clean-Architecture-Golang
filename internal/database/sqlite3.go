package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func InitializeDB(file string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return nil, err
	}

	if !isTableExists(db, "books") {
		err = createBooksTable(db)
		if err != nil {
			db.Close()
			return nil, err
		}
	}

	return db, nil
}

func isTableExists(db *sql.DB, tableName string) bool {
	query := "SELECT name FROM sqlite_master WHERE type='table' AND name=?"
	var name string
	err := db.QueryRow(query, tableName).Scan(&name)
	if err != nil {
		if err == sql.ErrNoRows {
			return false // Table does not exist
		}
		log.Printf("Failed to check table existence: %v", err)
		return false
	}
	return true
}

func createBooksTable(db *sql.DB) error {
	createTable := `
		CREATE TABLE books (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT,
			author TEXT,
			publicationYear INTEGER
		);`
	_, err := db.Exec(createTable)
	if err != nil {
		log.Printf("Failed to create books table: %v", err)
		return err
	}
	return nil
}
