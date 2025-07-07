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
	Count int `json:"count"`
}

const schema = `
	CREATE TABLE IF NOT EXISTS visits (
		id INTEGER PRIMARY KEY,
		count INTEGER
	);
`

type App struct {
	db *sql.DB
}

func main() {
	dbFile := flag.String("f", "example.db?_journal_mode=wal", "sqlite DB file")
	flag.Parse()

	db, err := sql.Open("sqlite3", "file:"+*dbFile)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("error closing db: %v", err)
		}
	}()

	app := App{db: db}

	ctx := context.Background()
	if _, err := app.db.ExecContext(ctx, schema); err != nil {
		log.Fatal(err)
	}

	if _, err := app.db.ExecContext(ctx, `INSERT OR IGNORE INTO visits VALUES (1, 0);`); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("GET /", app.handleGetVisits)
	http.HandleFunc("POST /", app.handleUpdateVisits)

	fmt.Println("Listening on http://0.0.0.0:9999")
	log.Fatal(http.ListenAndServe(":9999", nil))
}

func (a *App) handleGetVisits(w http.ResponseWriter, r *http.Request) {
	var visits Visits
	ctx := context.Background()
	err := a.db.QueryRowContext(ctx, `SELECT count FROM visits;`).Scan(&visits.Count)
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
}

func (a *App) handleUpdateVisits(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	_, err := a.db.ExecContext(ctx, `UPDATE visits SET count = count + 1 WHERE id = 1;`)
	if err != nil {
		msg := fmt.Sprintf("unable to UPDATE count in sqlite visits table: %v", err)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
}
