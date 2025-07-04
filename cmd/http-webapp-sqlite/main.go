// An example web application for tracking user visits, using sqlite3 for persistence.
package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
)

// Visits represents the number of visits to an HTTP endpoint.
type Visits struct {
	ID int `json:"id"`
	Count int `json:"count"`
}

const schema = `
	CREATE TABLE IF NOT EXISTS visits (
		id INTEGER PRIMARY KEY,
		count INTEGER
	);
`

func main() {
	dbFile := flag.String("f", "example.db?_journal_mode=wal", "sqlite DB file")
	flag.Parse()

	db, err := sql.Open("sqlite3", "file:" + *dbFile)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("error closing db: %v", err)
		}
	}()

	ctx := context.Background()
	if _, err := db.ExecContext(ctx, schema); err != nil {
		log.Fatal(err)
	}

	if _, err := db.ExecContext(ctx, `INSERT OR IGNORE INTO visits VALUES (1, 0);`); err != nil {
		log.Fatal(err)
	}

	var visits Visits
	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		err := db.QueryRowContext(ctx, `SELECT count FROM visits;`).Scan(&visits.Count)
		if err != nil {
			msg := fmt.Sprintf("unable to SELECT from visits table: %v", err)
			http.Error(w, msg, http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-type", "application/json")
		err = json.NewEncoder(w).Encode(&visits)
		if err != nil {
			msg := fmt.Sprintf("error encoding visits: %v", err)
			http.Error(w, msg, http.StatusInternalServerError)
			return
		}
	})

	http.HandleFunc("POST /", func(w http.ResponseWriter, r *http.Request) {
		_, err := db.ExecContext(ctx, `UPDATE visits SET count = count + 1 WHERE ID = 1;`)
		if err != nil {
			msg := fmt.Sprintf("unable to UPDATE count in sqlite visits table: %v", err)
			http.Error(w, msg, http.StatusInternalServerError)
			return
		}
	})

	fmt.Println("Listening on http://0.0.0.0:9999")
	panic(http.ListenAndServe(":9999", nil))
}
