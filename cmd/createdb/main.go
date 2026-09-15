package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

func main() {
	log.SetFlags(0)
	dbPath := flag.String("db", "./new.db", "path of the SQLite database file to create")
	flag.Parse()
	if flag.NArg() != 0 {
		log.Fatal("usage: go run ./cmd/createdb [-db path]")
	}

	path, err := filepath.Abs(*dbPath)
	if err != nil {
		log.Fatal(err)
	}
	if err := createDB(path); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Created SQLite database: %s\n", path)
}

func createDB(path string) error {
	// Reserve a new file without overwriting an existing database.
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0600)
	if err != nil {
		return fmt.Errorf("create database file: %w", err)
	}
	initialized := false
	defer func() {
		if !initialized {
			_ = os.Remove(path)
		}
	}()
	if err := file.Close(); err != nil {
		return fmt.Errorf("close new database file: %w", err)
	}

	db, err := sql.Open("sqlite", path)
	if err != nil {
		return fmt.Errorf("open SQLite database: %w", err)
	}
	defer db.Close()

	// sql.Open is lazy. VACUUM writes a valid SQLite header and an empty schema.
	if _, err := db.Exec("VACUUM;"); err != nil {
		return fmt.Errorf("initialize SQLite database: %w", err)
	}
	if err := db.Close(); err != nil {
		return fmt.Errorf("close SQLite database: %w", err)
	}
	initialized = true
	return nil
}
