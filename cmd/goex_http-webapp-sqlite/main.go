// An example web application for tracking user visits, using sqlite3 for persistence.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

// Visits represents the number of visits to an HTTP endpoint.
type Visits struct {
	Count int `json:"count"`
}

const schema = `
	CREATE TABLE IF NOT EXISTS visits (
		id INTEGER PRIMARY KEY,
		count INTEGER NOT NULL
	);
`

func main() {

	dbFile := flag.String("f", "example.db?_journal_mode=wal", "sqlite DB file")
	flag.Parse()

	db := sqlx.MustConnect("sqlite3", *dbFile)
	defer db.Close()

	db.MustExec(schema)
	db.MustExec(`INSERT OR IGNORE INTO visits VALUES (1, 0);`) // Seed row.

	var visits Visits
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if _, err := db.Exec(`UPDATE visits SET count = count + 1;`); err != nil {
			http.Error(w, "unable to UPDATE sqlite visits table", http.StatusInternalServerError)
			return
		}
		if err := db.Get(&visits, `SELECT count FROM visits;`); err != nil {
			http.Error(w, "unable to SELECT from sqlite visits table", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode(&visits)
	})

	fmt.Println("Listening on http://0.0.0.0:9999")
	panic(http.ListenAndServe(":9999", nil))
}
