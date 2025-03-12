// An example web application for tracking user visits, using sqlite3 for persistence.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"
)

// Visits represents the number of visits to an HTTP endpoint.
type Visits struct {
	Count int `json:"count"`
}

const schema = `
	CREATE TABLE IF NOT EXISTS visits (
		count INTEGER PRIMARY KEY
	);
`

func main() {
	dbFile := flag.String("f", "example.db?_journal_mode=wal", "sqlite DB file")
	flag.Parse()

	db := sqlx.MustConnect("sqlite", *dbFile)
	defer func() {
		err := db.Close()
		if err != nil {
			log.Fatalf("error closing db: %v", err)
		}
	}()

	db.MustExec(schema)
	db.MustExec(`INSERT OR IGNORE INTO visits VALUES (0);`) // Seed row if necessary.

	var visits Visits
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := db.Exec(`UPDATE visits SET count = count + 1;`)
		if err != nil {
			msg := fmt.Sprintf("unable to UPDATE count in sqlite visits table: %v", err)
			http.Error(w, msg, http.StatusInternalServerError)
			return
		}

		err = db.Get(&visits, `SELECT count FROM visits;`)
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

	fmt.Println("Listening on http://0.0.0.0:9999")
	panic(http.ListenAndServe(":9999", nil))
}
