package main

import (
	"context"
	"database/sql"
	"log"
	"reflect"

	"example.com/sqlc-playground/tutorial"
	_ "modernc.org/sqlite"
)

func run() error {
	ctx := context.Background()
	db, err := sql.Open("sqlite", "./playground.db?_pragma=foreign_keys(1)")
	if err != nil {
		return err
	}
	defer db.Close()

	queries := tutorial.New(db)
	// list all authors
	authors, err := queries.ListAuthors(ctx)
	if err != nil {
		return err
	}
	log.Println(authors)
	// create an author
	insertedAuthor, err := queries.CreateAuthor(ctx, tutorial.CreateAuthorParams{
		Name: "Brian Kernighan",
		Bio:  sql.NullString{String: "Co-author of The C Programming Language and The Go Programming Language", Valid: true},
	})
	if err != nil {
		return err
	}
	log.Println(insertedAuthor)
	// get the author we just inserted
	fetchedAuthor, err := queries.GetAuthor(ctx, insertedAuthor.ID)
	if err != nil {
		return err
	}
	// prints true
	log.Println(reflect.DeepEqual(insertedAuthor, fetchedAuthor))

	insertedBook, err := queries.CreateBook(
		ctx, tutorial.CreateBookParams{
			Name:     "Go Programming 101",
			AuthorID: sql.NullInt64{Int64: fetchedAuthor.ID, Valid: true},
		})
	if err != nil {
		return err
	}

	log.Println(insertedBook)

	return nil

}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
