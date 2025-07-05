package main

import (
	"context"
	"database/sql"
	_ "embed"
	"log"
	"reflect"

	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"

	dbsqlc "github.com/qjcg/examples-go/cmd/db-sqlc/db"
)

//go:embed schema.sql
var ddl string

func main() {
	db, err := sql.Open("sqlite3", "file:example.db")
	check(err)

	ctx := context.Background()

	// create tables
	_, err = db.ExecContext(ctx, ddl)
	check(err)

	queries := dbsqlc.New(db)

	// list all authors
	authors, err := queries.ListAuthors(ctx)
	check(err)
	log.Println(authors)

	// create an author
	insertedAuthor, err := queries.CreateAuthor(ctx, dbsqlc.CreateAuthorParams{
		Name: "Brian Kernighan",
		Bio: sql.NullString{
			String: "Co-author of The C Programming Language and The Go Programming Language",
			Valid:  true,
		},
	})
	check(err)
	log.Println(insertedAuthor)

	// get the author we just inserted
	fetchedAuthor, err := queries.GetAuthor(ctx, insertedAuthor.ID)
	check(err)

	// prints true
	log.Println(reflect.DeepEqual(insertedAuthor, fetchedAuthor))
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
